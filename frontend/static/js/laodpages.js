// import { log } from "console";
import { ProfileNav } from "./categories.js";
import { Inf } from "./checklogin.js";
import { fetchCommat, GetComments } from "./comment.js";
import { leftside } from "./component.js";
// import { fetchData } from "./forum.js";
import { login, register } from "./globa.js";
import { Login } from "./login.js";
import { classes } from "./popup.js";
import { Register } from "./register.js";
import { setupWs } from "./ws.js";

const section = document.querySelector("section");

document.addEventListener("DOMContentLoaded", async () => {
  const res = await fetch("/api/isLogged");
  if (res.ok) setupWs();
});

function loadPage() {
  const path = window.location.pathname.slice(1);

  switch (path) {
    case "login":
      section.classList.add("sectionLogin");
      section.innerHTML = login;
      Login();
      break;
    case "register":
      section.classList.add("sectionLogin");
      section.innerHTML = register;
      Register();
      break;
    case "":
    case "home":
      section.classList.remove("sectionLogin");
      leftside();
      classes();
      Inf();
      break;
    case "categories":
      section.classList.remove("sectionLogin");
      leftside();
      ProfileNav();
      classes();
      break;
    case "comment":
      leftside();
      classes();
      fetchCommat();
      GetComments();
      // Inf();
      break;
    default:
      section.innerHTML = "<h1>Page Not Found</h1>";
      break;
  }
}

loadPage();
let lastPath = window.location.pathname;

setInterval(() => {
  if (lastPath !== window.location.pathname) {
    lastPath = window.location.pathname;
    loadPage();
  }
}, 100);

export { loadPage };
