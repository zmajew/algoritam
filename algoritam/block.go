package algoritam

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/zmajew/zerr"
)

type BlockFunc func(*BlockStruct) error

type BlockStruct struct {
	Name             string
	Previous         Reference
	Next             Reference
	Func             BlockFunc
	Error            error
	creationCodeLine string
	// ReferenceAfterError defines the Reference that will be trigered on execution error
	ReferenceAfterError Reference
}

func (b *BlockStruct) GetName() string {
	return b.Name
}

func (b *BlockStruct) GetPrevious() Reference {
	return b.Previous
}

func (b *BlockStruct) GetType() string {
	return "BLOCK"
}

func (b *BlockStruct) Execute(Previous) {
	defer func() {
		if err := recover(); err != nil {
			b.Error = fmt.Errorf("%s \n %s", err, string(debug.Stack()))
			if b.ReferenceAfterError != nil {
				b.ReferenceAfterError.Execute(b)
				return
			}
			panic(err)
		}
	}()

	err := b.Func(b)
	if err != nil {
		b.Error = err
		if b.ReferenceAfterError != nil {
			b.ReferenceAfterError.Execute(b)
			return
		}
		panic(err)
	}
	b.Next.Execute(b)
}

func (b *BlockStruct) Exe() {
	b.Func(b)
}

func (b *BlockStruct) AddPrevious(d Reference) {
	if b.Previous != nil {
		fmt.Printf("\block %s already have previous block, cannot add another in the code.", b.Name)
		return
	}
	b.Previous = d
}

func (a *Algoritam) NewBlock(previous, next Reference, name string, f BlockFunc, rae Reference) *BlockStruct {
	if name == "" {
		err := fmt.Errorf("error: cannot create a Block with empty string name")
		zerr.Log(err, 2)
		os.Exit(1)
	}

	pc := make([]uintptr, 10)
	fk := runtime.FuncForPC(pc[1] - 1)
	osPath, _ := os.Getwd()
	_, fn, line, _ := runtime.Caller(1)
	fn = strings.TrimPrefix(fn, osPath)

	block := &BlockStruct{
		Name:                name,
		Previous:            previous,
		Func:                f,
		Next:                next,
		ReferenceAfterError: rae,
		creationCodeLine:    fmt.Sprintf("%s %d, %s", fn, line, fk.Name()),
	}

	if rae == nil {
		rae = &EndStruct{
			Previous: block,
			Func: func(es *EndStruct) {
				fmt.Println("error:", block.Error.Error())
				fmt.Printf(`error from the node: "%s"`, block.GetName())
				next = block
				actual := block.Previous
				for {
					if actual == nil {
						fmt.Printf("\n")
						break
					}
					romb, ok := actual.(*Romboid)
					if ok {
						if romb.NextNo == next {
							fmt.Printf("\nfrom NO of the ROMBOID: %s", romb.GetName())
						}
						if romb.NextYes == next {
							fmt.Printf("\nfrom YES of the ROMBOID: %s", romb.GetName())
						}
					} else {
						fmt.Printf("\n")
						fmt.Printf(`%s "%s" %s %s`, "from", actual.GetName(), "of the type", actual.GetType())
					}

					next = actual
					actual = actual.GetPrevious()
				}
			},
		}
	}

	block.ReferenceAfterError = rae

	romb, ok := previous.(*Romboid)
	if ok {
		if romb.NextYes == nil {
			romb.NextYes = block
		} else {
			if romb.NextNo == nil {
				romb.NextNo = block
			}
		}
	}

	if err := a.add(block); err != nil {
		zerr.Log(err, 2)
		os.Exit(1)
	}
	return block
}
