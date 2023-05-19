package algoritam

import (
	"fmt"
	"os"

	"github.com/zmajew/zerr"
)

type EndFunc func(*EndStruct)

type EndStruct struct {
	Previous Reference
	Func     EndFunc
	Name     string
}

func (a *Algoritam) NewEnd(previous Reference, name string, f EndFunc) *EndStruct {
	if name == "" {
		err := fmt.Errorf("error: cannot create an End with empty string name")
		zerr.Log(err, 2)
		os.Exit(1)
	}
	endStruct := &EndStruct{
		Previous: previous,
		Func:     f,
		Name:     name,
	}
	romb, ok := previous.(*Romboid)
	if ok {
		if romb.NextYes == nil {
			romb.NextYes = endStruct
		} else {
			if romb.NextNo == nil {
				romb.NextNo = endStruct
			}
		}
	}
	if err := a.add(endStruct); err != nil {
		zerr.Log(err, 2)
		os.Exit(1)
	}

	return endStruct
}

func (e *EndStruct) Execute(p Previous) {
	e.Func(e)
}

func (e *EndStruct) GetName() string {
	return e.Name
}

func (e *EndStruct) GetType() string {
	return "END"
}

func (e *EndStruct) GetPrevious() Reference {
	return e.Previous
}
