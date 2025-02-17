import { loadPage } from "./laodpages.js";

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
      div.classList.add("active");
      history.pushState(null, "", `/${id}`);
      loadPage();
    } else {
      e.preventDefault();

      console.log("Already on the Home page, no need to reload.");
    }
  });
}

const ids = ["chat", "home", "categories", "profile", "settings"];

function Change() {
  for (let idl of ids) {
    navigate(idl);
  }
}

export { navigate, Change };
