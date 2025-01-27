import { getTimeDifferenceInHours } from "../card.js";
import { throttle } from "../checklogin.js";
import { chat } from "./compenetChat.js";
// import { user_item } from "./compenetChat.js";
import { addUser } from "./create_user.js";

export function displayMessage(
  sender,
  createAt,
  content,
  isOwnMessage = false,
  fetched
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
  console.log(fetched);
  row.append(userIcon, messageUser);

  parent.prepend(row);
  if (!fetched) {
    log.appendChild(parent);
  } else {
    log.prepend(parent);
  }
  if (!fetched) {
    log.scrollBy(0, log.scrollHeight);
  }
  //log.scrollTop = log.scrollHeight;
}

export async function getMessage(receiver, offset = 0) {
  const log = document.querySelector(".chat-input");
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
      offset: offset,
    }),
  });
  if (response) {
    let data = await response.json();
    if (data) {
      let isOwen;
      if (!data) return;
      for (let i = 0; i < data.length; i++) {
        if (parsedData.id === data[i].sender) {
          isOwen = true;
        } else {
          isOwen = false;
        }
        displayMessage(
          data[i].username,
          data[i].createAt,
          data[i].content,
          isOwen,
          true
        );
      }
    }
  } else {
    console.log("error");
  }
}

let throttledScrollHandler = null;

export function GetMessage(receiver) {
  getMessage(receiver);

  let offset = 30;
  const chat = document.querySelector(".chat");

  if (throttledScrollHandler) {
    chat.removeEventListener("scroll", throttledScrollHandler);
  }
  throttledScrollHandler = throttle(() => {
    if (chat.scrollTop === 0) {
      getMessage(receiver, offset);
      offset += 30;
      console.log(offset);
    }
  }, 200);
  chat.addEventListener("scroll", throttledScrollHandler);
}

export async function fetchConnectedUsers() {
  console.log("user");

  const response = await fetch("/api/connected");
  if (response.ok) {
    const userList = document.querySelector(".aside-right");
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
