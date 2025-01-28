import { alertPopup } from "./alert.js";
import { setupWs } from "./chat/compenetChat.js";
import { loadPage } from "./laodpages.js";
export async function Register() {
  const res = await fetch("/api/isLogged")
  if (res.ok) {
    history.pushState( null, "" , "/")
    loadPage()
    return
  }
  document.querySelector("#login").addEventListener("click", () => {
    history.pushState(null, "", "/login");
    loadPage();
  });
  let register = document.querySelector("#form-submit");
  register.addEventListener("submit", async (e) => {
    e.preventDefault();
    let firstname = document.getElementById("firstname").value;
    let lastname = document.getElementById("lastname").value;
    let emailRegister = document.getElementById("emailRegister").value;
    let passwordRegister = document.getElementById("passwordRegister").value;
    let username = document.getElementById("username").value;
    let age = document.getElementById("age").value;
    const selectedGender = document.querySelector(
      'input[name="gender"]:checked'
    ).value;
    console.log(age);
    const response = await fetch("/api/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },

      body: JSON.stringify({
        firstname: firstname,
        lastname: lastname,
        email: emailRegister,
        password: passwordRegister,
        username: username,
        age: +age,
        gender: selectedGender,
      }),
    });

    if (response.ok) {
      const data = await response.json();
      console.log("Success:", data);
      window.alert("You have register successfuly");
      const userData = {
        firstname: data.message.firstname,
        lastname: data.message.lastname,
        email: data.message.email,
      };
      localStorage.setItem("data", JSON.stringify(userData));
      history.pushState(null, "", "/");
      setupWs();
      loadPage();
    } else if (response.status === 409 || response.status === 400) {
      const data = await response.json();
      alertPopup(data);
    } else {
      const errorData = await response.json();
      console.error("Error:", errorData);
      alert(`Error: ${errorData.message || "Request failed"}`);
    }
  });
}
