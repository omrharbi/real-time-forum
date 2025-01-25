// import { log } from "console";
import { ProfileNav } from "./categories.js";
import { checklogin, Inf } from "./checklogin.js";
import { fetchCommat, GetComments } from "./comment.js";
import {
  fetchConnectedUsers,
  messages,
  setupWs,
  user_item,
} from "./compenetChat.js";
import { leftside } from "./component.js";
// import { fetchData } from "./forum.js";
import { login, register } from "./globa.js";
import { Login } from "./login.js";
import { logout } from "./logout.js";
import { classes } from "./popup.js";
import { fetchData, Profile, profileInfo } from "./profile.js";
import { Register } from "./register.js";

const section = document.querySelector("section");

document.addEventListener("DOMContentLoaded", async () => {
  const res = await fetch("/api/isLogged");
  if (res.ok) setupWs();
});

window.addEventListener("popstate", (e) => {
  loadPage();
});

async function loadPage() {
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
    case "chat":
      await checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      fetchConnectedUsers();
      messages();
      user_item();
      break;
    case "":
    case "home":
      await checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      classes();
      Inf();
      break;
    case "categories":
      await checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      ProfileNav();
      classes();
      Inf();

      break;
    case "comment":
      await checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      classes();
      fetchCommat();
      GetComments();
      // Inf();
      break;
    case "profile":
      await checklogin();
      section.classList.remove("sectionLogin");
      leftside();
      classes();
      fetchData("posts");
      profileInfo();
      break;
    case "settings":
      await checklogin();
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
