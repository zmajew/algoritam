package algoritam

type Romboid struct {
	Name     string
	Previous Reference
	NextYes  Reference
	NextNo   Reference
	Next     Reference
	Func     func(*Romboid)
	Yes      bool
}

func (a *Algoritam) NewRomboid(name string, previous Reference, condition bool, yesNext, noNext Reference) (*Romboid, error) {
	f := func(b *Romboid) {
		if condition {
			b.Next = b.NextYes
		} else {
			b.Next = b.NextNo
		}
	}
	newRomb := &Romboid{
		Name:     name,
		Previous: previous,
		NextYes:  yesNext,
		Func:     f,
	}
	romb, ok := previous.(*Romboid)
	if ok {
		if romb.NextYes == nil {
			romb.NextYes = newRomb
		} else {
			if romb.NextNo == nil {
				romb.NextNo = newRomb
			}
		}
	}

	if err := a.add(newRomb); err != nil {
		return nil, err
	}

	return newRomb, nil
}

func (r *Romboid) Execute(p Previous) {
	r.Func(r)
	r.Next.Execute(p)
}

func (r *Romboid) GetName() string {
	return r.Name
}

func (r *Romboid) GetType() string {
	return "ROMBOID"
}

func (r *Romboid) GetPrevious() Reference {
	return r.Previous
}

func (r *Romboid) FirstPreviousBlockResult() interface{} {
	for {
		block, ok := r.Previous.(*BlockStruct)
		if ok {
			return block.Result
		}
		romboid, ok := r.Previous.(*Romboid)
		if ok {
			return romboid.FirstPreviousBlockResult()
		}
	}
}
