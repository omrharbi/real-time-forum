import { loadPage } from "../laodpages.js";
import { addUser, sendMessage } from "./compenetChat.js";
 



export function updateUserList(message) {
  const data = JSON.parse(localStorage.getItem("data"))
  if (data.id == message.online_users) {
    return
  }
  let id = document.getElementById(message.online_users);
  if (id) {
    let status = id.querySelector(".status");
    if (message.type === "online") {
      status.style.background = "green";
    } else {
      status.style.background = "red";
    }
  } else {
    const user = addUser(message.online_users, message.userName, message.type);
    document.querySelector(".aside-right").append(user);
  }
}

export function SetUserUp(message) {
  const userlist = document.querySelector(".aside-right");
  let useritem = document.getElementById(message.sender);
  if (useritem) {
    useritem.remove();
  } else {
    useritem = addUser(message.online_users, message.username, true);
  }
  userlist.prepend(useritem);
}
