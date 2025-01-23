

const messageInput = document.getElementById("messageInput");
const sendButton = document.getElementById("sendButton");

const cookies = document.cookie.split("token=")[1];
const storedData = localStorage.getItem("data");
const parsedData = JSON.parse(storedData);
let receiver = new URLSearchParams(location.search).get("receiver")
// document.addEventListener("DOMContentLoaded", (c) => {
//     sendMessage(receiver)
// })
let ws
export function setupWs() {
    fetchConnectedUsers()
    ws = new WebSocket("ws://localhost:8080/ws");
    ws.onopen = () => {
        console.log("is connected");

    };

    ws.onmessage = async (event) => {
        const message = JSON.parse(event.data);
        switch (message.type) {
            case "online":
                console.log(message);
                
                updateUserList(message)

                // updateUserList(message);
                break;
            case "broadcast":
                displayMessage(message.sender, message);
                break;
            case "typing":
                showTypingNotification(message.userId);
                break;
            case "offline":
                 
                // showTypingNotification(message.userId);
                break;
            default:
                console.warn("Unhandled message type:", message.type);
        }
    };

    ws.onclose = () => {
        console.log("WebSocket connection closed.");
        logMessage("WebSocket disconnected.");
    };

    ws.onerror = (error) => {
        console.error("WebSocket error:", error);
    };
}

export function messamges() {
    const chat = document.querySelector(".content_post");
    chat.style.height = "100%"
    chat.innerHTML += /*html*/`
      
            <div class="chat-message">
                    <div class="users">
                        <h1 class="user-online">User Online:</h1>
                        <ul class="user-list" id="userList">
                            <!-- User items will be added dynamically -->
                        </ul>
                    </div>
                    <div class="message">
                        <h1>User Online:</h1>
                        <div class="chat"></div>
                            <div>
                                <input type="text" id="messageInput" placeholder="Type your message here..." />
                                <button id="sendButton">Send</button>
                            </div>
                    </div>
            </div>
    `;

}


async function fetchConnectedUsers() {
    const response = await fetch("/api/connected");
    if (response.ok) {
        const userList = document.getElementById("userList");
        userList.innerHTML = ""
        const users = await response.json();
        users.forEach((user) => {
            console.log(user);
            
            addUser(user.id, user.username, user.status)
        })
        user_item()
    }
}
function addUser(userId, userName, status) {
    const userList = document.getElementById("userList");
    const userItem = document.createElement("li");
    userItem.className = "user-item";
    userItem.id = userId;
    userItem.dataset.id = userId

    const userIcon = document.createElement("div");
    userIcon.className = "user-icon";
    userIcon.textContent = userName[0].toUpperCase();

    const userNameDiv = document.createElement("div");
    userNameDiv.className = "user-name";
    userNameDiv.textContent = userName;

    const statusDot = document.createElement("span");
    statusDot.className = "status";
    // if (userId === status.online_users) {//online_users type
    //     statusDot.style.background ="green"
    // }else{
    //     statusDot.style.background ="red"
    // }
    statusDot.style.background = status === "online" ? "green" : "red";
    userItem.append(userIcon, userNameDiv, statusDot);
    userList.appendChild(userItem);


}

function genreteMessages() {
    let chat = document.querySelector(".chat")
}

function updateUserList(message) {
     
    let id = document.getElementById(message.online_users)
    if(id){
        id.style.background = "green"
    }
    console.log(id);
 
}

function displayMessage(sender, content, isOwnMessage = false) {
    let log = document.querySelector(".chat");

    const messageDiv = document.createElement("div");
    messageDiv.textContent = `${isOwnMessage ? "You" : sender}: ${content.content}`;
    log.appendChild(messageDiv);
    log.scrollTop = log.scrollHeight;
}

function logMessage(message) {
    const logDiv = document.createElement("div");
    logDiv.textContent = message;
    log.appendChild(logDiv);
    log.scrollTop = log.scrollHeight;
}

function showTypingNotification(userId) {
    logMessage(`User ${userId} is typing...`);
    setTimeout(() => {
        logMessage("");
    }, 3000);
}

export function user_item() {
    let items = document.querySelectorAll(".user-item")
     items.forEach((clik) => {
        clik.addEventListener("click", () => {
            let id = clik.getAttribute("data-id")
            let url = `chat?receiver=${id}`
            history.pushState(null, "", url)

        })

    })
}

function sendMessage(receiver) {
    sendButton.addEventListener("click", () => {
        console.log(+receiver, parsedData.id);
        const message = messageInput.value.trim();
        if (message) {
            ws.send(
                JSON.stringify({
                    type: "broadcast",
                    content: message,
                    sender: parsedData.id,
                    receiver: +receiver
                })
            );

        }
    });
}
