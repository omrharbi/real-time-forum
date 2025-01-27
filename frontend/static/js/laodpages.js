// import { log } from "console";
import { ProfileNav } from "./categories.js";
import { checklogin, Inf } from "./checklogin.js";
import { fetchCommat, GetComments } from "./comment.js";
import {
  addStyle,
  messages,
  setupWs,
  // user_item,
} from "./chat/compenetChat.js";
import { leftside } from "./component.js";
import { about, login, register } from "./globa.js";
import { Login } from "./login.js";
import { logout } from "./logout.js";
import { classes } from "./popup.js";
import { fetchData, Profile, profileInfo } from "./profile.js";
import { Register } from "./register.js";
import { fetchConnectedUsers } from "./chat/displyMessage.js";

const section = document.querySelector("section");

document.addEventListener("DOMContentLoaded", async () => {
  const res = await fetch("/api/isLogged");
  if (res.ok) setupWs();
});

window.addEventListener("popstate", (e) => {
  loadPage();
});

function loadPage() {
  const path = window.location.pathname.slice(1);
  switch (path) {
    case "login":
      document.head.querySelector("title").innerText = path;
      section.classList.add("sectionLogin");
      section.innerHTML = login;
      Login();
      break;
    case "register":
      document.head.querySelector("title").innerText = path;
      section.classList.add("sectionLogin");
      section.innerHTML = register;
      Register();
      break;
    case "chat":
      document.head.querySelector("title").innerText = path;
      checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      fetchConnectedUsers();
      messages();
      addStyle();
      break;
    case "":
    case "home":
      document.head.querySelector("title").innerText = "home";
      checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      fetchConnectedUsers();
      classes();
      Inf();
      break;
    case "comment":
      document.head.querySelector("title").innerText = path;

      checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      fetchConnectedUsers();
      classes();
      fetchCommat();
      GetComments();
      // Inf();
      break;
    case "profile":
      document.head.querySelector("title").innerText = path;

      checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      fetchConnectedUsers();
      classes();
      fetchData("posts");
      profileInfo();
      break;
    case "settings":
      document.head.querySelector("title").innerText = path;
      checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      classes();
      fetchConnectedUsers();
      logout();
      break;
    default:
      section.innerHTML = "<h1>Page Not Found</h1>";
      break;
  }
}

loadPage();
export { loadPage };
