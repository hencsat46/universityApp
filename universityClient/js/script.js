
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

async function getUniversity(url, data) {
    const response = await fetch(url, {
        method: "GET",
        mode: "cors",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })

    
    return response.json()
}

function makeJsonUniversity() {
    const universityCount = document.querySelectorAll(".uni-wrapper").length
    console.log(document.documentElement.getBoundingClientRect().bottom)
    console.log(document.documentElement.clientHeight)
    if (document.documentElement.getBoundingClientRect().bottom + 50 < document.documentElement.clientHeight) {
        console.log("you reached the end of page")
        const universityOrder = universityCount - 2
        const requestJson = {order: universityOrder}
        const response = getUniversity("http://localhost:3000/getUniversity", requestJson)
        console.log(response)
    }
}

function tempFunc() {
    
    console.log()
    console.log()

    
}

window.onscroll=makeJsonUniversity

