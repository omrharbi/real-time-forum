import { fetchData } from "./forum.js";

export async function checklogin() {
  const res = await fetch("/api/isLogged");
  if (res.ok) {
    history.pushState(null, "", "/");
  }
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
  let offset = 1;
  console.log(offset);
  const main = document.querySelector("main");
  console.log(main);

  //   main.addEventListener("scroll", () => {
  //     console.log("ok");
  //   });
  main.addEventListener("scroll", () => {
    console.log("ok");
  });

  //   window.addEventListener(
  //     "scroll",
  //     throttle(() => {
  //       let windowHight = window.innerHeight;
  //       let scrol = window.scrollY;
  //       if (scrol + windowHight > document.body.scrollHeight - 1000) {
  //         fetchData(offset);
  //         offset += 20;
  //       }
  //     })
  //   );
}
