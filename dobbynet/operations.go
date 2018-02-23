package dobbynet

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/swathiGiligar/dobbyS/dobbydb"
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
	task := new(dobbydb.PTask)
	if err := c.Bind(task); err != nil {
		fmt.Printf("Error" + err.Error())
		return c.JSON(http.StatusBadRequest, "Invlid Request")
	}
	// task.ID = bson.NewObjectId()
	if errDb := Db.Insert(*task); errDb != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, task)
}

func UpdateTask(c echo.Context) error {
	task := new(dobbydb.PTask)
	if err := c.Bind(task); err != nil {
		fmt.Printf("Error" + err.Error())
		return c.JSON(http.StatusBadRequest, "Invlid Request")
	}
	if err := Db.Update(*task); err != nil {
		fmt.Printf("Error" + err.Error())
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, task)
}

func FindTaskByOwner(c echo.Context) error {
	userName := c.Param("userName")
	fmt.Printf("The username sent is: " + userName)
	tasks, err := Db.FindByOwner(userName)
	if err != nil {
		fmt.Printf("The username sent is: " + err.Error())
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, tasks)
}
