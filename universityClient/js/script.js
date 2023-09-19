function clickButton(elem) {
    const elemName = elem.textContent.trim();
    const element = document.querySelectorAll('.form');
    const formBg = document.querySelector('.form-bg')
    formBg.style.display = "block"
    document.body.classList.add('no-scroll')
    switch(elemName) {
        case "Войти":
            element[0].style.display = "block"
            element[1].style.display = "none"
            //document.querySelector("ul.ul").style.filter = "blur(3px)"
            break;
        case "Зарегистрироваться":
            element[0].style.display = "none"
            element[1].style.display = "block"
            break;
    }
    
}

function closeForm() {
    document.body.classList.remove('no-scroll')
    const formBg = document.querySelector('.form-bg')
    formBg.style.display = "none"
}

function showDocs(elem) {

}

function sendData(elem) {
    const dataArr = elem.parentElement.parentElement.getElementsByTagName("input")
    const dataJson = {username: dataArr[0].value, password: dataArr[1].value}
    let response = undefined
    switch (elem.className) {
        case "sign-up":
            postData("http://localhost:3000/signup", dataJson)
            //console.log(response)
            break
        case "sign-in":
            signIn("http://localhost:3000/token", dataJson)
            console.log(dataJson)
            break
    }
    
}

function getCookie() {
    const cookies = document.cookie
    let cookieArr = cookies.split(';')
    const length = cookieArr.length
    for (let i = 0; i < length; i++) {
        cookieArr[i] = cookieArr[i].trim()
    }

    for (let i = 0; i < length; i++) {
        const tempArr = cookieArr[i].split('=')
        if (tempArr[0] == 'Token') return tempArr[1]
    }
    return null
}

function setCookie(stringToken) {
    console.log(stringToken)

    document.cookie = `Token=${stringToken}`
    console.log(document.cookie)
}

function signIn(url, data) {
    const response = signPost(url, data)
    console.log(response, "PENIS")
    response.then(value => {return value}).then(value => {
        if ('Token' in value) {
            console.log("PENIS!!!")
            setCookie(value.Token)
            
        }
    })
}

async function signPost(url, data) {
    const request = new Request(url, {
        method: "POST",
        mode: "cors",
        headers: {
            "Content-Type": "application/json",
            "Token": getCookie()
        },
        body: JSON.stringify(data)
    })

    const response = await fetch(request)
    return response.json()

}

async function postData(url, data) {

    const request = new Request(url, {
        method: "POST",
        mode: "cors",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })

    const response = await fetch(request)
    console.log(response.json())
    //return response.json()
}

async function getUniversity(url, data) {
    const response = await fetch(url, {
        method: "POST",
        mode: "cors",
        headers: {
            "Content-Type": "application/json"
        },
        body: data
    })
    
    return response.json()
}
let remained = 1

let flag = true

let universityId = 0

function setUniversity(element) {
    const docsForm = document.querySelector(".docs-form")
    docsForm.style.display = "flex"
    docsForm.style.position = "absolute"
    docsForm.style.left = "50%"
    docsForm.style.top = "50%"
    docsForm.style.transform = "translate(-50%, -50%)"
    const formBg = document.querySelector('.form-bg')
    formBg.style.display = "block"
    universityId = element.parentNode.parentNode.parentNode.getAttribute('id')
    console.log(universityId)
}


async function makeJsonUniversity() {
    flag = false
    let universityCount = document.querySelectorAll(".uni-wrapper").length
    if (remained <= universityCount - 2) {
        return
    }

    if (document.documentElement.getBoundingClientRect().bottom < document.documentElement.clientHeight + 10) {
        const universityOrder = universityCount - 2
        const requestJson = `{"order": ${universityOrder}}`
        console.log(requestJson)
        const response = await getUniversity("http://localhost:3000/getUniversity", requestJson)
        
        makeUniversityElem(response.name, response.description, response.imagePath)
        if (!response.left) {
            return
        }
        remained = await parseInt(response.left)
    }
    flag = true
}

function access() {
    if (flag) {
        makeJsonUniversity()
    }
}

function makeUniversityElem(name, description, img) {
    const newLiElem = document.createElement("li")
    const newUniWrapper = document.createElement("div")
    const newInfoContainer = document.createElement("div")
    newInfoContainer.classList.add("info-container")
    newUniWrapper.classList.add("uni-wrapper")
    newUniWrapper.setAttribute("id", document.querySelectorAll(".uni-wrapper").length)
    const newH2Elem = document.createElement("h2")
    const newImgElem = document.createElement("img")
    const newUniText = document.createElement("uni-text")
    const newTextContainer = document.createElement("div")
    const newButtonElem = document.createElement("button")
    newButtonElem.classList.add("sub-docs")
    const newTextElem = document.createElement("div")
    newTextElem.classList.add("btn-text")
    newTextElem.innerHTML = "Подать документы"
    newButtonElem.append(newTextElem)
    newTextContainer.classList.add("text-container")
    newUniText.classList.add("uni-text")
    const ulClass = document.querySelector("ul.ul")
    newInfoContainer.append(newImgElem)
    ulClass.append(newLiElem)
    newLiElem.append(newUniWrapper)
    newUniWrapper.append(newInfoContainer)
    newTextContainer.append(newH2Elem)
    newTextContainer.append(newUniText)
    newTextContainer.append(newButtonElem)
    newInfoContainer.append(newTextContainer)
    newH2Elem.innerText = name
    newImgElem.setAttribute("src", img)
    newImgElem.setAttribute("alt", "x")
    newUniText.innerText = description
}



window.onscroll=access

//function 