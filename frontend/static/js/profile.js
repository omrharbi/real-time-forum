import { navigate } from "./home.js";
import { likes } from "./likescomment.js";
import { cards } from "./card.js";

import { alertPopup } from "./alert.js";
import { fetchupdateCard } from "./createcomment.js";
// navigate();
let content = [];
export function Profile() {
  const profileNav = document.querySelectorAll(".profile-nav span");
  profileNav.forEach((navItem) => {
    navItem.addEventListener("click", (e) => {
      const navId = navItem.getAttribute("id");
      if (navId === undefined) {
        return;
      }
      document.querySelector(".profile").innerHTML = "";
      fetchData(navId);
      navItem.className = "active";
      profileNav.forEach((item) => {
        if (item != navItem) {
          item.className = "";
        }
      });
    });
  });
}

export function profileInfo() {
  const profileInfo = document.querySelector(".profile-info");
  let dataUser = JSON.parse(localStorage.getItem("data"));
  console.log(dataUser.firstname);
  const profileName = document.createElement("h1");
  profileName.classList.add("profile-name");
  profileName.textContent = dataUser.firstname + " " + dataUser.lastname;

  const profileHandle = document.createElement("p");
  profileHandle.classList.add("profile-handle");
  profileHandle.textContent = `🌟 ` + dataUser.email;

  profileInfo.appendChild(profileName);
  profileInfo.appendChild(profileHandle);
}

//--------------------------------------
export async function fetchData(id) {
  const response = await fetch("/api/profile/" + id, {
    method: "GET",
  });
  if (response.ok) {
    let data = await response.json();
    let user_info = document.querySelector(".profile");
    content = cards(data, user_info);
    let like = document.querySelectorAll("#likes");
    likes(like);
    fetchupdateCard()
  } else if (response.status === 409 || response.status === 400) {
    const data = await response.json();
    alertPopup(data);
  } else {
    let data = response.json();
    console.log(data);
  }
}
