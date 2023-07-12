// let hamburgerIsOpen = false
// function openHamburger() {
//     let hamburgerContainer = document.getElementById("hamburger-container")
//     if (!hamburgerIsOpen) {
//         hamburgerContainer.style.display = "flex"
//         hamburgerIsOpen = true;
//     } else {
//         hamburgerContainer.style.display = "none"
//         hamburgerIsOpen = false;
//     }
// }

let hamburgerIsOpen = false
function openHamburger() {
    let hamburgerContainer = document.getElementById("hamburger-menu-container")
    if (!hamburgerIsOpen) {
        // hamburgerContainer.style.display = "flex"

        hamburgerContainer.classList.toggle("showContainerMenu");
        hamburgerIsOpen = true;
    } else {
        // hamburgerContainer.style.display = "none"
        // hamburgerContainer.style.transform = "translateY(-150%)";
        hamburgerContainer.classList.toggle("showContainerMenu");

        hamburgerIsOpen = false;
    }
}

