import { displayMessage } from "./displyMessage.js";

export function addUser(userId, userName, status) {
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



