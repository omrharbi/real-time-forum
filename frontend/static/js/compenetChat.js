
export function messamges() {
     
    const chat = section.querySelector(".content_post");
    chat.style.height = "100%"
    chat.innerHTML += /*html*/`
      
            <div class="chat-message">
                  <div class="users">
                    <h1>User Online:</h1>
                    <ul class="user-list" id="userList">
                        <!-- User items will be added dynamically -->
                    </ul>
                  </div>
                  <div class="message">
                    <h1>User Online:</h1>
                  </div>
            </div>
    `;

    console.log(chat);

}