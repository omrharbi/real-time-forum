import { InitialComment, likes } from "./createcomment.js";
// import { checklogin } from "./checklogin.js";

import { alertPopup } from "./alert.js";
import { addLikes, deletLikes } from "./likescomment.js";
// await checklogin()

async function fetchCommat() {
  const urlParams = new URLSearchParams(window.location.search);
  const cardData = urlParams.get("card_id");
  let fullname = document.querySelector(".full-name");
  let content = document.querySelector(".content");
  let time = document.querySelector(".time");
  let username = document.querySelector(".username");
  let cards = document.querySelector("#likes");
  let disliked = document.querySelector("#dislikes")
  let is_liked = document.querySelector("#is_liked");
  let is_Dislikes = document.querySelector("#is_Dislikes");
  let comments = document.querySelector(".comments");
  let data = "";
  let path = window.location.pathname;

  if (path !== "/comment") {
    return "";
  } else {
    const response = await fetch(`/api/card?id=${cardData}`, {
      method: "GET",
    });
    if (response.ok) {
      data = await response.json();
      fullname.textContent = data.lastName + " " + data.firstName;
      content.textContent = data.content;
      username.textContent = data.lastName;
      is_liked.textContent = data.likes;
      is_Dislikes.textContent = data.dislikes;
      comments.textContent = data.comments;
      let catgory = document.querySelector(".catgory");
      let c = data.categories.split(",");
      console.log(data);
      
      if (data.categories==="") {
        catgory.style.display = "none"
      }
      c.forEach((element) => {
        console.log(element);

        let CreatCate = document.createElement("span");
        CreatCate.className = "category-item categories";
        CreatCate.textContent = element;
        catgory.appendChild(CreatCate);
      });


      likes(cards, disliked, data.id)
      cards.addEventListener("click", () => {
        if (cards.classList.contains("clicked")) {
          deletLikes(data.id);
          cards.classList.remove("clicked");
          data.likes--;
        } else {
          addLikes(data.id, true)
          if (disliked.classList.contains("clicked_disliked")) {
            disliked.classList.remove("clicked_disliked");
            data.dislikes--;
          }
          data.likes++;
          cards.classList.add("clicked");
        }
        is_Dislikes.innerHTML = data.dislikes;
        is_liked.innerHTML = data.likes;
      });

      disliked.addEventListener("click", () => {
        if (disliked.classList.contains("clicked_disliked")) {
          deletLikes(data.id);
          disliked.classList.remove("clicked_disliked");
          data.dislikes--;
        } else {
          addLikes(data.id, false)
          if (cards.classList.contains("clicked")) {
            cards.classList.remove("clicked");
            data.likes--;
          }
          data.dislikes++;
          disliked.classList.add("clicked_disliked");
        }
        is_Dislikes.innerHTML = data.dislikes;
        is_liked.innerHTML = data.likes;
      });
    } else if (response.status === 409 || response.status === 400) {
      const data = await response.json();
      alertPopup(data);
    }
  }
}
async function GetComments() {
  const urlParams = new URLSearchParams(window.location.search);
  const cardData = urlParams.get("card_id");
  let path = window.location.pathname;
  if (path !== "/comment") {
    return "";
  } else {
    const response = await fetch(`/api/comment?target_id=${cardData}`, {
      method: "GET",
    });

    if (response.ok) {
      let textResponse = await response.text();
      if (textResponse.trim() === "") {
        console.log("Empty response body");
        return;
      }
      let datacomment = JSON.parse(textResponse); // Manually parse JSON
      let comments = document.querySelector(".allcomment");
      comments.innerHTML = "";
      await InitialComment(datacomment, comments);
      localStorage.setItem("card_id", cardData);
    } else if (response.status === 400) {
      const data = await response.json();
      alertPopup(data);
    } else {
      console.log("err");
    }
  }
}
// await GetComments()
export { GetComments, fetchCommat };
