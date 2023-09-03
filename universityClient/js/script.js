
function clickButton(elem) {
    const elemName = elem.textContent.trim();
    const element = document.querySelectorAll('.form');
    let displaySignIn = element[0].style.display
    let displaySignUp = element[1].style.display
    switch(elemName) {
        case "Войти":
            element[0].style.display = "block"
            element[1].style.display = "none"
            break;
        case "Зарегистрироваться":
            element[0].style.display = "none"
            element[1].style.display = "block"
            break;
    }
    
}

function closeForm(number) {
    const element = document.querySelectorAll('.form');
    element[number].style.display = "none"
}
