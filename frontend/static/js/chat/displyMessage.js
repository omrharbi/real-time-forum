import { getTimeDifferenceInHours } from "../card.js";
import { chat } from "./compenetChat.js";
// import { user_item } from "./compenetChat.js";
import { addUser } from "./create_user.js";

export function displayMessage(
  sender,
  createAt,
  content,
  isOwnMessage = false
) {
  let log = document.querySelector(".chat");
  const parent = document.createElement("div");
  const messageUser = document.createElement("div");
  const message_content = document.createElement("div");
  const time = document.createElement("div");
  const userIcon = document.createElement("div");
  const row = document.createElement("div");

  messageUser.className = "messages";
  message_content.className = "message-content";
  parent.className = "parent";

  userIcon.className = "user-icon message-icon";
  userIcon.textContent = sender[0].toUpperCase();

  message_content.textContent = `${content}`;
  time.textContent = getTimeDifferenceInHours(createAt);
  if (isOwnMessage) {
    messageUser.classList = "messages sander";
    time.className = "time sander";
    row.className = "row sander";
  } else {
    messageUser.className = "messages resiver";
    time.className = "time resiver";
    row.className = "row resiver";
  }
  messageUser.append(message_content, time);
  row.append(userIcon, messageUser);
  parent.appendChild(row);
  log.appendChild(parent);
  //log.scrollTop = log.scrollHeight;
}

export async function getMessage(receiver) {
  const log = document.querySelector(".chat-input");
  console.log(log);
  log.innerHTML = chat;
  const storedData = localStorage.getItem("data");
  let sendButton = document.getElementById("sendButton");
  let messageInput = document.getElementById("messageInput");
  sendButton.style.display = "block";
  messageInput.style.display = "block";
  const parsedData = JSON.parse(storedData);
  const response = await fetch("/api/messages", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      receiver: +receiver,
    }),
  });
  if (response) {
    let data = await response.json();
    if (data) {
      let isOwen;
      data.forEach((d) => {
        if (parsedData.id === d.sender) {
          isOwen = true;
        } else {
          isOwen = false;
        }

        displayMessage(d.username, d.createAt, d.content, isOwen);
      });
    }
  } else {
    console.log("error");
  }
}

export async function fetchConnectedUsers() {
  const response = await fetch("/api/connected");
  if (response.ok) {
    const userList = document.getElementById("userList");
    userList.innerHTML = "";
    const users = await response.json();
    users.forEach((user) => {
      console.log(user);
      addUser(user.id, user.username, user.status);
    });
  } else {
    console.error("Failed to fetch connected users:", response.status);
  }
}
