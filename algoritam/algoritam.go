package algoritam

import "fmt"

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
	Name     string
	Previous Reference
	Elements []interface{}
	First    *StartStruct
}

func NewAlgoritam(name string, previous Reference, first *StartStruct) *Algoritam {
	return &Algoritam{
		Name:     name,
		Previous: previous,
		First:    first,
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
	a.Elements = append(a.Elements, element)
	return nil
}

func (a *Algoritam) Arange() {
	for i, v := range a.Elements {
		if i == 0 {
			if a.First == nil {
				s, _ := v.(*BlockStruct)
				start := &StartStruct{
					Next: s,
				}
				a.First = start
			}
		}
		switch m := v.(type) {
		case *BlockStruct:
			for _, k := range a.Elements {
				switch m2 := k.(type) {
				case *BlockStruct:
					if m2.Previous != nil {
						if m2.Previous == v {
							m.Next = m2
						}
					}

				case *Romboid:
					if m2.Previous != nil {
						if m2.Previous == v {
							m.Next = m2
						}
					}
				case *EndStruct:
					if m2.Previous != nil {
						if m2.Previous == v {
							m.Next = m2
						}
					}
				case *StartStruct:
				default:
				}
			}

		case *Romboid:
			for _, k := range a.Elements {
				switch m2 := k.(type) {
				case *BlockStruct:
					if m2.Previous != nil {
						if m2.Previous == v {
							m.Next = m2
						}
					}

				case *Romboid:
					if m2.Previous != nil {
						if m2.Previous == v {
							m.Next = m2
						}
					}
				case *EndStruct:
					if m2.Previous != nil {
						if m2.Previous == v {
							m.Next = m2
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
