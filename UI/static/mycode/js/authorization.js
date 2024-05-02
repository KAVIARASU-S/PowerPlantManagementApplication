document.addEventListener('DOMContentLoaded', function() {
    let role = localStorage.getItem('role');
console.log("Value of role",role)
    if (role === 'SuperAdmin'){
        document.getElementById('dashboard0').style.display='none';
        document.getElementById('roi0').style.display='none';
        document.getElementById('accounting0').style.display='none';
        document.getElementById('transactions0').style.display='none';
        document.getElementById('tasks0').style.display='none';
        document.getElementById('inventory0').style.display='none';
        document.getElementById('purchase0').style.display='none';
        document.getElementById('livedata0').style.display='none';
        document.getElementById('sensors0').style.display='none';
        document.getElementById('calculator0').style.display='none';        
    }
    if (role == 'CompanyAdmin'){
        document.getElementById('superadmindashboard0').style.display='none';
        document.getElementById('ipaddress0').style.display='none';
        document.getElementById('transactions0').style.display='none';
        document.getElementById('tasks0').style.display='none';
        document.getElementById('inventory0').style.display='none';
        document.getElementById('purchase0').style.display='none';
        document.getElementById('sensors0').style.display='none';
        document.getElementById('calculator0').style.display='none';    
    }
    if (role == 'Accounts'){
        document.getElementById('dashboard0').style.display='none';
        document.getElementById('superadmindashboard0').style.display='none';
        document.getElementById('ipaddress0').style.display='none';
        document.getElementById('roi0').style.display='none';
        document.getElementById('tasks0').style.display='none';
        document.getElementById('inventory0').style.display='none';
        document.getElementById('purchase0').style.display='none';
        document.getElementById('livedata0').style.display='none';
        document.getElementById('sensors0').style.display='none';
        document.getElementById('calculator0').style.display='none';
        document.getElementById('users0').style.display='none';
    }
    if (role == 'TechnicalManager'){
        document.getElementById('dashboard0').style.display='none';
        document.getElementById('superadmindashboard0').style.display='none';
        document.getElementById('ipaddress0').style.display='none';
        document.getElementById('roi0').style.display='none';
        document.getElementById('accounting0').style.display='none';
        document.getElementById('transactions0').style.display='none';
        document.getElementById('inventory0').style.display='none';
        document.getElementById('purchase0').style.display='none';
        document.getElementById('sensors0').style.display='none';
        document.getElementById('calculator0').style.display='none'; 
        document.getElementById('users0').style.display='none';
    }
    if (role == 'Inventory'){
        document.getElementById('dashboard0').style.display='none';
        document.getElementById('superadmindashboard0').style.display='none';
        document.getElementById('roi0').style.display='none';
        document.getElementById('accounting0').style.display='none';
        document.getElementById('transactions0').style.display='none';
        document.getElementById('tasks0').style.display='none';
        document.getElementById('livedata0').style.display='none';
        document.getElementById('sensors0').style.display='none';
        document.getElementById('calculator0').style.display='none';
        document.getElementById('ipaddress0').style.display='none';
        document.getElementById('users0').style.display='none';
    }
    if (role == 'Engineer'){
        document.getElementById('dashboard0').style.display='none';
        document.getElementById('superadmindashboard0').style.display='none';
        document.getElementById('roi0').style.display='none';
        document.getElementById('accounting0').style.display='none';
        document.getElementById('transactions0').style.display='none';
        document.getElementById('tasks0').style.display='none';
        document.getElementById('inventory0').style.display='none';
        document.getElementById('purchase0').style.display='none';
        document.getElementById('ipaddress0').style.display='none';
        document.getElementById('users0').style.display='none';
    }
})
