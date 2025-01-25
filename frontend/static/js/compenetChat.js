
const cookies = document.cookie.split("token=")[1];
const storedData = localStorage.getItem("data");
const parsedData = JSON.parse(storedData);

let ws
export function setupWs() {

    ws = new WebSocket("ws://localhost:8080/ws");
    ws.onopen = () => {
        console.log("is connected");

    };

    ws.onmessage = async (event) => {
        const message = JSON.parse(event.data);
        console.log(message);

        switch (message.type) {
            case "online":
                updateUserList(message)
                break;
            case "broadcast":
                console.log(message, "herererer");
                displayMessage(message.sender, message.content, false);
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
        //logMessage("WebSocket disconnected.");
    };

    ws.onerror = (error) => {
        console.error("WebSocket error:", error);
    };
}

export function messages() {
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
    sendMessage()
}


export async function fetchConnectedUsers() {
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

    userItem.append(userIcon, userNameDiv, statusDot);
    userList.appendChild(userItem);
    statusDot.style.background = status === "online" ? "green" : "red";

}

function updateUserList(message) {

    let id = document.getElementById(message.online_users)
    let status = id.querySelector(".status")
    console.log(message.type);


    if (id) {
        if (message.type === "online") {
            status.style.background = "green"
        } else {
            status.style.background = "red"
        }
    }
    console.log(id);

}

function displayMessage(sender, content, isOwnMessage = false) {
    let log = document.querySelector(".chat");

    const messageUser = document.createElement("div");// 
    const message_content = document.createElement("div");
    const time = document.createElement("div");

    messageUser.className = "message";
    message_content.className = "message-content"
    time.className = "time";
    message_content.textContent = `${isOwnMessage ? "You" : sender}: ${content}`;

    if (isOwnMessage) {
        messageUser.classList = "bot";
    } else {
        messageUser.className = "user";
    }
    messageUser.append(message_content, time);
    log.appendChild(messageUser);
    log.scrollTop = log.scrollHeight;
}

// function showTypingNotification(userId) {
//     logMessage(`User ${userId} is typing...`);
//     setTimeout(() => {
//         logMessage("");
//     }, 3000);
// }

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


function sendMessage() {
    const chat = document.querySelector(".content_post");

    let message = chat.querySelector("#messageInput");
    let sendButton = chat.querySelector("#sendButton");
    console.log(sendButton);

    sendButton.addEventListener("click", () => {
        let receiver = new URLSearchParams(location.search).get("receiver")
        const messages = message.value.trim();
        if (messages) {
            displayMessage("You", messages, true);
            ws.send(
                JSON.stringify({
                    type: "broadcast",
                    content: messages,
                    sender: parsedData.id,
                    receiver: +receiver
                })
            );

        }
    });


}


export function addStyle() {
    let style = document.createElement("link")
    style.rel = "stylesheet"
    style.href = "../static/css/chat.css"
    document.head.appendChild(style)

}