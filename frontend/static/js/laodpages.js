import { ProfileNav } from "./categories.js";
import { leftside } from "./component.js";
import { login } from "./globa.js";
import { Login } from "./login.js";
import { classes } from "./popup.js";
import { addStyle, fetchConnectedUsers, messages, setupWs, user_item } from "./compenetChat.js";

const section = document.querySelector("section");

document.addEventListener("DOMContentLoaded", () => {
  setupWs();
});

function loadPage() {
  const path = window.location.pathname.slice(1);
  switch (path) {
    case "login":
      section.classList.add("sectionLogin");
      section.innerHTML = login;
      Login();
      break;

    case "home":
      section.classList.remove("sectionLogin");
      leftside();
      classes();
      break;
    case "chat":
      section.classList.remove("sectionLogin");
      addStyle()
      leftside();
      messages()
      fetchConnectedUsers()
      user_item()
      break;
    case "categories":
      section.classList.remove("sectionLogin");
      leftside();
      ProfileNav();
      classes();
      break;
    default:
      section.innerHTML = "<h1>Page Not Found</h1>";
      break;
  }
}

loadPage();

export { loadPage };
