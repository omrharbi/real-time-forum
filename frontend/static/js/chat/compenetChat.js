import { debounce } from "../checklogin.js";
import { loadPage } from "../laodpages.js";
import { getCookie, logout } from "../logout.js";
import { SetUserUp, updateUserList } from "./create_user.js";
import { displayMessage, GetMessage, getMessage } from "./displyMessage.js";

let ws;
export function setupWs() {
  ws = new WebSocket(`ws://${window.location.host}/ws`);
  ws.onopen = () => {
    console.log("is connected");
  };

  ws.onmessage = async (event) => {
    const message = JSON.parse(event.data);
    console.log(message);

    switch (message.type) {
      case "online":
        updateUserList(message);
        break;
      case "broadcast":
        const query = new URLSearchParams(window.location.search);
        if (window.location.pathname === "/chat") {
          if (query.get("receiver") == message.sender) {
            displayMessage(
              message.username,
              message.createAt,
              message.content,
              false
            );
            addtoOfset++;
          } else {
            document.getElementById("notify").play();
            showPopup(`you have message from ${message.username}`);
          }
          SetUserUp(message);
        } else {
          document.getElementById("notify").play();
          showPopup(`you have message from ${message.username}`);
        }
        break;
      case "typing":
        showTypingNotification(message.userId);
        break;
      case "offline":
        updateUserList(message);
        break;
      case "reload":
        window.location.reload();
        break;
      default:
        console.warn("Unhandled message type:", message.type);

    }
  };

  ws.onclose = () => {
    console.log("WebSocket connection closed.");
  };

  ws.onerror = (error) => {
    console.error("WebSocket error:", error);
  };
}

export const chat = /*html*/ `
    
            <input type="text" id="messageInput" placeholder="Type your message here..." />
            <div id="sendButton" >
               <svg   xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-send" viewBox="0 0 16 16">
                  <path d="M15.854.146a.5.5 0 0 1 .11.54l-5.819 14.547a.75.75 0 0 1-1.329.124l-3.178-4.995L.643 7.184a.75.75 0 0 1 .124-1.33L15.314.037a.5.5 0 0 1 .54.11ZM6.636 10.07l2.761 4.338L14.13 2.576zm6.787-8.201L1.591 6.602l4.339 2.76z"/>
               </svg>
             </div>
    
`;

export function addUser(userId, userName, status) {
  const userList = document.querySelector(".aside-right");
  const userItem = document.createElement("li");
  userItem.className = "user-item";
  userItem.id = userId;
  userItem.dataset.id = userName;

  const userIcon = document.createElement("div");
  userIcon.className = "user-icon";
  userIcon.textContent = userName[0].toUpperCase();

  const userNameDiv = document.createElement("div");
  userNameDiv.className = "user-name";

  const type = document.createElement("div");
  type.className = "type";
  userNameDiv.textContent = userName;

  const statusDot = document.createElement("span");
  statusDot.className = "status";
  userNameDiv.appendChild(type)
  userItem.append(userIcon, userNameDiv, statusDot);
  userList.appendChild(userItem);
  userItem.addEventListener("click", () => {
    let url = `chat?receiver=${userId}`;
    history.pushState(null, "", url);
    let log = document.querySelector(".chat");
    if (log) {
      log.innerHTML = "";
    }
    addtoOfset = 0;
    loadPage();
    // sendMessage();
  });
  statusDot.style.background = status === "online" ? "green" : "red";
  return userItem;
}

export function messages() {
  const chat = document.querySelector(".content_post");
  chat.style.height = "100%";
  chat.innerHTML += /*html*/ `
      
            <div class="chat-message chat-container">
                    <div class="message">
                         <div class="chat"></div>
                         <div class="chat-input"></div>
                    </div>
            </div>
    `;
  const query = new URLSearchParams(window.location.search);
  if (query.get("receiver")) {
    addtoOfset = 0;
    GetMessage(query.get("receiver"));
    let chat = document.querySelector(".chat");
    chat.scrollBy(0, chat.scrollHeight);
    sendMessage();
  } else {
    let chat = document.querySelector(".chat");
    chat.className = "chat welcome";
    chat.textContent = "WELCOME TO CHAT";
  }
}

export let addtoOfset = 0;

export function sendMessage() {
  const storedData = localStorage.getItem("data");
  const parsedData = JSON.parse(storedData);
  let token = getCookie("token");


  const chat = document.querySelector(".content_post");
  let message = chat.querySelector("#messageInput");
  let sendButton = chat.querySelector("#sendButton");
  message.addEventListener("keypress", (e) => {
    if (!token) {
      logout()
    }
    if (e.key === "Enter") {
      e.preventDefault();
      sendButton.click();
    }
  });
  sendButton.addEventListener(
    "click",
    debounce(() => {

      if (!token) {
        logout()
      }
      let receiver = new URLSearchParams(location.search).get("receiver");
      const messages = message.value.trim();
      if (messages) {
        displayMessage(parsedData.firstname, new Date(), messages, true);
        addtoOfset++;
        SetUserUp({ sender: receiver });
        ws.send(
          JSON.stringify({
            type: "broadcast",
            content: messages,
            sender: parsedData.id,
            receiver: +receiver,
            createAt: new Date(),
          })
        );
        message.value = "";
      }
    }, 200)
  );
}
export function showPopup(message) {
  const popup = document.getElementById("message-popup");
  const popupMessage = document.getElementById("popup-message");

  popupMessage.textContent = message;

  popup.style.display = "block";
  setTimeout(() => {
    popup.style.display = "none";
  }, 5000);
}

export function addStyle() {
  let style = document.createElement("link");
  style.rel = "stylesheet";
  style.href = "../static/css/chat.css";
  document.head.appendChild(style);
}
