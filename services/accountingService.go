package services

import (
	"PowerPlantManagementApplication/interfaces"
	"PowerPlantManagementApplication/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccountingService struct {
	AccountingCollection *mongo.Collection
}

func InitAccounting (accountingCollection *mongo.Collection)interfaces.IAccounting{
	return &AccountingService{
		AccountingCollection: accountingCollection,
	}
}

func (accountingData *AccountingService) DisplayTransactions(searchFilter *models.SearchFilter)(allTransactions *[]models.Transaction,err error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"CompanyName":searchFilter.CompanyName}
	opts := options.Find().SetSort(map[string]int{"Date": -1})

	result, err := accountingData.AccountingCollection.Find(ctx, filter,opts)
	if err != nil {
		log.Println("Error Finding transactions in mongoDB ", err)
		return nil, err
	}

	log.Println("Successfully found Transactions in mongoDb")

	var transactions []models.Transaction

	err = result.All(ctx, &transactions)

	if err != nil {
		log.Println("Error parsing the Transactions to slice",err)
	}

	return &transactions,nil
}

func (accountingData *AccountingService) InsertTransaction(transaction *models.Transaction)(err error){
	log.Println("Transaction to be inserted ", transaction)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := accountingData.AccountingCollection.InsertOne(ctx, transaction)

	if err != nil {
		log.Println("Error inserting transaction in mongoDB",err)
		return err
	}

	log.Println("Successfully Inserted Transaction", result)
	return nil
}

func (accountingData *AccountingService) DisplayAccounting(searchFilter *models.SearchFilter)(allAccounts *models.Accounting,err error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Now we find the total of incomes

	totIncomefilter := bson.M{
		"CompanyName":searchFilter.CompanyName,
		"TransactionType": "Income",
	}

	totIncomepipeline := bson.A{
        bson.D{{"$match", totIncomefilter}},
        bson.D{{"$group", bson.D{{"_id", nil}, {"totalIncome", bson.D{{"$sum", "$Amount"}}}}}},
    }

	totIncomeresult,err := accountingData.AccountingCollection.Aggregate(ctx,totIncomepipeline)
	if err != nil {
		log.Println("Error Finding transactions in mongoDB ", err)
		return nil,err
	}

	var income []bson.M
	if err := totIncomeresult.All(ctx, &income); err != nil {
        log.Println("Unable to decode total income to slice",err)
		return nil,err
    }

	var totalincome int

	if len(income) > 0 {
		switch val := income[0]["totalIncome"].(type) {
		case int:
			totalincome = val
		case int64:
			totalincome = int(val)
		case int32:
			totalincome = int(val)
		case float64:
			totalincome = int(val)
		case float32:
			totalincome = int(val)
		default:
			log.Println("Failed to convert totalIncome to int")
		}
	} else {
		// Handle the case where there are no income transactions foud
		log.Println("No income transactions found.")
	}
	
	//Now we find the total of expenses

	totExpensefilter := bson.M{
		"CompanyName": searchFilter.CompanyName,
		"TransactionType": "Expense",
	}

	totExpensepipeline := bson.A{
        bson.D{{"$match", totExpensefilter}},
        bson.D{{"$group", bson.D{{"_id", nil}, {"totalExpense", bson.D{{"$sum", "$Amount"}}}}}},
    }

	totExpenseresult,err := accountingData.AccountingCollection.Aggregate(ctx,totExpensepipeline)
	if err != nil {
		log.Println("Error Finding transactions in mongoDB ", err)
		return nil,err
	}

	var expense []bson.M
	if err := totExpenseresult.All(ctx, &expense); err != nil {
        log.Println("Failed to decode total expense",err)
    }

	var totalExpense int

	if len(expense) > 0 {
		switch val := expense[0]["totalExpense"].(type) {
		case int:
			totalExpense = val
		case int64:
			totalExpense = int(val)
		case int32:
			totalExpense = int(val)
		case float64:
			totalExpense = int(val)
		case float32:
			totalExpense = int(val)
		default:
			log.Println("Failed to convert totalIncome to int")
		}
	} else {
		// Handle the case where there are no income transactions foud
		log.Println("No income transactions found.")
	}

	// Now we get all the income

	incomefilter := bson.M{
		"CompanyName": searchFilter.CompanyName,
		"TransactionType": "Income",
	}

	incomeopts := options.Find().SetSort(map[string]int{"Date": 1})

	incomes,err := accountingData.AccountingCollection.Find(ctx,incomefilter,incomeopts)
	if err != nil {
		log.Println("Error finding incomes in mongoDB",err)
		return nil,err
	}

	var allincomes []models.Transaction

	err = incomes.All(ctx,&allincomes)
	if err != nil {
		log.Println("Error parsing the Incomes to slice",err)
	}

	// Now we get all the expenses

	expenseFilter := bson.M{
		"CompanyName": searchFilter.CompanyName,
		"TransactionType": "Expense",
	}

	expenseopts := options.Find().SetSort(map[string]int{"Date": 1})

	expenses,err := accountingData.AccountingCollection.Find(ctx,expenseFilter,expenseopts)
	if err != nil {
		log.Println("Error finding expenses in mongoDB",err)
		return nil,err
	}

	var allexpenses []models.Transaction

	err = expenses.All(ctx,&allexpenses)
	if err != nil {
		log.Println("Error parsing the Expenses to slice",err)
	}

	// Now we calculate the profit
	profit := totalincome - totalExpense

	//We are assigning the values to the response
	var accounts models.Accounting

	accounts.Incomes = allincomes
	accounts.Expenses = allexpenses
	accounts.TotalIncome = totalincome
	accounts.TotalExpense = totalExpense
	accounts.Profit = profit

	log.Println("The accounts is",accounts)
	return &accounts,nil
}