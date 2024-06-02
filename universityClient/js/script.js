function clickButton(elem) {
    const elemName = elem.textContent.trim();
    const element = document.querySelectorAll('.form');
    document.body.classList.add('no-scroll')

    const windowHeight = window.innerHeight || document.documentElement.clientHeight || document.body.clientHeight;
    
    const scrollHeight = window.scrollY || document.documentElement.scrollTop || document.body.scrollTop;

    const topOffset = (windowHeight / 2) - (element[0].offsetHeight / 2) + scrollHeight;
    
    switch(elemName) {
        case "Войти":
            element[0].style.display = "flex"
            element[1].style.display = "none"
            element[0].style.height = "230px"
            element[0].style.top = topOffset + 'px';
            document.querySelector("ul.ul").style.filter = "blur(3px)"
            break;
        case "Зарегистрироваться":
            element[0].style.display = "none"
            element[1].style.display = "flex"
            element[1].style.height = "390px"
            element[1].style.top = topOffset + 'px';
            document.querySelector("ul.ul").style.filter = "blur(3px)"
            break;
    }
    
}

function closeForm(element) {
    document.body.classList.remove('no-scroll')
    document.querySelector("ul.ul").style.filter = "blur(0px)"
    element.parentNode.style.display = "none"
}

function sendData(elem) {
    const dataArr = elem.parentElement.parentElement.getElementsByTagName("input")
    const dataJson = {username: dataArr[0].value, password: dataArr[1].value}
    let response = undefined
    console.log(elem.className)
    switch (elem.className) {
        case "sign-up":
            signUpButton(elem)
            closeForm(elem.parentNode)
            break
        case "sign-in":
            signInButton(elem)
            closeForm(elem.parentNode)
            break
    }
    
}

async function signInButton(element) {
    const dataArr = element.parentElement.parentElement.getElementsByTagName("input")
    const json = {Username: dataArr[0].value, Password: dataArr[1].value}
    
    const request = new Request("http://localhost:3000/signin", {
        method: "POST",
        mode: "cors",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(json)
    })
    
    
    const response = await (await fetch(request)).json()
    
    if (response.Status != 401) {
        setCookie("")
        let date = new Date()
        date.setMinutes(date.getMinutes() + 30)
        setCookie(response.Payload + "; expires=" + date)
        document.querySelector(".sign-btns").style.display = "none"
        document.querySelector(".username-icon").style.display = "flex"
    }
}

async function signUpButton(element) {
    const dataArr = element.parentElement.parentElement.getElementsByTagName("input")
    const data = {
        StudentName: dataArr[0].value, 
        StudentSurname: dataArr[1].value, 
        Username: dataArr[2].value, 
        Password: dataArr[3].value,
    }

    const request = new Request("http://localhost:3000/signup", {
        method: "POST",
        mode: "cors",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data)
    })

    console.log(request)

    const response = (await fetch(request)).json()
    console.log(await response)
}

function getToken() {
    const cookies = document.cookie
    let cookieArr = cookies.split(';')
    const length = cookieArr.length
    for (let i = 0; i < length; i++) {
        cookieArr[i] = cookieArr[i].trim()
    }

    for (let i = 0; i < length; i++) {
        const tempArr = cookieArr[i].split('=')
        console.log(tempArr[0])
        if (tempArr[0] == 'Token') return tempArr[1]
    }
    return null
}

function setCookie(stringToken) {
    console.log(stringToken)
    document.cookie = `Token=${stringToken}`
    console.log(document.cookie)
}

async function getUniversity(url, data) {
    const response = await fetch(url, {
        method: "GET",
        mode: "cors",
    })
    
    return response.json()
}

let flag = true

let universityName = ""

