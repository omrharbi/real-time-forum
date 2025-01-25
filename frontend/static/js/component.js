import { Inf } from "./checklogin.js";
import { comments } from "./globa.js";
import { Change } from "./home.js";

const section = document.querySelector("section");
const main = /*html*/ `
      <main class ="scroll">
        <div class="alert"></div>
        <div class="headMian">
          <img class="logo" src="../static/imgs/logo.png" alt="logo" />
        </div>
        <article class="main content_post">

        </article>
      </main>

      <div id="side-right"></div>
      <aside class="aside-right">
        <input class="search" type="text" placeholder="Search.." data-search />
        <div class="link-list">
          <a href="/about">about</a>·<a href="/contact">contact</a>
        </div>
      </aside>`;
const nav_item = [
  {
    id: "chat",
    name: "chatbubbles",
  },
  {
    id: "settings",
    name: "settings",
  },
  {
    id: "profile",
    name: "person-circle",
  },
  {
    id: "categories",
    name: "filter-circle",
  },
  {
    id: "home",
    name: "home",
  },

];

const categories = [
  "General",
  "Technology",
  "Sports",
  "Entertainment",
  "Science",
  "Health",
  "Food",
  "Travel",
  "Fashion",
  "Art",
  "Music",
];

function rightSide() {
  const aside = section.querySelector(".aside-right");
  aside.innerHTML = /*html*/ `
       <input class="search" type="text" placeholder="Search.." data-search />
          <div class="header-nav">
                <h1>Choose Your Categories:</h1>
                <nav class="profile-nav">
                </nav> 
            </div>
            <div class="link-list">
                <a href="/about">about</a>·<a href="/contact">contact</a>
            </div>
  `;
  const nav = aside.querySelector("nav");
  SetcategoriesOption(nav);
  section.append(aside);
}

function leftside() {
  section.innerHTML = main;
 
  const aside = document.createElement("aside");
  aside.className = "aside-left";
  aside.innerHTML = /*html*/ `
    <div class="aside-nav">
          <div class="logo-user">
            <img src="../static/imgs/profilePic.png" class="avatar" alt="Profile picture" />
          </div>
          <nav>
            
            <button class="newPost-popup">
              <ion-icon name="duplicate"></ion-icon>
              <h1>New Post</h1>
            </button>
  
            <div id="creatPost-popup" class="add_post">
              <div class="newPost">
                <div class="newPost-header">
                  <button class="cancel-btn post-close">Cancel</button>
                  <button class="post-btn create-post create-comment">Post</button>
                </div>
                <div class="newPost-content">
                  <img src="../static/imgs/profilePic.png" class="avatar" alt="Profile picture" />
                  <textarea maxlength="1000" placeholder="What's up?" id="content" required></textarea>
                </div>
                <div class="openCategories" id="choice-categories">
                  <h1>Categories</h1>
                </div>
              </div>
            </div>
            <div id="categories-popup">
              <div class="newPost">
                <div class="newPost-header">
                  <button class="cancel-btn category">Cancel</button>
                  <h1>Choice you are categories:</h1>
                  <button class="post-btn done-post">Done</button>
                </div>
                <div class="categories-content">
                  <div class="categories-list">
                </div>
            </div>
          </nav>
        </div>
      `;

  const nav = aside.querySelector("nav");

  const div = aside.querySelector(".categories-list");
  SetIcon(nav);
  Setcategories(div);
  section.prepend(aside);
  Change();
  if (window.location.pathname === "/categories") {
    rightSide();
  } else if (window.location.pathname === "/comment") {
    CommtSide();
    document.querySelector("input").remove();
    return;
  }
}

function CommtSide() {
  const main = document.querySelector("main");
  main.innerHTML = comments;
}

function SetcategoriesOption(nav) {
  for (let obj of categories) {
    const divC = document.createElement("span");
    divC.innerText = obj;
    // divC.href = "#";
    nav.append(divC);
  }
}

function SetIcon(nav) {
  for (let obj of nav_item) {
    const a = document.createElement("span");
    // a.href = "#";
    a.id = obj.id;
    a.className = "nav-item";
    // console.log(window.location.pathname.slice(1), obj.id);

    if ("home" == obj.id) {
      a.classList.add("active");
    }
    a.innerHTML = /*html*/ `
      <ion-icon name="${obj.name}-outline"></ion-icon>
      <ion-icon class="active" name="${obj.name}"></ion-icon>
      <h1>${obj.id}</h1>
    `;
    nav.prepend(a);
  }
}

function Setcategories(div) {
  for (let obj of categories) {
    const divC = document.createElement("div");
    divC.className = "category-item";
    divC.innerText = obj;
    div.append(divC);
  }
}

export { leftside, rightSide };
