package main

import (
	authentication "chat_server/adapter/Authentication"
	"chat_server/adapter/database"
	"chat_server/router"
	"chat_server/usecase/interactor"
	"fmt"
	"log"
	"net/http"
)

	func main() {
		db, err := database.NewPostgresDB()
		if err != nil {
			log.Fatal(err)
		}
		defer func(){
			postgres, err := db.DB()
			if err != nil {
				log.Fatal(err)
			}
			
			err = postgres.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		err = database.Migrate(db)
		if err != nil {
			log.Fatal(err)
		}

		authRepo := authentication.NewUserAuth()
		userRepo := database.NewUserRepository(db)

		userUC := interactor.NewUserUsecase(interactor.NewUserUsecaseArgs{
			Auth: authRepo,
			User: userRepo,
		})
		s := router.NewServer(userUC)
		defer s.Close()

		server := &http.Server{
			Addr:    ":8080",
			Handler: s.Handler,
		}
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
		
	}

