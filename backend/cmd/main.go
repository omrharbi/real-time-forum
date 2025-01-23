package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

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
	ctx := context.Background()
	SetupPageRoutes(mux)
	SetupAPIRoutes(mux, ctx)

	serverAddr := ":8080"
	log.Printf("Server running at http://localhost%s/home\n", serverAddr)
	err = http.ListenAndServe(serverAddr, mux)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}

func SetupAPIRoutes(mux *http.ServeMux, ctx context.Context) {
	db := config.Config()
	userRepo := repo.NewUserRepository(db.Connection)
	cardRepo := repo.NewcardRepository(db.Connection)
	categoryRepo := repo.NewCategoryRepository(db.Connection)
	postRepo := repo.NewPostRepository(db.Connection)
	commentRepo := repo.NewCommentRepository(db.Connection)
	profiletRepo := repo.NewProfileRepository(db.Connection)
	likesRepo := repo.NewLikesRepository(db.Connection)
	messageRepo := repo.NewMessageRepository(db.Connection)

	// serveses
	userService := services.NewUserService(userRepo)
	cardService := services.NewcardssService(cardRepo)
	likesService := services.NewLikesServer(likesRepo)
	profileService := services.NewProfilesService(profiletRepo, postRepo)
	commentService := services.NewCommentService(commentRepo, cardRepo)
	postService := services.NewPostService(postRepo, cardRepo, categoryRepo)
	categoryService := services.NewcategorysService(categoryRepo, postRepo)
	messService := services.NewMessageService(messageRepo)
	// hubService := services.NewHub()
	// userService := services.NewUserService(userRepo)

	// 4. Initialize the controller
	userController := controllers.NewUserController(userService, ctx)
	homeController := controllers.NewHomeController(cardService)
	likesController := controllers.NewLikesController(likesService, userController)
	categoryController := controllers.NewcategoryController(categoryService)
	commentController := controllers.NewCommentController(commentService, userController)
	postController := controllers.NewpostController(postService, userController, cardService)
	profileController := controllers.NewprofileController(profileService, userController)
	middlewareController := middlewares.NewMeddlewireController(userService) //.NewMeddlewireController(userService)
	// hubController := controllers.NewHubController(hubService, userController) //.NewMeddlewireController(userService)
	// handlers
	mux.HandleFunc("/api/register", userController.HandleRegister) // done
	mux.HandleFunc("/api/login", userController.HandleLogin)
	mux.HandleFunc("/api/isLogged", userController.HandleIsLogged)
	// done
	newWs := controllers.NewManager(userController, messService, userService)
	mux.Handle("/api/post", middlewareController.AuthenticateMiddleware(http.HandlerFunc(postController.HandlePost)))
	mux.Handle("/api/home", middlewareController.AuthenticateMiddleware(http.HandlerFunc(homeController.HomeHandle)))
	mux.Handle("/api/card", middlewareController.AuthenticateMiddleware(http.HandlerFunc(homeController.GetCard_handler)))
	mux.Handle("/api/addcomment", middlewareController.AuthenticateMiddleware(http.HandlerFunc(commentController.Handler_AddComment)))
	mux.Handle("/api/category", middlewareController.AuthenticateMiddleware(http.HandlerFunc(categoryController.HandelCategory)))
	mux.Handle("/api/profile/posts", middlewareController.AuthenticateMiddleware(http.HandlerFunc(profileController.HandleProfilePosts)))
	mux.Handle("/api/profile/likes", middlewareController.AuthenticateMiddleware(http.HandlerFunc(profileController.HandleProfileLikes)))
	mux.Handle("/api/logout", middlewareController.AuthenticateMiddleware(http.HandlerFunc(userController.HandleLogOut)))
	mux.Handle("/api/connected", middlewareController.AuthenticateMiddleware(http.HandlerFunc(userController.HandleUserConnected)))
	mux.Handle("/api/likescheked", middlewareController.AuthenticateMiddleware(http.HandlerFunc(likesController.LikesCheckedHandle))) /// this get user liked card
	mux.Handle("/api/addlike", middlewareController.AuthenticateMiddleware(http.HandlerFunc(likesController.HandleAddLike)))
	mux.Handle("/api/deleted", middlewareController.AuthenticateMiddleware(http.HandlerFunc(likesController.HandleDeletLike)))
	mux.Handle("/ws", http.HandlerFunc(newWs.ServWs))

	mux.Handle("/api/comment", middlewareController.AuthenticateMiddleware(http.HandlerFunc(commentController.Handel_GetCommet)))
	mux.Handle("/api/messages", http.HandlerFunc(newWs.HandleGetMessages))
	// mux.Handle("/ws", http.HandlerFunc(hubController.Messages))
}

func SetupPageRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			// handlers.JsoneResponseError(w, r, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		suffix := r.URL.Path[len("/static/"):]

		if strings.Contains(suffix, ".css/") || strings.Contains(suffix, ".js/") || strings.Contains(suffix, ".png/") {
			// handlers.JsoneResponseError(w, r, "Not Found", http.StatusNotFound)
			return
		}

		if strings.Contains(suffix, ".js") {
			http.ServeFile(w, r, "../../frontend/static/"+suffix)
			return
		}

		allowedFiles := map[string]bool{
			"css/alert.css":       true,
			"css/styles.css":      true,
			"imgs/logo.png":       true,
			"imgs/profilePic.png": true,
		}

		if !allowedFiles[suffix] {
			// handlers.JsoneResponseError(w, r, "Access Forbidden", http.StatusForbidden)
			return
		}
		http.ServeFile(w, r, "../../frontend/static/"+suffix)
	})

	// mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
	// 	cookies, err := r.Cookie("token")
	// 	if err != nil || cookies == nil {
	// 		http.ServeFile(w, r, "../../frontend/templates/register.html")
	// 	} else {
	// 		http.Redirect(w, r, "/home", http.StatusSeeOther)
	// 	}
	// })
	// mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
	// 	cookies, err := r.Cookie("token")
	// 	if err != nil || cookies == nil {
	// 		http.ServeFile(w, r, "../../frontend/templates/login.html")
	// 	} else {
	// 		http.Redirect(w, r, "/home", http.StatusSeeOther)
	// 	}
	// })
	// mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "../../frontend/templates/about.html")
	// })
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../frontend/templates/home.html")
	})
	// mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "../../frontend/templates/home.html")
	// })
	// mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "../../frontend/templates/categories.html")
	// })
	// mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "../../frontend/templates/contact.html")
	// })
	// mux.HandleFunc("/comment", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "../../frontend/templates/comment.html")
	// })
	// // mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// // 	validatePath(w, r)
	// // })
	// mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
	// 	cookies, err := r.Cookie("token")
	// 	if err != nil || cookies == nil {
	// 		http.Redirect(w, r, "/login", http.StatusSeeOther)
	// 	} else {
	// 		http.ServeFile(w, r, "../../frontend/templates/profile.html")
	// 	}
	// })
	// mux.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
	// 	cookies, err := r.Cookie("token")
	// 	if err != nil || cookies == nil {
	// 		http.Redirect(w, r, "/login", http.StatusSeeOther)
	// 	} else {
	// 		http.ServeFile(w, r, "../../frontend/templates/settings.html")
	// 	}
	// })
	// mux.HandleFunc("/err", func(w http.ResponseWriter, r *http.Request) {
	// 	filePath := "../../frontend/templates/err.html"
	// 	fileContent, err := os.ReadFile(filePath)
	// 	if err != nil {
	// 		// handlers.JsoneResponse(w, r, "Error loading the error page", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Write(fileContent)
	// })
}

// func SetupPageRoutes(mux *http.ServeMux) {
// 	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodGet {
// 			fmt.Println("Errorororororo")
// 			return
// 		}

// 		suffix := r.URL.Path[len("/static/"):]

// 		if strings.Contains(suffix, ".css/") || strings.Contains(suffix, ".js/") || strings.Contains(suffix, ".png/") {
// 			fmt.Println("Errorororororo")
// 			return
// 		}

// 		if strings.Contains(suffix, ".js") {
// 			http.ServeFile(w, r, "../../frontend/static/"+suffix)
// 			return
// 		}

// 		allowedFiles := map[string]bool{
// 			"css/alert.css":       true,
// 			"css/styles.css":      true,
// 			"imgs/logo.png":       true,
// 			"imgs/profilePic.png": true,
// 		}

// 		if !allowedFiles[suffix] {
// 			fmt.Println("Errorororororo")
// 			return
// 		}
// 		http.ServeFile(w, r, "../../frontend/static/"+suffix)
// 	})

// 	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "../../frontend/templates/login.html")
// 	})

// 	mux.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "../../frontend/templates/message.html")
// 	})
// }
