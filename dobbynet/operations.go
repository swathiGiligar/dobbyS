package dobbynet

import (
	"fmt"
	"net/http"

	"github.com/varunamachi/vaali/vlog"

	"github.com/varunamachi/vaali/vsec"

	"github.com/labstack/echo"
	"github.com/swathiGiligar/dobbyS/dobbydb"
	"github.com/varunamachi/vaali/vnet"
)

var Db = dobbydb.DobbyDAO{}

func GetAllTasks(c echo.Context) error {
	tasks, err := Db.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			"Internal Server Error: "+err.Error())
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

func DeleteTaskByID(c echo.Context) error {
	id := c.Param("taskId")
	vlog.Info("Dobby:ID", id)
	err := Db.Delete(id)
	if err != nil {
		vlog.LogError("Dobby:Mongo", err)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, "Deleted")
}

func GetEndpoints() []*vnet.Endpoint {
	return []*vnet.Endpoint{
		&vnet.Endpoint{
			Method:   echo.POST,
			URL:      "dobby/tasks",
			Access:   vsec.Normal,
			Category: "tasks",
			Func:     CreatTask,
			Comment:  "Create new Task",
		},
		&vnet.Endpoint{
			Method:   echo.PUT,
			URL:      "dobby/tasks",
			Access:   vsec.Normal,
			Category: "tasks",
			Func:     UpdateTask,
			Comment:  "Update a task",
		},
		&vnet.Endpoint{
			Method:   echo.GET,
			URL:      "dobby/tasks/users/:userName",
			Access:   vsec.Normal,
			Category: "tasks",
			Func:     FindTaskByOwner,
			Comment:  "Retrieves a task for given user",
		},
		&vnet.Endpoint{
			Method:   echo.GET,
			URL:      "dobby/tasks",
			Access:   vsec.Normal,
			Category: "tasks",
			Func:     GetAllTasks,
			Comment:  "Retrieves all tasks",
		},
		&vnet.Endpoint{
			Method:   echo.DELETE,
			URL:      "dobby/tasks/:taskId",
			Access:   vsec.Normal,
			Category: "tasks",
			Func:     DeleteTaskByID,
			Comment:  "Delete a task",
		},
	}
}
