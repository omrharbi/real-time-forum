import { ProfileNav } from "./categories.js";
import { leftside } from "./component.js";
import { login } from "./globa.js";
import { Login } from "./login.js";
import { classes } from "./popup.js";

const section = document.querySelector("section");

function loadPage() {
  const path = window.location.pathname.slice(1);

  switch (path) {
    case "login":
      section.classList.add("sectionLogin");
      section.innerHTML = login;
      Login();
      document.body.addEventListener("keydown", (e) => {
        if (e.key === "p") {
          history.pushState(null, "", "/");
          loadPage();
        }
      });
      break;

    case "":
    case "home":
      section.classList.remove("sectionLogin");
      leftside();
      classes();
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
