package algoritam

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Reference interface {
	Execute(Previous)
	GetName() string
	GetPrevious() Reference
	GetType() string
}

type Previous interface {
	Exe()
}

// type Dilema interface {
// 	Execute()
// 	FromHere()
// }

type StartStruct struct {
	Next Reference
}

type Algoritam struct {
	Name             string
	Previous         Reference
	Elements         []interface{}
	First            *StartStruct
	creationCodeLine string
}

func NewAlgoritam(name string, previous Reference, first *StartStruct) *Algoritam {
	pc := make([]uintptr, 10)
	fk := runtime.FuncForPC(pc[1] - 1)
	osPath, _ := os.Getwd()
	_, fn, line, _ := runtime.Caller(1)
	fn = strings.TrimPrefix(fn, osPath)

	return &Algoritam{
		Name:             name,
		Previous:         previous,
		First:            first,
		creationCodeLine: fmt.Sprintf("----------ii %s %d, %s", fn, line, fk.Name()),
	}
}

func (a *Algoritam) Start() {
	a.First.Next.Execute(nil)
}

func (a *Algoritam) add(element interface{}) error {
	if len(a.Elements) == 0 {
		_, ok := element.(*BlockStruct)
		if !ok {
			return fmt.Errorf("error: first element of the algoritam must be of type Block, it is %T", element)
		}
	}
	err := a.checkName(element)
	if err != nil {
		return err
	}

	a.Elements = append(a.Elements, element)
	return nil
}

func (a *Algoritam) Arange() {
	for i, thisElemntInterface := range a.Elements {
		if i == 0 {
			if a.First == nil {
				s, _ := thisElemntInterface.(*BlockStruct)
				start := &StartStruct{
					Next: s,
				}
				a.First = start
			}
		}
		switch thisElemntStruct := thisElemntInterface.(type) {
		case *BlockStruct:
			for _, k := range a.Elements {
				switch otherElement := k.(type) {
				case *BlockStruct:
					if otherElement.Previous != nil {
						if otherElement.Previous == thisElemntInterface {
							thisElemntStruct.Next = otherElement
						}
					}

				case *Romboid:
					if otherElement.Previous != nil {
						if otherElement.Previous == thisElemntInterface {
							thisElemntStruct.Next = otherElement
						}
					}
				case *EndStruct:
					if otherElement.Previous != nil {
						if otherElement.Previous == thisElemntInterface {
							thisElemntStruct.Next = otherElement
						}
					}
				case *StartStruct:
				default:
				}
			}

		case *Romboid:
			for _, k := range a.Elements {
				switch otherElement := k.(type) {
				case *BlockStruct:
					if otherElement.Previous != nil {
						if otherElement.Previous == thisElemntInterface {
							thisElemntStruct.Next = otherElement
						}
					}

				case *Romboid:
					if otherElement.Previous != nil {
						if otherElement.Previous == thisElemntInterface {
							thisElemntStruct.Next = otherElement
						}
					}
				case *EndStruct:
					if otherElement.Previous != nil {
						if otherElement.Previous == thisElemntInterface {
							thisElemntStruct.Next = otherElement
						}
					}
				case *StartStruct:
				default:
				}
			}
		case *EndStruct:
		case *StartStruct:
		default:
		}
	}
}

func (a *Algoritam) Execute(Previous) {
	a.Start()
}

func (a *Algoritam) GetName() string {
	return a.Name
}

func (a *Algoritam) GetPrevious() Reference {
	return a.Previous
}

func (a *Algoritam) GetType() string {
	return "ALGORITAM"
}

func (a *Algoritam) getElemenType(element interface{}) string {
	return fmt.Sprintf("%T", element)
}

func (a *Algoritam) checkName(newElement interface{}) error {
	var elementName string
	switch m := newElement.(type) {
	case *BlockStruct:
		elementName = m.GetName()
	case *Romboid:
		elementName = m.GetName()
	case *EndStruct:
		elementName = m.GetName()
	}

	for _, v := range a.Elements {
		switch m := v.(type) {
		case *BlockStruct:
			if m.GetName() == elementName {
				return fmt.Errorf(`error: element with the name "%s" already exist, created on the location: %s`, elementName, m.creationCodeLine)
			}
		case *Romboid:
			if m.GetName() == elementName {
				return fmt.Errorf(`error: element with the name "%s" already exist, created on the location: %s`, elementName, m.creationCodeLine)
			}
		case *EndStruct:
			if m.GetName() == elementName {
				return fmt.Errorf(`error: element with the name "%s" already exist, created on the location: %s`, elementName, m.creationCodeLine)
			}
		}
	}
	return nil
}
