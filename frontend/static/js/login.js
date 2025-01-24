import { alertPopup } from "./alert.js";

import { checklogin } from "./checklogin.js";
import { setupWs } from "./compenetChat.js";

export async function Login() {
  checklogin();
  document.querySelector("#register").addEventListener("click", () => {
    history.pushState(null, "", "/register");
  });
  let login = document.querySelector("#login");

  login.addEventListener("submit", async (e) => {
    e.preventDefault();
    let email = document.querySelector("#email").value;
    let password = document.querySelector("#password").value;
    const response = await fetch("/api/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify({
        email: email,
        password: password,
      }),
    });

    if (response.ok) {
      const data = await response.json();
      const userData = {
        uuid: data.message.uuid,
        id: data.message.id,
        firstname: data.message.firstname,
        lastname: data.message.lastname,
        email: data.message.email,
      };
      localStorage.setItem("data", JSON.stringify(userData));
      console.log(localStorage);
      setupWs();
      history.pushState(null, "", "/");
    } else if (response.status === 400) {
      const data = await response.json();
      console.log(data);

      alertPopup(data);
    } else {
      const errorData = await response.json();
      console.error("Error:", errorData);
      alert(`Error: ${errorData.message || "Request failed"}`);
    }
  });
}
