# Real-Time Forum

A real-time forum application built with **Go** (Golang) for the backend and **HTML, CSS, JavaScript** for the frontend. This application allows users to interact with posts, comments, and enjoy real-time interactions.

---

## Project Structure

### Backend

```plaintext
backend/
├── config/
│   └── db.go          # Database connection setup
├── controllers/
│   └── userController.go  # Handles user-related actions
├── middlewares/
│   └── authMiddleware.go  # Middleware for authentication
├── models/
│   └── user.go        # User schema
├── repository/
│   └── userRepository.go  # Database queries for users
├── routes/
│   └── routes.go      # Main router with route definitions
├── services/
│   └── userService.go  # Business logic related to users
├── server.go          # Entry point of the backend server
└── .env               # Environment variables (DB credentials, JWT secret)

frontend/
├── public/
│   └── index.html      # Main HTML file
├── src/
│   └── app.js          # Main JavaScript file handling app logic
│   └── styles.css      # CSS file for styling
│   └── images/         # Folder for images (if required)
├── .env                # Environment variables (API URL)
