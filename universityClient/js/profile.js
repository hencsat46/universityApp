async function getProfile() {
    const request = new Request("http://localhost:3000/profile", {
        method: "GET",
        mode: "cors",
        headers: {
            "Token": getToken(),
        }
    })

    const response = await (await fetch(request)).json()

    const newHtml = `

        <div class="user-bio">
            <div class="user-photo">
                <img src="../img/usernameIcon.jpg" alt="x">
            </div>
            <div class="user-data">
                <div class="user-name">${response.Payload.Name}</div>
                <div class="user-surname">${response.Payload.Surname}</div>
                <div class="user-university">
                    ${response.Payload.University}
                </div>
            </div>
        </div>
    
    `

    document.querySelector(".wrapper").innerHTML = newHtml
}

getProfile()

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