// function toggleNightMode() {
//     let element = document.body;
//     element.classList.toggle("night-mode");
// }
// let switchButton = (toggleNightMode);


// add key and value to local storage to keep the dark mode on when in every page
let isDarkMode = localStorage.getItem("isDarkMode");
const toggleNightMode = document.getElementById("night-mode");

const darkMode = () => {
  document.body.classList.add("night-mode"); 
  localStorage.setItem("isDarkMode", "true"); 
}; 

const lightMode = () => {
  document.body.classList.remove("night-mode");
  localStorage.setItem("isDarkMode", "false");
};

if (isDarkMode === "true") {
  darkMode();
}

toggleNightMode.addEventListener("click", () => {
  isDarkMode = localStorage.getItem("isDarkMode");
  if (isDarkMode !== "true") {
    darkMode();
  } else {
    lightMode();
  }
});