function setUniversity(element) {
    
    const docsForm = document.querySelector(".docs-form")
    docsForm.style.display = "flex"
    document.body.classList.add('no-scroll')
    document.querySelector("ul.ul").style.filter = "blur(3px)"
    const popup = document.querySelector('.docs-form');
    const windowHeight = window.innerHeight || document.documentElement.clientHeight || document.body.clientHeight;
    const scrollHeight = window.scrollY || document.documentElement.scrollTop || document.body.scrollTop;
    const topOffset = (windowHeight / 2) - (popup.offsetHeight / 2) + scrollHeight;
    popup.style.top = topOffset + 'px';
    popup.style.display = 'flex';

    universityName = element.parentNode.querySelector("h2").innerHTML
    
}

async function autoLogin() {
    const token = getToken()
    console.log(token)
    if (token == null) return
    const request = new Request("http://localhost:3000/", {
        method: "GET",
        mode: "cors",
        headers : {
            "Token": token,
        }
    })
    console.log("aaaaaaaaaaaaa")
    console.log(request)
    console.log("hello")
    const response = await (await fetch(request)).json()

    console.log(response)

    if (response.Status == 200 && response.Payload == "Sign in ok") {
        document.querySelector(".sign-btns").style.display = "none"
        document.querySelector(".username-icon").style.display = "flex"
    }
}

autoLogin()

function submitUniversity(element) {
    const pointsString = element.parentNode.querySelector(".docs-input").value
    
    const dataObject = {
        University: universityName,
        Points: pointsString
    }
    

    const response = requestUniversity("http://localhost:3000/addStudent", dataObject)
    response.then(value => console.log(value))
}

async function requestUniversity(url, data) {
    const request = new Request(url, {
        method: "POST",
        mode: "cors",
        headers: {
            "Content-Type": "application/json",
            "Token": getToken(),
        },
        body: JSON.stringify(data)
    })

    const response = (await fetch(request)).json()

    return response



}

makeJsonUniversity()

async function makeJsonUniversity() {
    let universityCount = document.querySelectorAll(".uni-wrapper").length
    
    
    const response = await fetch("http://localhost:3000/get_universities", {
        method: "GET",
        mode: 'cors',
    })
    const universityObject = await response.json()
    
    console.log(universityObject)


    const universities = new Array()

    for (let i = 0; i < universityObject.Payload.length / 2; i++) {
        const universityElem = new Array()
        universityElem.push(universityObject.Payload[i * 2])
        universityElem.push(universityObject.Payload[i * 2 + 1])
        universities.push(universityElem)
        console.log("жопа")
    }
    console.log(universities)
    //makeUniversityElem(universityObject[0].Uni_name, universityObject[0].Uni_des, universityObject[0].Uni_img, universityObject[1].Uni_name, universityObject[1].Uni_des, universityObject[1].Uni_img)

    for (let i = 0; i< universities.length; i++) {
        makeUniversityElem(universities[i][0].Uni_name, universities[i][0].Uni_des, universities[i][0].Uni_img, universities[i][1].Uni_name, universities[i][1].Uni_des, universities[i][1].Uni_img)
    }
        
}

async function makeUniversityElem(firstName, firstDescription, firstImg, secondName, secondDescription, secondImg) {
    let firstUni = ""
    let secondUni = ""
    if (firstName.length != 0) {
        firstUni = `
        <div class="uni-wrapper">
            <div class="info-container">
                <img src="${firstImg}" alt="x">
                <div class="text-container">
                    <h2>${firstName}</h2>
                    <div class="uni-text">
                        ${firstDescription} 
                    </div>
                    <button class="sub-docs" onclick="setUniversity(this)">
                        <div class="btn-text">Подать документы</div>
                    </button>
                </div>
            </div>
        </div>
        `
    }
    if (secondName.length != 0) {
        secondUni = `
        <div class="uni-wrapper">
            <div class="info-container">
                <img src="${secondImg}" alt="x">
                <div class="text-container">
                    <h2>${secondImg}</h2>
                    <div class="uni-text">
                        ${secondDescription} 
                    </div>
                    <button class="sub-docs" onclick="setUniversity(this)">
                        <div class="btn-text">Подать документы</div>
                    </button>
                </div>
            </div>
        </div>
        `
    }
    const liElement = document.createElement("li")
    liElement.innerHTML = firstUni + secondUni
    document.querySelector("ul.ul").append(liElement)
}





