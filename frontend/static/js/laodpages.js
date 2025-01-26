// import { log } from "console";
import { ProfileNav } from "./categories.js";
import { checklogin, Inf } from "./checklogin.js";
import { fetchCommat, GetComments } from "./comment.js";
import {
  addStyle,
  messages,
  setupWs,
 } from "./chat/compenetChat.js";
import { leftside } from "./component.js";
 import { login, register } from "./globa.js";
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
  console.log(path);

  switch (path) {
    case "login":
      // document.head.title = "login";
      section.classList.add("sectionLogin");
      section.innerHTML = login;
      Login();
      break;
    case "register":
      section.classList.add("sectionLogin");
      section.innerHTML = register;
      Register();
      break;
    case "chat":
       checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      messages();
      fetchConnectedUsers();
      addStyle();
      break;
    case "":
    case "home":
      checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      classes();
      Inf();
      break;
    case "categories":
       checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      ProfileNav();
      classes();
      Inf();

      break;
    case "comment":
       checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      classes();
      fetchCommat();
      GetComments();
      // Inf();
      break;
    case "profile":
       checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      classes();
      fetchData("posts");
      profileInfo();
      break;
    case "settings":
       checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      classes();
      logout();
      break;
    default:
      section.innerHTML = "<h1>Page Not Found</h1>";
      break;
  }
}

loadPage();
export { loadPage };
