import { alertPopup } from "./alert.js";
import { loadPage } from "./laodpages.js";

export async function logout() {
  let Useruuid = getCookie("token");
  const LogoutItem = document.querySelector(".signOut");

  if (LogoutItem) {
    LogoutItem.addEventListener("click", async () => {
      const response = await fetch("http://localhost:8080/api/logout", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ uuid: Useruuid }),
      });
      history.pushState(null, "", "/login");
      loadPage();
    });
  } else {
    console.error("Logout button not found");
  }
}

function getCookie(name) {
  const cookies = document.cookie.split("; ");
  for (let i = 0; i < cookies.length; i++) {
    const [key, value] = cookies[i].split("=");
    if (key === name) {
      return value;
    }
  }
  return null;
}
