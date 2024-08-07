function parseJSON() {
    const jsonInput = document.getElementById('jsonInput').value;
    fetch('http://localhost:8080/parse', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: jsonInput
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('jsonOutput').textContent = JSON.stringify(data, null, 2);
    })
    .catch(error => {
        document.getElementById('jsonOutput').textContent = 'Error: ' + error;
    });
}
