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
	// handlers
	mux.HandleFunc("/api/register", userController.HandleRegister)//done
	mux.HandleFunc("/api/login", userController.HandleLogin)//done
	mux.HandleFunc("/api/post", postController.HandlePost)//done
	mux.HandleFunc("/api/home", homeController.HomeHandle)//done
	mux.HandleFunc("/api/card", homeController.GetCard_handler)//done
	mux.HandleFunc("/api/addcomment", commentController.Handler_AddComment)//done
	mux.HandleFunc("/api/comment", commentController.Handel_GetCommet)//dome
	mux.HandleFunc("/api/category", categoryController.HandelCategory)//done
	mux.HandleFunc("/api/profile/posts", profileController.HandleProfilePosts)//done
	mux.HandleFunc("/api/profile/likes", profileController.HandleProfileLikes)//done

	// mux.Handle("/api/likes", handlers.AuthenticateMiddleware((http.HandlerFunc(handlers.LikesHandle))))
	// mux.HandleFunc("/api/isLogged", handlers.HandleIsLogged)
	// mux.Handle("/api/like", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandelLike)))
	// mux.Handle("/api/deleted", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandelDeletLike)))
	// mux.Handle("/api/logout", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandleLogOut)))
}
