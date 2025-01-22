import { leftside } from "./component.js";

// import { checklogin } from "./checklogin.js";

// await checklogin(id);
function navigate(id) {
  const div = document.querySelector(`#${id}`);
  div.addEventListener("click", (e) => {
    const currentUrl = window.location.pathname;
    if (currentUrl !== `/${id}`) {
      for (let idl of ids) {
        document.getElementById(idl).classList.remove("active");
      }
      history.pushState(null, "", `/${id}`);
      div.classList.add("active");
    } else {
      e.preventDefault();

      console.log("Already on the Home page, no need to reload.");
    }
  });
}

const ids = ["home", "categories", "profile", "settings"];

leftside();
for (let idl of ids) {
  navigate(idl);
}

export { navigate };
