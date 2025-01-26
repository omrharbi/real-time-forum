import { loadPage } from "../laodpages.js";
import { updateUserList } from "./create_user.js";
import { displayMessage, getMessage } from "./displyMessage.js";


let ws
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
                updateUserList(message)
                break;
            case "broadcast":

                displayMessage(message.sender, message.createAt, message.content, false);
                break;
            case "typing":
                showTypingNotification(message.userId);
                break;
            case "offline":
                updateUserList(message)
                break;
            default:
                console.warn("Unhandled message type:", message.type);
        }
    };

    ws.onclose = () => {
        console.log("WebSocket connection closed.");
        history.pushState("", "", "/login")
        loadPage()
    };

    ws.onerror = (error) => {
        console.error("WebSocket error:", error);
    };
}

export function messages() {
    const chat = document.querySelector(".content_post");
    chat.style.height = "100%"
    chat.innerHTML += /*html*/`
      
            <div class="chat-message chat-container">
                    <div class="users">
                        <h1 class="user-online">User Online:</h1>
                        <ul class="user-list" id="userList">
                            <!-- User items will be added dynamically -->
                        </ul>
                    </div>
                    <div class="message">
                         <div class="chat"></div>
                            <div class="chat-input">
                                <input type="text" id="messageInput" placeholder="Type your message here..." />
                                
                                <ion-icon id="sendButton" name="send"></ion-icon>
                            </div>
                    </div>
            </div>
    `;
    const query = new URLSearchParams(window.location.search);
    let sendButton = document.getElementById("sendButton")
    let messageInput = document.getElementById("messageInput")
    if (query.get("receiver")) {


        //

        getMessage(query.get("receiver"))
        sendMessage()
    } else {
        let chat = document.querySelector(".chat")
        chat.className = "chat welcome"
        chat.textContent = "WELCOME TO CHAT"
        sendButton.style.display = "none"
        messageInput.style.display = "none"
    }

}

export function sendMessage() {
    const storedData = localStorage.getItem("data");
    const parsedData = JSON.parse(storedData);
    const chat = document.querySelector(".content_post");
    let message = chat.querySelector("#messageInput");
    let sendButton = chat.querySelector("#sendButton");
    sendButton.addEventListener("click", () => {

        let receiver = new URLSearchParams(location.search).get("receiver")
        const messages = message.value.trim();
        if (messages) {
            displayMessage(parsedData.firstname, new Date(), messages, true);
            ws.send(
                JSON.stringify({
                    type: "broadcast",
                    content: messages,
                    sender: parsedData.id,
                    receiver: +receiver,
                    createAt: new Date()
                })
            );

        }
    });

}





// function showTypingNotification(userId) {
//     logMessage(`User ${userId} is typing...`);
//     setTimeout(() => {
//         logMessage("");
//     }, 3000);
// }

export function addStyle() {
    let style = document.createElement("link")
    style.rel = "stylesheet"
    style.href = "../static/css/chat.css"
    document.head.appendChild(style)

}
