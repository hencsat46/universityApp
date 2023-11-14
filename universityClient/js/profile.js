async function getProfile() {
    const request = new Request("http://localhost:3000/profile", {
        method: "GET",
        mode: "cors",
        headers: {
            "Token": getToken(),
        }
    })

    const response = await (await fetch(request)).json()

    console.log(response)

    const newHtml = `

        <div class="user-bio">
            <div class="user-photo">
                <img src="../img/usernameIcon.jpg" alt="x">
            </div>
            <div class="user-data">
                <div class="user-name">${response.Payload.Student_name}</div>
                <div class="user-surname">${response.Payload.Student_surname}</div>
                <div class="user-university">
                    ${response.Payload.Uni_name}
                </div>
                <button onclick="quit()">Выйти</button>
            </div>
        </div>
    
    `

    document.querySelector(".wrapper").innerHTML = newHtml
}

function quit() {
    document.cookie = "Token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;"
    window.location.href = "/";
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