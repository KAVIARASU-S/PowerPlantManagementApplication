document.addEventListener('DOMContentLoaded', function() {
    // // Call the getPublicIP function to fetch the public IP address
    getPublicIP()
    // Add event listener to the submit button click event
    document.getElementById('submitBtn').addEventListener('click', function(event) {
        event.preventDefault(); // Prevent the default form submission behavior

        // Call the login function
        login();
    });
});

// Define the login function
function login() {
    // Get the username and password input fields
    let usernameInput = document.getElementById('username');
    let passwordInput = document.getElementById('password');
    let ipaddressInput = document.getElementById('publicip');

    // Get the username and password values
    let username = usernameInput.value;
    let password = passwordInput.value;
    let ipaddress = ipaddressInput.innerText;

    // Create a data object to send to the server
    let logindata = {
        UserName: username,
        Password: password
    };

    let ipdata = {
        IPaddress:ipaddress
    }
console.log("the ip adddres",ipdata.IPaddress)
    // Send a POST request to the server
    fetch('http://localhost:2000/signIn', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            user: logindata,
            ip: ipdata
        })
    })
    .then(function(response) {
        if (response.ok) {
            localStorage.setItem('username',username)
            window.location.href = '../totp/totp.html'; // Redirect for successful response
        } else {
            if (response.status === 400 || response.status === 401) {
                response.json().then(function(data) {
                    if (data && data.Error) {
                        document.getElementById('result').innerText = data.Error; // Display error message
                    } else {
                        document.getElementById('result').innerText = 'An error occurred'; // Default error message
                    }
                }).catch(function(error) {
                    console.error('Error parsing JSON:', error);
                    document.getElementById('result').innerText = 'An error occurred'; // Default error message
                });
            } else {
                document.getElementById('result').innerText = 'An error occurred'; // Generic error message
            }
        }
    })
    .catch(function(error) {
        console.error('Error fetching data:', error);
        document.getElementById('result').innerText = 'An error occurred'; // Default error message
    });
}


function getPublicIP() {
    fetch('https://api.ipify.org?format=json')
        .then(response => response.json())
        .then(data => {
            document.getElementById('publicip').textContent = data.ip;
        })
        .catch(error => {
            console.error('Error fetching public IP:', error);
            document.getElementById('publicip').textContent = 'Error fetching IP';
        });
}