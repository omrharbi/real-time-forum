package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"real-time-froum/config"
	"real-time-froum/controllers"
	"real-time-froum/middlewares"
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
	profiletRepo := repo.NewProfileRepository(db.Connection)
	
	// serveses
	userService := services.NewUserService(userRepo)
	cardService := services.NewcardssService(cardRepo)
	profileService := services.NewProfilesService(profiletRepo, postRepo)
	commentService := services.NewCommentService(commentRepo, cardRepo)
	postService := services.NewPostService(postRepo, cardRepo, categoryRepo)
	categoryService := services.NewcategorysService(categoryRepo, postRepo)
	// userService := services.NewUserService(userRepo)

	// 4. Initialize the controller
	userController := controllers.NewUserController(userService)
	homeController := controllers.NewHomeController(cardService)
	categoryController := controllers.NewcategoryController(categoryService)
	commentController := controllers.NewCommentController(commentService, userController)
	postController := controllers.NewpostController(postService, userController)
	profileController := controllers.NewprofileController(profileService, userController)
	middlewareController := middlewares.NewMeddlewireController(userService) //.NewMeddlewireController(userService)
	// handlers
	mux.HandleFunc("/api/register", userController.HandleRegister) // done
	mux.HandleFunc("/api/login", userController.HandleLogin)
	mux.HandleFunc("/api/isLogged", userController.HandleIsLogged)
	// done
	mux.Handle("/api/post", middlewareController.AuthenticateMiddleware(http.HandlerFunc(postController.HandlePost)))                     // Protected
	mux.Handle("/api/home", middlewareController.AuthenticateMiddleware(http.HandlerFunc(homeController.HomeHandle)))                     // Protected
	mux.Handle("/api/card", middlewareController.AuthenticateMiddleware(http.HandlerFunc(homeController.GetCard_handler)))                // Protected
	mux.Handle("/api/addcomment", middlewareController.AuthenticateMiddleware(http.HandlerFunc(commentController.Handler_AddComment)))    // Protected
	mux.Handle("/api/comment", middlewareController.AuthenticateMiddleware(http.HandlerFunc(commentController.Handel_GetCommet)))         // Protected
	mux.Handle("/api/category", middlewareController.AuthenticateMiddleware(http.HandlerFunc(categoryController.HandelCategory)))         // Protected
	mux.Handle("/api/profile/posts", middlewareController.AuthenticateMiddleware(http.HandlerFunc(profileController.HandleProfilePosts))) // Protected
	mux.Handle("/api/profile/likes", middlewareController.AuthenticateMiddleware(http.HandlerFunc(profileController.HandleProfileLikes))) // Protected
	mux.Handle("/api/logout", middlewareController.AuthenticateMiddleware(http.HandlerFunc(userController.HandleLogOut)))
	// mux.Handle("/api/likes", handlers.AuthenticateMiddleware((http.HandlerFunc(handlers.LikesHandle))))
	// mux.Handle("/api/like", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandelLike)))
	// mux.Handle("/api/deleted", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandelDeletLike)))
}
