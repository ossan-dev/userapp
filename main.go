package main

import (
	"net/http"

	"userapp/controllers"
	user_db "userapp/db"
	"userapp/models"
	"userapp/services"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	userDb := &user_db.UserDb{
		Conn: db,
	}

	userDb.Conn.AutoMigrate(&models.User{})

	userService := services.NewUserService(userDb)
	userCtrl := controllers.NewUserController(userService)

	r := mux.NewRouter()

	r.HandleFunc("/users", userCtrl.Save)

	if err := http.ListenAndServe(":8000", r); err != nil {
		panic(err)
	}
}
