package dobbyS

import (
	"github.com/swathiGiligar/dobbyS/dobbynet"
	"github.com/varunamachi/vaali/vapp"
	"gopkg.in/urfave/cli.v1"
)

//NewModule - creates new dobby module
func NewModule() *vapp.Module {
	return &vapp.Module{
		Name:        "dobby",
		Description: "The dobby server",
		Endpoints:   dobbynet.GetEndpoints(),
		Commands:    []cli.Command{},
		Factories:   []vapp.Factory{},
		Initialize:  nil,
		Setup:       nil,
		Reset:       nil,
	}
}
