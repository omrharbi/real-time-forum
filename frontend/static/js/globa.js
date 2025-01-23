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
              ><a href="#">Forgot Password?</a>
            </div>
            <button type="submit" class="btn">Login</button>
            <div class="login-register">
              <p>
                Don't have an account?
                <a
                  href="#"
                  class="register-link"
                  >Register</a
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
                            <a href="#" class="login-link" onclick="window.location.href='/login'">Login</a>
                        </p>
                        <p>
                            Do you want to go back?
                            <a class="home-link" href="/home">Home</a>
                        </p>
                    </div>
                </form>
            </div>
        </div>
    </div>
`;

export { register, login };
