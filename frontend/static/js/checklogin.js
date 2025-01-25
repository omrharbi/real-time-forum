import { fetchData } from "./forum.js";
 import { loadPage } from "./laodpages.js";

export function checklogin() {
  fetch("/api/isLogged").then((res) => {
    const path = window.location.pathname;
    if (res.ok) {
      if (path === "/login" || path === "/register") {
        history.pushState(null, "", "/");
        loadPage();
      }
    } else {
      if (path !== "/login" && path !== "/register") {
        history.pushState(null, "", "/login");
        loadPage();
      }
    }
  });
}

export const throttle = (func, wait = 100) => {
  let shouldWait = false;
  let waitArg;
  const timeFunc = () => {
    if (!waitArg) shouldWait = false;
    else {
      func(...waitArg);
      waitArg = null;
      setTimeout(timeFunc, wait);
    }
  };
  return (...arg) => {
    if (shouldWait) {
      waitArg = arg;
      return;
    }
    func(...arg);
    shouldWait = true;
    setTimeout(timeFunc, wait);
  };
};

export function Inf() {
  fetchData();
  let offset = 2;
  console.log(offset);
  const main = document.querySelector(".main");
  main.addEventListener(
    "scroll",
    throttle(() => {
      const windowHeight = main.clientHeight;
      const scrollTop = main.scrollTop;
      const scrollHeight = main.scrollHeight;

      // Load more data when the user scrolls close to the bottom
      if (scrollTop + windowHeight >= scrollHeight - 100) {
        fetchData(offset);
        offset += 1;
      }
    }, 500)
  );
}
