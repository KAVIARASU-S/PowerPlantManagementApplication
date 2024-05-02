document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('otpForm').addEventListener('submit', function(event) {
         event.preventDefault(); // Prevent the form from submitting normally
 
         let totpValue = document.getElementById('totp').value; // Get the value of TOTP
         let username = localStorage.getItem('username'); // Get the username from local storage
 
         fetch('http://localhost:2000/totp', {
             method: 'POST',
             headers: {
                 'Content-Type': 'application/json',
             },
             body: JSON.stringify({ UserName: username, Totp: totpValue }),
         })
         .then(function(response) {
             if (response.ok){
                response.json().then(function(data) {
                     if (data && data.Company) {
                         localStorage.setItem('company',data.Company); // Store company in local storage
                         localStorage.setItem('role',data.Role);
                         localStorage.setItem('powerplanttype',data.PowerPlantType);
                     }
                    //  window.location.href='../static/index.html';
                    redirectUser()
                 }).catch(function(error) {
                     console.error('Error parsing JSON:', error);
                     document.getElementById('responseMessage').innerText = 'An error occurred'; // Default error message
                 });
             }else{
                 if (response.status === 400 || response.status === 401) {
                     response.json().then(function(data) {
                         if (data && data.Error) {
                             document.getElementById('responseMessage').innerText = data.Error; // Display error message
                         } else {
                             document.getElementById('responseMessage').innerText = 'An error occurred'; // Default error message
                         }
                     }).catch(function(error) {
                         console.error('Error parsing JSON:', error);
                         document.getElementById('responseMessage').innerText = 'An error occurred'; // Default error message
                     });
                 } else {
                     document.getElementById('responseMessage').innerText = 'An error occurred'; // Generic error message
                 }
             } 
         })
         .catch(function(error) {
             console.error('Fetch error:', error);
             document.getElementById('responseMessage').textContent = 'Error: ' + error.message;
         });
     });
 });
 
 function redirectUser(){
    let role = localStorage.getItem('role')

    if(role === 'SuperAdmin'){
        window.location.href="../static/dashboard-superadmin.html"
    }
    
    if (role === 'CompanyAdmin'){
        window.location.href='../static/index.html';
    }

    if (role === 'Accounts'){
        window.location.href='../static/pages-accounting.html';
    }

    if (role === 'TechnicalManager'){
        window.location.href='../static/pages-tasks.html';
    }

    if (role === 'Inventory'){
        window.location.href='../static/pages-inventory.html';
    }

    if (role === 'Engineer'){
        window.location.href='../static/pages-livedata.html';
    }
 }