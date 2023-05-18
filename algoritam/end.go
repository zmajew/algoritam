package algoritam

type EndFunc func(*EndStruct)

type EndStruct struct {
	Previous Reference
	Func     EndFunc
	Name     string
}

func (a *Algoritam) NewEnd(previous Reference, f EndFunc) (*EndStruct, error) {
	endStruct := &EndStruct{
		Previous: previous,
		Func:     f,
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
		return nil, err
	}

	return endStruct, nil
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
