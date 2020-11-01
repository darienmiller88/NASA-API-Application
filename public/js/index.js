const buttonRef = document.getElementById('date-button')
const inputRef = document.getElementById('date')

buttonRef.addEventListener('click', async e => {
    const date = inputRef.value
    const inputData = {
        date,
    }
    console.log(inputData)

    const response = await fetch(`/apod/${date}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
    });
    const response_json = await response.json();
    console.log(response_json);
})