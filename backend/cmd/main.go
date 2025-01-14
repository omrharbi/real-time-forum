package main

import (
	"fmt"
	"log"
	"net/http"

	"real-time-froum/config"
	"real-time-froum/controllers"
	"real-time-froum/repo"
	"real-time-froum/services"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := config.InitDataBase()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to initialize database: %w", err))
	}

	mux := http.NewServeMux()

	SetupAPIRoutes(mux)
	// route.SetupPageRoutes(mux)

	serverAddr := ":3333"
	log.Printf("Server running at http://localhost%s/home\n", serverAddr)
	err = http.ListenAndServe(serverAddr, mux)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}

func SetupAPIRoutes(mux *http.ServeMux) {
	db := config.Config()
	userRepo := repo.NewUserRepository(db.Connection)
	cardRepo := repo.NewcardRepository(db.Connection)
	categoryRepo := repo.NewCategoryRepository(db.Connection)
	postRepo := repo.NewPostRepository(db.Connection)
	commentRepo := repo.NewCommentRepository(db.Connection)

	// serveses
	userService := services.NewUserService(userRepo)
	cardService := services.NewcardssService(cardRepo)
	commentService := services.NewCommentService(commentRepo, cardRepo)
	postService := services.NewPostService(postRepo, cardRepo, categoryRepo)
	// userService := services.NewUserService(userRepo)

	// 4. Initialize the controller
	userController := controllers.NewUserController(userService)
	homeController := controllers.NewHomeController(cardService)
	commentController := controllers.NewCommentController(commentService, userController)
	postController := controllers.NewpostController(postService, userController)

	mux.HandleFunc("/api/register", userController.HandleRegister)
	mux.HandleFunc("/api/login", userController.HandleLogin)
	mux.HandleFunc("/api/post", postController.HandlePost)
	mux.HandleFunc("/api/home", homeController.HomeHandle)
	mux.HandleFunc("/api/addcomment", commentController.Handler_AddComment)
	mux.HandleFunc("/api/comment", commentController.Handel_GetCommet)
	// mux.HandleFunc("/api/category", handlers.HandelCategory)
	// mux.HandleFunc("/api/login", handlers.HandleLogin)
	// mux.HandleFunc("/api/comment", handlers.Handel_GetCommet)
	// mux.HandleFunc("/api/card", handlers.GetCard_handler)
	// mux.HandleFunc("/api/isLogged", handlers.HandleIsLogged)
	// mux.Handle("/api/likes", handlers.AuthenticateMiddleware((http.HandlerFunc(handlers.LikesHandle))))
	// mux.Handle("/api/profile/posts", handlers.AuthenticateMiddleware((http.HandlerFunc(handlers.HandleProfilePosts))))
	// mux.Handle("/api/profile/likes", handlers.AuthenticateMiddleware((http.HandlerFunc(handlers.HandleProfileLikes))))
	// mux.Handle("/api/like", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandelLike)))
	// mux.Handle("/api/deleted", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandelDeletLike)))
	// mux.Handle("/api/logout", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandleLogOut)))
}
