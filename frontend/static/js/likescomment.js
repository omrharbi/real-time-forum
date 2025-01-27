import { alertPopup } from "./alert.js";

export function likes(likeElements) {
  // if (document.cookie != "") {
  // console.log(likeElements);

  likeElements.forEach(async (click) => {


    let card_id = click.getAttribute("data-id_card");
    let like = click.getAttribute("data-id_card");



    const response = await fetch("/api/likescheked", {
      method: "POST",
      body: JSON.stringify({ card_id: +card_id }),
    });
    if (response.ok) {
      let data = await response.json();
      data.forEach(d => {
        // if (like === d.Uuid&& d.UserLiked) {
        //   console.log(click,"is liked ");

        // }
        console.log(like, "uuid", d.Uuid, d.UserLiked);

      })

      // console.log(data);
      // data.forEach((el) => {
      //   let tokens = document.cookie.split("token=");
      //   if (el.Uuid === tokens[1]) {
      //     localStorage.setItem("user_login", el.User_id);
      //     if (el.UserLiked && like === "like") {
      //       click.classList.add("clicked");
      //       click.setAttribute("data-liked", "true");
      //     } else if (el.UserDisliked && like === "Dislikes") {
      //       click.classList.add("clicked_disliked");
      //       click.setAttribute("data-liked", "true");
      //     }
      //   }
      // });
    }
  });
}
// }

export async function addLikes(card_id, liked, lik, dislk, click) {
  try {
    if (document.cookie != "") {
      let response = await fetch("/api/like", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify({
          is_liked: +liked,
          card_id: +card_id,
          UserLiked: lik,
          Userdisliked: dislk,
        }),
      });
      if (response.status === 400) {
        const data = await response.json();
        alertPopup(data);
      }
    }
  } catch (error) {
    console.log(error);
  }
}

export async function deletLikes(card_id) {
  try {
    if (document.cookie != "") {
      let response = await fetch("/api/deleted", {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify({ card_id: +card_id }),
      });

      if (response.status === 400) {
        const data = await response.json();
        alertPopup(data);
      }
    }
  } catch (error) { }
}
