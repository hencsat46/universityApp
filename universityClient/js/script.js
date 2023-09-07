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

async function makeJsonUniversity() {
    let universityCount = document.querySelectorAll(".uni-wrapper").length
    if (remained <= universityCount - 2) {
        return
    }
    
    // console.log(document.documentElement.getBoundingClientRect().bottom)
    // console.log(document.documentElement.clientHeight)

    if (document.documentElement.getBoundingClientRect().bottom < document.documentElement.clientHeight + 50) {
        const universityOrder = universityCount - 2
        const requestJson = `{"order": ${universityOrder}}`
        
        const response = await getUniversity("http://localhost:3000/getUniversity", requestJson)
        console.log(response)
        makeUniversityElem(response.name, response.description, response.imagePath)
        if (!response.left) {
            return
        }
        remained = await parseInt(response.left)
    }
}

function makeUniversityElem(name, description, img) {
    const newLiElem = document.createElement("li")
    const newUniWrapper = document.createElement("div")
    newUniWrapper.classList.add("uni-wrapper")
    const newH2Elem = document.createElement("h2")
    const newImgElem = document.createElement("img")
    const newUniText = document.createElement("uni-text")
    newUniText.classList.add("uni-text")
    const ulClass = document.querySelector("ul.ul")
    newH2Elem.innerText = name
    newImgElem.setAttribute("src", img)
    newImgElem.setAttribute("alt", "x")
    newUniText.innerText = description
    ulClass.append(newLiElem)
    newLiElem.append(newUniWrapper)
    newUniWrapper.append(newH2Elem)
    newUniWrapper.append(newImgElem)
    newUniWrapper.append(newUniText)
}

window.onscroll=makeJsonUniversity

