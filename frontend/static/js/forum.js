
import { cards } from "./card.js";
import { alertPopup } from "./alert.js";
import { likes } from "./likescomment.js";
import { fetchupdateCard } from "./createcomment.js";
let content = [];
export async function fetchData(page = 1) {
  const response = await fetch(`/api/home?page=${page}`, {
    method: "GET",
  });

  if (response.ok) {
    let path = window.location.pathname;
    if (path !== "/profile") {
      let data = await response.json();
      let user_info = document.querySelector(".main");
      content = cards(data.posts, user_info);
      
    }
  } else if (response.status === 400) {
    const data = await response.json();
    // alertPopup(data);
  } else if (response.status === 401) {
    history.pushState(null, "", "/login");
    //loadPage();
  }
} 