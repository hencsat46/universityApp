
function clickButton(elem) {
    const elemName = elem.textContent.trim();
    const element = document.querySelectorAll('.form');
    let displaySignIn = element[0].style.display
    let displaySignUp = element[1].style.display
    switch(elemName) {
        case "Войти":
            element[0].style.display = "block"
            element[1].style.display = "none"
            break;
        case "Зарегистрироваться":
            element[0].style.display = "none"
            element[1].style.display = "block"
            break;
    }
    
}

function closeForm(number) {
    const element = document.querySelectorAll('.form');
    element[number].style.display = "none"
}

function sendData(elem) {
    console.log()
    const dataArr = elem.parentElement.parentElement.getElementsByTagName("input")
    const dataJson = {status: elem.className, username: dataArr[0].value, password: dataArr[1].value}
    

    postData("http://localhost:3000", dataJson)
}

async function postData(url, data) {
    const response = await fetch(url, {
        method: "POST",
        mode: "cors",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
    console.log(await response.text())
}
