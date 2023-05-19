package algoritam

import (
	"fmt"
	"os"

	"github.com/zmajew/zerr"
)

type Romboid struct {
	Name      string
	Previous  Reference
	NextYes   Reference
	NextNo    Reference
	Next      Reference
	Condition func(*Romboid)
	Yes       bool
}

func (a *Algoritam) NewRomboid(name string, previous Reference, condition func() bool, yesNext, noNext Reference) *Romboid {
	if name == "" {
		err := fmt.Errorf("error: cannot create a Romboid with empty string name")
		zerr.Log(err, 2)
		os.Exit(1)
	}
	f := func(b *Romboid) {
		if condition() {
			b.Next = b.NextYes
		} else {
			b.Next = b.NextNo
		}
	}
	newRomb := &Romboid{
		Name:      name,
		Previous:  previous,
		NextYes:   yesNext,
		Condition: f,
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
		zerr.Log(err, 2)
		os.Exit(1)
	}

	return newRomb
}

func (r *Romboid) Execute(p Previous) {
	r.Condition(r)
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
