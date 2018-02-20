package dobbydb

import "gopkg.in/mgo.v2/bson"

//PTask represents a task
type PTask struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	Description    string        `json:"description" bson:"description"`
	Priority       string        `json:"priority" bson:"priority"`
	CreatedDate    string        `json:"createdDate" bson:"createdDate"`
	CreatedBy      string        `json:"createdBy" bson:"createdBy"`
	Owner          string        `json:"owner" bson:"owner"`
	ProjectName    string        `json:"projectName" bson:"projectName"`
	Dependencies   string        `json:"dependencies" bson:"dependencies"`
	Status         string        `json:"status" bson:"status"`
	AdditionalInfo string        `json:"additionalInfo" bson:"additionalInfo"`
}

//GetSample gets you the sample data for Task type
func GetSample() PTask {
	task := PTask{Description: "Task2",
		Priority:       "Medium",
		CreatedDate:    "13-02-2018",
		CreatedBy:      "Swathi",
		Owner:          "Swathi",
		ProjectName:    "Dobby",
		Dependencies:   "primeNG",
		Status:         "Active",
		AdditionalInfo: "Testing"}
	// tasks := make([]PTask, 0, 100)
	// tasks = append(tasks, task)
	return task
}

func (t *PTask) String() string {
	return "description: " + t.Description +
		"\npriority: " + t.Priority +
		"\ncreatedDate: " + t.CreatedDate +
		"\ncreatedBy: " + t.CreatedBy +
		"\nowner: " + t.Owner +
		"\nprojectName: " + t.ProjectName +
		"\ndependecies: " + t.Dependencies +
		"\nstatus: " + t.Status +
		"\nadditionlInfo: " + t.AdditionalInfo
}
