package main

import (
	"os"

	"github.com/swathiGiligar/dobbyS"
	"github.com/varunamachi/vaali/vapp"
	"github.com/varunamachi/vaali/vcmn"
	"github.com/varunamachi/vaali/vmgo"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {

	app := vapp.NewWebApp(
		"dobby",
		vcmn.Version{
			Major: 0,
			Minor: 0,
			Patch: 0,
		},
		"0",
		[]cli.Author{
			cli.Author{
				Name: "Swathi Giligar",
			},
		},
		true,
		"Dobby The Task Manager",
	)
	app.Modules = append(app.Modules, dobbyS.NewModule())
	vmgo.SetDefaultDB("tasks_db")
	app.Exec(os.Args)

	// server := "localhost:27017"
	// database := "tasks_db"
	// dobbynet.ConnectToTaskDB(server, database)

	// e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*", "http://localhost:8080/tasks"},
	// 	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	// }))
	// e.GET("/tasks/users/:userName", dobbynet.FindTaskByOwner)
	// e.GET("/tasks", dobbynet.GetAllTasks)
	// e.POST("/tasks", dobbynet.CreatTask)
	// e.PUT("/tasks", dobbynet.UpdateTask)
	// e.Logger.Fatal(e.Start(":8080"))
}
