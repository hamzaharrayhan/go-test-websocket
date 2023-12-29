package pages

const FormPage = `
<html>
<head>
    <title>Phone Number and Provider Form</title>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.6.0/jquery.min.js"></script>
</head>
<body>
    <h1>Phone Number and Provider Form</h1>
    <form id="phone-form" onsubmit="return false;">
        <label for="phone-number">Phone Number:</label>
        <input type="text" id="phone-number" name="number"><br><br>
        <label for="provider_id">Provider:</label>
        <select id="provider_id" name="provider_id"></select><br><br>
        <button type="button" id="random-button">Fill Random</button>
        <button id="save-button">Save</button>
    </form>

    <script>
        fetch('http://localhost:9090/provider')
            .then(response => response.json())
            .then(data => {
                const providers = data.data;

                providers.forEach(provider => {
                    const option = document.createElement('option');
                    option.value = provider.ID;  
                    option.text = provider.Provider;
                    document.getElementById('provider_id').appendChild(option);
                });
            })
            .catch(error => console.error('Error fetching providers:', error));

        $('#random-button').click(() => {
            const phoneNumber = generateRandomPhoneNumber();
            const randomProviderIndex = Math.floor(Math.random() * $('#provider_id option').length);
            $('#phone-number').val(phoneNumber);
            $('#provider_id').val($('#provider_id option')[randomProviderIndex].value);
        });

        $('#save-button').click(() => {
            const formData = $('#phone-form').serialize();
            $.post('http://localhost:9090/add-phonenumber', formData)
                .then(response => {
                    console.log('Data saved successfully!');
                    window.location.replace("http://localhost:9090/dashboard")
                })
                .catch(error => {
                    console.error('Error saving data:', error);
                });
            });

        function generateRandomPhoneNumber() {
            const length = Math.floor(Math.random() * (12 - 4 + 1)) + 4;
            let randomString = '';

            for (let i = 0; i < length; i++) {
                randomString += Math.floor(Math.random() * 10); // Generate a random digit (0-9)
            }
            return randomString;  // Example placeholder
        }
    </script>
</body>
</html>
`
