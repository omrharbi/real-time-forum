

export function alertPopup(data) {
    const params = location.pathname
    let cardId = localStorage.getItem("card_id");
        console.log(cardId);
        
     let alert=document.querySelector(".alert")
     
    alert.innerHTML+=`
        <div class="popup-container" id="popup">
            <div class="popup-message">
            <div class="success-icon">
                <i class="fas fa-check"></i>
            </div>
            <div class="message-content">
                <h3 class="meesage">Error</h3>
                <p class="error">${data.message}</p>
            </div>
            </div>
        </div>
    `
    function showPopup() {
        const popup = document.getElementById('popup');
        popup.classList.add('show');
    }

    function hidePopup() {
        const popup = document.getElementById('popup');
        popup.classList.remove('show');
        if(params==="/comment"){
            location.href=`/comment?card_id=${cardId}`
        }
        
    }
    showPopup()
    setTimeout(hidePopup, 900);

} 