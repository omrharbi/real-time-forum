import { sendMessage } from "./compenetChat.js";
import { getMessage } from "./displyMessage.js";

export function addUser(userId, userName, status) {
    const userList = document.getElementById("userList");
    const userItem = document.createElement("li");
    userItem.className = "user-item";
    userItem.id = userId;
    userItem.dataset.id = userId;
    const userIcon = document.createElement("div");
    userIcon.className = "user-icon";
    userIcon.textContent = userName[0].toUpperCase();
 
    const userNameDiv = document.createElement("div");
    userNameDiv.className = "user-name";
    userNameDiv.textContent = userName;
 
    const statusDot = document.createElement("span");
    statusDot.className = "status";
    statusDot.style.background = status === "online" ? "green" : "red";
    userItem.append(userIcon, userNameDiv, statusDot);
    userList.appendChild(userItem);
    userItem.addEventListener("click", () => {
        document.querySelectorAll(".user-item").forEach((e) => {
            e.classList.remove("user-clicked");
        });

        userItem.classList.add("user-clicked");
        let url = `chat?receiver=${userId}`;
        history.pushState(null, "", url);
        const log = document.querySelector(".chat");
        if (log) {
            log.innerHTML = "";
        }

        getMessage(userId);
        sendMessage();
    });
}


export function updateUserList(message) {
    let id = document.getElementById(message.online_users)
    if (id) {
        let status = id.querySelector(".status")
        if (message.type === "online") {
            status.style.background = "green"
        } else {
            status.style.background = "red"
        }
    }
    console.log(id);

}



