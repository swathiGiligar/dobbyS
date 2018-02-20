package dobbynet

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/swathiGiligar/dobbyS/dobbydb"
	"gopkg.in/mgo.v2/bson"
)

// //GetAllTasks will fetch all the tasks
// func GetAllTasks() {
// 	e := echo.New()
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())
// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 		AllowOrigins: []string{"*", "http://localhost:8080/tasks"},
// 		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
// 	}))
// 	e.GET("/tasks", getTask)
// 	e.Logger.Fatal(e.Start(":8080"))
// }

// func getTask(c echo.Context) error {
// 	t := dobbydb.GetSample()
// 	return c.JSON(http.StatusOK, t)
// // }

var Db = dobbydb.DobbyDAO{}

func ConnectToTaskDB(server string, database string) {
	Db.Server = server
	Db.Database = database
	Db.Connect()
}

func GetAllTasks(c echo.Context) error {
	tasks, err := Db.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, tasks)
}

func FindTaskByID(c echo.Context) error {
	id := c.QueryParam("id")
	task, err := Db.FindById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, task)
}

func CreatTask(c echo.Context) error {
	task := dobbydb.GetSample()
	task.ID = bson.NewObjectId()
	if err := Db.Insert(task); err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, task)
}

func UpdateTask(c echo.Context) error {
	var task dobbydb.PTask
	if err := Db.Update(task); err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, task)
}
