async function signInButton(element) {
    const dataArr = element.parentElement.getElementsByTagName("input")
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
    if (response.Status == "Ok") {
        document.querySelector(".login-form").style.display = 'none'
        const newDiv = document.createElement("div")
        newDiv.classList.add("action-wrapper")
        const newHtml = `
            <ul class="actions">
                <li class="delete-students">
                    <button class="del-st">Удалить</button>
                    <div class="students-text">Удалить записи о студентах</div>
                </li>
                <li class="delete-records">
                    <button class="del-rec">Удалить</button>
                    <div class="records-text">Удалить записи с документами</div>
                </li>
                <li class="stop-submission">
                    <button class="stop-sub">Остановить</button>
                    <div class="sub-text">Остановить подачу документов</div>
                </li>
            </ul>
        `

        newDiv.innerHTML = newHtml
        document.body.append(newDiv)

    }

}