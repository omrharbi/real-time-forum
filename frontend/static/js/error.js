import { loadPage } from "./laodpages.js";

export const pageNotFound = `
<div class="container">
        <h1 class="status-code">${window.statusCode}</h1>
        <p class="message-error">${window.textCode}</p>
        <span >
            <button class="back-button">Back to Home</button>
        </span>
    </div>
`;

export function RetunHome() {
  document.querySelector(".back-button").addEventListener("click", () => {
    history.pushState(null, "", "/");
    loadPage();
  });
}
