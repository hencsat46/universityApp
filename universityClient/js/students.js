


async function getRecords() {
    const request = new Request("http://localhost:3000/records", {
        method: "GET",
        mode: "cors",
    })

    const response = await (await fetch(request)).json()

    const recordsArray = (await response).Payload.Message

    console.log(recordsArray)

    const len = recordsArray.length
    const table = document.querySelector("table")
    for (let i = 0; i < len; i++) {
        const newHtml = `
            
            <td>${recordsArray[i][0]}</td>
            <td>${recordsArray[i][1]}</td>
            <td>${recordsArray[i][2]}</td>
            <td>${recordsArray[i][3]}</td>
            
        `
        const newTr = document.createElement("tr")
        newTr.innerHTML = newHtml
        table.append(newTr)
    }

}

getRecords()