package dobbydb

import (
	"github.com/varunamachi/vaali/vdb"
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

// func (m *DobbyDAO) Connect() {
// 	session, err := mgo.Dial(m.Server)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db = session.DB(m.Database)
// }

func (m *DobbyDAO) FindAll() ([]PTask, error) {
	conn := vdb.DefaultMongoConn()
	defer conn.Close()
	var tasks []PTask = make([]PTask, 0, 100)
	err := conn.C(COLLECTION).Find(bson.M{}).All(&tasks)
	return tasks, err
}

func (m *DobbyDAO) FindById(id string) (PTask, error) {
	conn := vdb.DefaultMongoConn()
	defer conn.Close()
	var task PTask
	err := conn.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&task)
	return task, err
}

func (m *DobbyDAO) Insert(task PTask) error {
	conn := vdb.DefaultMongoConn()
	defer conn.Close()
	err := conn.C(COLLECTION).Insert(&task)
	return err
}

func (m *DobbyDAO) Delete(id string) error {
	conn := vdb.DefaultMongoConn()
	defer conn.Close()
	err := conn.C(COLLECTION).Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (m *DobbyDAO) Update(task PTask) error {
	conn := vdb.DefaultMongoConn()
	defer conn.Close()
	err := conn.C(COLLECTION).UpdateId(task.ID, &task)
	return err
}

func (m *DobbyDAO) FindByOwner(user string) ([]PTask, error) {
	conn := vdb.DefaultMongoConn()
	defer conn.Close()
	var tasks []PTask
	err := conn.C(COLLECTION).Find(bson.M{"owner": user}).All(&tasks)
	return tasks, err
}
