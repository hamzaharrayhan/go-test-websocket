package pages

const DashboardPage = `
<html>
<head>
<title>Phone Number Data</title>
<style>
  table {
    border-collapse: collapse;
    width: 100%;
    margin-bottom: 20px;
  }
  th, td {
    border: 1px solid black;
    padding: 8px;
    text-align: left;
  }
  th {
    background-color: lightgray;
  }
  .selected {
    background-color: lightblue;
  }
</style>
</head>
<body>

<h1>Phone Number Data</h1>

<h2>Odd</h2>
<table id="oddTable">
  <thead>
    <tr>
      <th>ID</th>
      <th>Number</th>
      <th>Action</th>
    </tr>
  </thead>
  <tbody></tbody>
</table>

<h2>Even</h2>
<table id="evenTable">
  <thead>
    <tr>
      <th>ID</th>
      <th>Number</th>
      <th>Action</th>
    </tr>
  </thead>
  <tbody></tbody>
</table>

<script>
    const ws = new WebSocket("ws://localhost:9090/ws"); 

    ws.onmessage = (event) => {
        const data = JSON.parse(event.data); // Parse the JSON data
        updateTable(data); // Call a function to update the table
    }

    function updateTable(data) {
        const oddTable = document.getElementById('oddTable').querySelector('tbody');
        const evenTable = document.getElementById('evenTable').querySelector('tbody');
        
        oddTable.innerHTML = '';
        evenTable.innerHTML = '';
        
        data.forEach(phoneNumber => {
            const row = document.createElement('tr');
            const idCell = document.createElement('td');
            const numberCell = document.createElement('td');
            const actionsCell = document.createElement('td');

            idCell.textContent = phoneNumber.ID;
            numberCell.textContent = phoneNumber.Number;

            const deleteButton = document.createElement('button');
            const updateButton = document.createElement('button');
            deleteButton.textContent = 'Delete';
            updateButton.textContent = 'Update';
            deleteButton.addEventListener('click', () => {
                handleDelete(phoneNumber.ID);
            });
            updateButton.addEventListener('click', () => {
                let headers = new Headers();
                window.location.replace("http://localhost:9090/update-phonenumber?id="+phoneNumber.ID)
            })
            actionsCell.appendChild(deleteButton);
            actionsCell.appendChild(updateButton);

            row.appendChild(idCell);
            row.appendChild(numberCell);
            row.appendChild(actionsCell); 

            if (parseInt(phoneNumber.Number) % 2 === 0) {
                evenTable.appendChild(row);
            } else {
                oddTable.appendChild(row);
            }
        });
    }

    function handleDelete(id) {
        fetch('http://localhost:9090/phonenumber/'+id, {
            method: 'DELETE'
        })
        .then(response => {
            console.log('Success Delete PhoneNumber with ID '+id)
        })
        .catch(error => {
            console.error('Error deleting phone number:', error);
        });
    }

</script>

</body>
</html>
`
