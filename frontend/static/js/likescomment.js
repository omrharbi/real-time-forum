// import { alertPopup } from "./alert.js";
// import { fetchupdateCard } from "./createcomment.js";

export async function likes(likeElements ,alldislike,card_id ) {
  const storedData = localStorage.getItem("data");
    const response = await fetch("/api/likescheked", {
      method: "POST",
      body: JSON.stringify({ card_id: +card_id }),
    });
    if (response.ok) {
      let data = await response.json();
      
       data.forEach(d => {
        if ( d.UserLiked) {
          likeElements.classList.add("clicked");
          return
        } else if (d.UserDisliked) {
          alldislike.classList.add("clicked_disliked");
        }
      })
    }   
} 

export async function addLikes(card_id, liked ) {  
  try {
    if (document.cookie != "") {
      let response = await fetch("/api/addlike", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify({
          card_id: +card_id, 
          is_liked: liked
        }),
      });
      if(response.ok){
        console.log(response);
        
        // alertPopup("success", "You have successfully liked the card");  
      }else if (response.status === 400) {
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
