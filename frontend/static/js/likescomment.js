import { alertPopup } from "./alert.js";
import { fetchupdateCard } from "./createcomment.js";

export function likes(likeElements) {
  const storedData = localStorage.getItem("data");
  const parsedData = JSON.parse(storedData);
  likeElements.forEach(async (click) => {
    let card_id = click.getAttribute("data-id_card");
    let data_like = click.getAttribute("data-like");
    const response = await fetch("/api/likescheked", {
      method: "POST",
      body: JSON.stringify({ card_id: +card_id }),
    });
    if (response.ok) {
      let data = await response.json();
       data.forEach(d => {
        if (parsedData.id === d.Id_user && d.UserLiked && data_like === "like") {
          click.classList.add("clicked");
          return
        } else if (parsedData.id === d.Id_user && d.UserDisliked && data_like === "Dislikes") {
          click.classList.add("clicked_disliked");
        }
      })
    }
  });
   
} 

export async function addLikes(card_id, liked ) {
  try {
    if (document.cookie != "") {
      let response = await fetch("/api/like", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify({
          card_id: +card_id, 
          is_liked: +liked
        }),
      });
      if(response.ok){
        alertPopup("success", "You have successfully liked the card");  
      }
      if (response.status === 400) {
        const data = await response.json();
        console.log();
        
       // alertPopup(data);
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
        // alertPopup(data);
      }
    }
  } catch (error) { }
}
