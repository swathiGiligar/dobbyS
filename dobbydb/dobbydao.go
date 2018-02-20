package dobbydb

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DobbyDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "tasks"
)

func (m *DobbyDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *DobbyDAO) FindAll() ([]PTask, error) {
	var tasks []PTask
	err := db.C(COLLECTION).Find(bson.M{}).All(&tasks)
	return tasks, err
}

func (m *DobbyDAO) FindById(id string) (PTask, error) {
	var task PTask
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&task)
	return task, err
}

func (m *DobbyDAO) Insert(task PTask) error {
	err := db.C(COLLECTION).Insert(&task)
	return err
}

func (m *DobbyDAO) Delete(task PTask) error {
	err := db.C(COLLECTION).Remove(&task)
	return err
}

func (m *DobbyDAO) Update(task PTask) error {
	err := db.C(COLLECTION).UpdateId(task.ID, &task)
	return err
}
