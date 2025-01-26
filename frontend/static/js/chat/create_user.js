import { sendMessage } from "./compenetChat.js";
import { displayMessage, GetMessage, getMessage } from "./displyMessage.js";

export function addUser(userId, userName, status) {
  const userList = document.getElementById("userList");
  const userItem = document.createElement("li");
  userItem.className = "user-item";
  userItem.id = userId;
  userItem.dataset.id = userName;

  const userIcon = document.createElement("div");
  userIcon.className = "user-icon";
  userIcon.textContent = userName[0].toUpperCase();

  const userNameDiv = document.createElement("div");
  userNameDiv.className = "user-name";
  userNameDiv.textContent = userName;

  const statusDot = document.createElement("span");
  statusDot.className = "status";

  userItem.append(userIcon, userNameDiv, statusDot);
  userList.appendChild(userItem);
  userItem.addEventListener("click", () => {
    let url = `chat?receiver=${userId}`;
    history.pushState(null, "", url);
    let log = document.querySelector(".chat");
    if (log) {
      log.innerHTML = "";
    }
    GetMessage(userId);
    sendMessage();
  });
  statusDot.style.background = status === "online" ? "green" : "red";
  return userItem;
}

export function updateUserList(message) {
  console.log(message);

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
    document.querySelector("#userList").append(user);
  }
}

export function SetUserUp(message) {
  const userlist = document.querySelector("#userList");
  console.log(message);

  let useritem = document.getElementById(message.sender);
  if (useritem) {
    useritem.remove();
  } else {
    useritem = addUser(message.online_users, message.username, true);
  }
  userlist.prepend(useritem);
}
