
function clickButton(elem) {
    const elemName = elem.textContent.trim();
    const element = document.querySelectorAll('.form');
    console.log(element)
    switch(elemName) {
        case "Войти":
            element[0].style.display = "block"
            element[1].style.display = "none"
            document.querySelector("ul.ul").style.filter = "blur(3px)"
            break;
        case "Зарегистрироваться":
            element[0].style.display = "none"
            element[1].style.display = "block"
            document.querySelector("ul.ul").style.filter = "blur(3px)"
            break;
    }
    
}

function closeForm(number) {
    const element = document.querySelectorAll('.form');
    element[number].style.display = "none"
    document.querySelector("ul.ul").style.filter = "blur(0px)"
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

async function tempFunc() {
    
    console.log(document.documentElement.getBoundingClientRect().bottom)
    console.log(document.documentElement.clientHeight)
}

window.onscroll=tempFunc

