import { loadPage } from "./laodpages.js";

export async function logout() {
  let Useruuid = getCookie("token");
  const response = await fetch("/api/logout", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ uuid: Useruuid }),
  });
  history.pushState(null, "", "/login");
  loadPage();
}

export function getCookie(name) {
  const cookies = document.cookie.split("; ");
  for (let i = 0; i < cookies.length; i++) {
    const [key, value] = cookies[i].split("=");
    if (key === name) {
      return value;
    }
  }
  return null;
}
