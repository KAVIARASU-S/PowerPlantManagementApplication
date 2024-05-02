document.addEventListener('DOMContentLoaded', () => {
    let form = document.getElementById('adminForm');
    let messageBox = document.getElementById('message');
    let usernameInput = document.getElementById('username');
    let passwordInput = document.getElementById('password');
    let qrbox = document.getElementById('qr')

    form.addEventListener('submit', (event) => {
        event.preventDefault();
        messageBox.innerHTML = '';

        let username = usernameInput.value;
        let password = passwordInput.value;
        let company = localStorage.getItem('company')
        let powerplant = document.getElementById('powerplant').value;
        let powerplanttype = localStorage.getItem('powerplanttype')

        // Construct JSON data
        let userData = {
            "UserName": username,
            "Password": password,
            "Company": company,
            "Role": powerplant, // Assuming role as Admin
            "PowerplantType": powerplanttype
        };

        // Send JSON data to the server
        fetch('http://localhost:2000/createUser', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(userData)
        })
        .then(response => {
            if (response.status === 409) {
                return response.json().then(data => {
                    throw new Error(data.Error); // Throw the error message
                });
            } else if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            // Hide the form
            form.style.display = 'none';

            let container = document.querySelector('.titleText');
            container.style.display = 'none';

            let formcontainer = document.querySelector('.form-container')

            formcontainer.style.width = '25%';
            formcontainer.style.padding = '20px';
            formcontainer.style.paddingLeft = '40px';
            formcontainer.style.paddingRight = '1px'
            formcontainer.style.borderRadius = '8px';
            formcontainer.style.backgroundColor = '#fff';
            formcontainer.style.boxShadow = '0 2px 4px rgba(0, 0, 0, 0.1)';
            
            // Display success message and QR code
            qrbox.innerHTML = '<h3 style="color: rgb(46, 250, 63);">' + "User Created Successfully" + '</h3>';
            qrbox.innerHTML +='<p style="width: 100%;">' + "Please scan and securely store the OTP token in your Authenticator app." + "</p>"
            qrbox.innerHTML += '<p style="width: 100%;">' + "Once lost, tokens cannot be retrieved." + "</p>"
            qrbox.innerHTML += '<div style="margin-left: 15px;">' + '<img src="data:image/png;base64,' + data.QR + '" alt="QR Code">' + "</div>";
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
            messageBox.textContent = error.message; // Display the error message
        });
    });
});
