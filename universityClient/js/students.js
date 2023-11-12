


async function getRecords() {
    const request = new Request("http://localhost:3000/records", {
        method: "GET",
        mode: "cors",
    })

    const response = await (await fetch(request)).json()

    const recordsArray = response.Payload

    console.log(recordsArray)

    const len = recordsArray.length
    const table = document.querySelector("table")
    for (let i = 0; i < len; i++) {
        const newHtml = `
            
            <td>${recordsArray[i].Student_name}</td>
            <td>${recordsArray[i].Student_surname}</td>
            <td>${recordsArray[i].Uni_name}</td>
            <td>${recordsArray[i].Student_points}</td>
            
        `
        const newTr = document.createElement("tr")
        newTr.innerHTML = newHtml
        table.append(newTr)
    }

}

getRecords()