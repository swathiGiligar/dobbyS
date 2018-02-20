package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/swathiGiligar/dobbyS/dobbynet"
)

func main() {

	server := "localhost:27017"
	database := "tasks_db"
	dobbynet.ConnectToTaskDB(server, database)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "http://localhost:8080/tasks"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.GET("/tasks", dobbynet.GetAllTasks)
	// e.GET("/create", dobbynet.CreatTask)
	e.Logger.Fatal(e.Start(":8080"))
}
