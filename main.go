package main

import (
	"Golang-RestAPI/app"
	"Golang-RestAPI/controller"
	"Golang-RestAPI/helper"
	"Golang-RestAPI/middleware"
	"Golang-RestAPI/repository"
	"Golang-RestAPI/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	fmt.Println("Application is running...")

	err := server.ListenAndServe()
	helper.Panic(err)

}
