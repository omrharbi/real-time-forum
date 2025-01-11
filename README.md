# real-time-forum
 backend/
├── config/
│   └── db.go          # Database connection setup
├── controllers/
│   └── userController.go
├── middlewares/
│   └── authMiddleware.go
├── models/
│   └── user.go        # User schema
├── repository/
│   └── userRepository.go # Database queries for users
├── routes/
│   └── routes.go      # Main router
├── services/
│   └── userService.go # Business logic for users
├── server.go          # Entry point
└── .env               # Environment variables
