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
    switch (elem.className) {
        case "sign-up":
            signUpButton(elem)
            break
        case "sign-in":
            signInButton(elem)
            break
    }
    
}

function signInButton(element) {
    const dataArr = element.parentElement.parentElement.getElementsByTagName("input")
    const json = {Username: dataArr[0].value, Password: dataArr[1].value}
    const response = signIn("http://localhost:3000/signin", json)

    response.then(value => {
        console.log(value)
        if ('Token' in value.Payload) {
            setCookie(value.Payload.Token)
            
        }
    })
}

async function signIn(url, json) {
    const request = new Request(url, {
        method: "POST",
        mode: "cors",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(json)
    })

    const response = (await fetch(request)).json()
    return response
}

function signUpButton(element) {
    const dataArr = element.parentElement.parentElement.getElementsByTagName("input")
    const data = {
        StudentName: dataArr[0].value, 
        StudentSurname: dataArr[1].value, 
        Username: dataArr[2].value, 
        Password: dataArr[3].value,
    }
    const response = signUp("http://localhost:3000/signup", data)
    response.then(value => console.log(value))
}

async function signUp(url, json) {
    const request = new Request(url, {
        method: "POST",
        mode: "cors",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(json)
    })

    const response = (await fetch(request)).json()
    return response
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
        method: "POST",
        mode: "cors",
        headers: {
            "Content-Type": "application/json"
        },
        body: data
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

let remained = 0

function getRemain() {
    const response = remainRequest("http://localhost:3000/getRemain")
    response.then(value => {
        remained = value.Payload.Message
        const promise = new Promise((resolve, reject) => {
            access()
            console.log("hello from first")
            console.log(document.querySelectorAll(".uni-wrapper").length)
            resolve()
        }).then(() => {
            console.log("hello from second")
            access()
        })
        
    })
}

getRemain()

async function remainRequest(url) {
    const request = new Request(url, {
        method: "GET",
        mode: "cors",
    })

    const response = await fetch(request)

    return await response.json()

}


async function makeJsonUniversity() {
    console.log("start function")
    flag = false
    let universityCount = document.querySelectorAll(".uni-wrapper").length
    console.log(universityCount)
    if (remained <= universityCount) {
        
        return
    }
    if (document.documentElement.getBoundingClientRect().bottom < document.documentElement.clientHeight + 10) {
        const universityOrder = universityCount
        const requestJson = `{"order": ${universityOrder}}`
        const response = getUniversity("http://localhost:3000/getUniversity", requestJson)
        let firstUniversity = []
        let secondUniversity = []
        await response.then(value => {
            firstUniversity = value.Payload.FirstUni.split("|");
            secondUniversity = value.Payload.SecondUni.split("|");
            remained = value.Payload.Remain.Message
            console.log(remained)
        })
        
        await makeUniversityElem(firstUniversity[0], firstUniversity[1], firstUniversity[2], secondUniversity[0], secondUniversity[1], secondUniversity[2])
        // makeUniversityElem(response.name, response.description, response.imagePath)
        // if (!response.left) {
        //     return
        // }
        // remained = await parseInt(response.left)
    }
    console.log("end function")
    flag = true
}

function access() {
    if (flag) {
        makeJsonUniversity()
    }
}

function makeUniversityElem(firstName, firstDescription, firstImg, secondName, secondDescription, secondImg) {
    const htmlElement = `
        <div class="uni-wrapper" id="0">
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
        <div class="uni-wrapper" id="1">
            <div class="info-container">
                <img src="${secondImg}" alt="x">
                <div class="text-container">
                    <h2>${secondName}</h2>
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
    const liElement = document.createElement("li")
    liElement.innerHTML = htmlElement
    document.querySelector("ul.ul").append(liElement)
    console.log("stop pushing elements")
}



window.onscroll=access

