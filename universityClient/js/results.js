function createTable(response) {

    let html = ""

    if (response.Status == 200) {

        for (let i = 0; i < response.Payload.length; ++i) {
            console.log(response.Payload[i].Student_university)
            const header = `<h3 class="table-header">${response.Payload[i].Student_university}</h3>`
            let table = "<table>"
            for (let j = 0; j < response.Payload[i].Student_information.length; ++j) {
                
                let columns = ""
                columns += `<td>${response.Payload[i].Student_information[j].Student_name}</td>`
                columns += `<td>${response.Payload[i].Student_information[j].Student_surname}</td>`
                columns += `<td>${response.Payload[i].Student_information[j].Student_points}</td>`
                console.log(columns)
                table += `<tr>${columns}</tr>`
                console.log(table)
            }
            table += "</table>"
            html += header + table
            
        }
    }
    console.log(html)
    const tables = document.createElement("div")
    tables.innerHTML = html
    document.querySelector("header").after(tables)
}

async function getResults() {
    

    const request = new Request("http://localhost:3000/getresult", {
        method: "GET",
        mode: "cors"
    })

    const response = await (await fetch(request)).json()
    console.log(response)
    createTable(response)

}

getResults()