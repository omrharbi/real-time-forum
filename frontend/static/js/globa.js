const login = /*html*/ `
    <div class="logoAndName">
        <a href="/home">
          <img class="logo" src="../static/imgs/logo.png" alt="logo" />
        </a>
        <h2>Bluezone</h2>
      </div>
      <div class="alert"></div>
      <div class="wrapper">
        <div class="form-box login">
          <h2>Login</h2>
          <form action="" id="login">
            <div class="input-box">
              <span class="icon"><ion-icon name="mail"></ion-icon></span>
              <input type="text" id="email" required />
              <label>Email</label>
            </div>
            <div class="input-box">
              <span class="icon"><ion-icon name="lock-closed"></ion-icon></span>
              <input type="password" id="password" required />
              <label>Password</label>
            </div>
            <div class="remember-forgot">
              <label><input type="checkbox" />Remember me</label
              >
            </div>
            <button type="submit" class="btn">Login</button>
            <div class="login-register">
              <p>
                Don't have an account?
                <span
                  id="register"
                  class="register-link"
                  >Register</span
                >
              </p>
            </div>
          </form>
        </div>
      </div>
`;

const register = /*html*/ `
    <div class="alert"></div>
        <div class="logoAndName">
            <a href="/home">
                <img class="logo" src="../static/imgs/logo.png" alt="logo" />
            </a>
            <h2>Bluezone</h2>
        </div>
        <div class="wrapper active">
            <div class="form-box register">
                <h2>Registration</h2>
                <form action="" id="form-submit">
                    <div class="input-box">
                        <span class="icon"><ion-icon name="person-circle"></ion-icon></span>
                        <input type="text" required id="firstname" />
                        <label>Firstname</label>
                    </div>
                    <div class="input-box">
                        <span class="icon"><ion-icon name="person-circle"></ion-icon></span>
                        <input type="text" required id="lastname" />
                        <label>Lastname</label>
                    </div>
                    <div class="input-box">
                        <span class="icon"><ion-icon name="mail"></ion-icon></span>
                        <input type="text" required id="emailRegister" />
                        <label>Email</label>
                    </div>
                    <div class="input-box">
                        <span class="icon"><ion-icon name="lock-closed"></ion-icon></span>
                        <input type="password" required id="passwordRegister" />
                        <label>Password</label>
                    </div>
                    <div class="remember-forgot">
                        <label><input class="checkbox" type="checkbox" /> I agree to the terms
                            & conditions</label>
                    </div>
                    <button type="submit" class="btn">register</button>
                    <div class="login-register">
                        <p>
                            Already have an account?
                            <span class="login-link" id="login">Login</span>
                        </p>
                    </div>
                </form>
            </div>
        </div>
    </div>
`;

const comments = /*html*/ `
  <div class="alert"></div>
            <div class="post commens-card">
                <div class="post-header">
                    <img src="../static/imgs/profilePic.png" class="avatar" alt="Profile picture" />
                    <div class="user-info">
                        <div class="display-name full-name"></div>
                        <span class="username"></span>
                        <span class="timestamp time">1h</span>
                    </div>
                </div>
                <div class="post-content content">
                    
                </div>
                <div class="post-actions ">
                    <div class="action active is_liked" data-context="comment" id="likes" data-liked="false" data-like="like" data-id_card="" >
                      <svg width="17" height="17" viewBox="0 0 20 20" fill="currentColor">
                        <path d="M10 19c-.072 0-.145 0-.218-.006A4.1 4.1 0 0 1 6 14.816V11H2.862a1.751 1.751 0 0 1-1.234-2.993L9.41.28a.836.836 0 0 1 1.18 0l7.782 7.727A1.751 1.751 0 0 1 17.139 11H14v3.882a4.134 4.134 0 0 1-.854 2.592A3.99 3.99 0 0 1 10 19Zm0-17.193L2.685 9.071a.251.251 0 0 0 .177.429H7.5v5.316A2.63 2.63 0 0 0 9.864 17.5a2.441 2.441 0 0 0 1.856-.682A2.478 2.478 0 0 0 12.5 15V9.5h4.639a.25.25 0 0 0 .176-.429L10 1.807Z"></path>
                    </svg>
                        <span id="is_liked"></span>
                    </div>
                    <div  class="action disliked" id="likes" data-liked="false"  data-like="Dislikes" data-id_card="">
                      <svg width="17" height="17" viewBox="0 0 20 20" fill="currentColor">
                        <path d="M10 1c.072 0 .145 0 .218.006A4.1 4.1 0 0 1 14 5.184V9h3.138a1.751 1.751 0 0 1 1.234 2.993L10.59 19.72a.836.836 0 0 1-1.18 0l-7.782-7.727A1.751 1.751 0 0 1 2.861 9H6V5.118a4.134 4.134 0 0 1 .854-2.592A3.99 3.99 0 0 1 10 1Zm0 17.193 7.315-7.264a.251.251 0 0 0-.177-.429H12.5V5.184A2.631 2.631 0 0 0 10.136 2.5a2.441 2.441 0 0 0-1.856.682A2.478 2.478 0 0 0 7.5 5v5.5H2.861a.251.251 0 0 0-.176.429L10 18.193Z"></path>
                    </svg>  
                        <span id="is_Dislikes" data-disliked="disliked"></span>
                    </div>
                    <div class="action">
                      <svg width="17" height="17" viewBox="0 0 20 20" fill="currentColor">
                        <path d="M10 19H1.871a.886.886 0 0 1-.798-.52.886.886 0 0 1 .158-.941L3.1 15.771A9 9 0 1 1 10 19Zm-6.549-1.5H10a7.5 7.5 0 1 0-5.323-2.219l.54.545L3.451 17.5Z"></path>
                    </svg>
                        <span class="comments"></span>
                    </div>
                </div>
            </div>
            <div class="post">
                <div class="postReply">
                    <img src="../static/imgs/profilePic.png" class="avatar" alt="Profile picture" />
                    <div class="writeReply">Write your reply</div>
                </div>
            </div>  
            <div class="allcomment"></div>
`;

export { register, login, comments };
