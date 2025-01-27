import { debounce } from "../checklogin.js";
import { loadPage } from "../laodpages.js";
import { SetUserUp, updateUserList } from "./create_user.js";
import { displayMessage, GetMessage, getMessage } from "./displyMessage.js";

let ws;
export function setupWs() {
  console.log("from login");

  ws = new WebSocket(`ws://${window.location.host}/ws`);
  ws.onopen = () => {
    console.log("is connected");
  };

  ws.onmessage = async (event) => {
    const message = JSON.parse(event.data);
    switch (message.type) {
      case "online":
        if (window.location.pathname === "/chat") updateUserList(message);
        break;
      case "broadcast":
        const query = new URLSearchParams(window.location.search);
        if (window.location.pathname === "/chat") {
          console.log(message);

          if (query.get("receiver") == message.sender) {
            displayMessage(
              message.username,
              message.createAt,
              message.content,
              false
            );
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
        if (window.location.pathname === "/chat") updateUserList(message);
        break;
      default:
        console.warn("Unhandled message type:", message.type);
    }
  };

  ws.onclose = () => {
    console.log("WebSocket connection closed.");
    history.pushState("", "", "/login");
    loadPage();
  };

  ws.onerror = (error) => {
    console.error("WebSocket error:", error);
  };
}

export const chat = /*html*/ `
    
            <input type="text" id="messageInput" placeholder="Type your message here..." />
            <button id="sendButton">Send</button>
    
`;

export function messages() {
  const chat = document.querySelector(".content_post");
  chat.style.height = "100%";
  chat.innerHTML += /*html*/ `
      
            <div class="chat-message chat-container">
                    <div class="users">
                        <h1 class="user-online">My Friends  </h1>
                        <ul class="user-list" id="userList">
                            <!-- User items will be added dynamically -->
                        </ul>
                    </div>
                    <div class="message">
                         <div class="chat"></div>
                         <div class="chat-input"></div>
                    </div>
            </div>
    `;
  const query = new URLSearchParams(window.location.search);
  if (query.get("receiver")) {
    GetMessage(query.get("receiver"));
    sendMessage();
  } else {
    let chat = document.querySelector(".chat");
    chat.className = "chat welcome";
    chat.textContent = "WELCOME TO CHAT";
  }
}

export function sendMessage() {
  const storedData = localStorage.getItem("data");
  const parsedData = JSON.parse(storedData);
  const chat = document.querySelector(".content_post");
  let message = chat.querySelector("#messageInput");
  let sendButton = chat.querySelector("#sendButton");
  sendButton.addEventListener(
    "click",
    debounce(() => {
      let receiver = new URLSearchParams(location.search).get("receiver");
      const messages = message.value.trim();
      if (messages) {
        displayMessage(parsedData.firstname, new Date(), messages, true);
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
