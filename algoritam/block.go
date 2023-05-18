package algoritam

import (
	"fmt"
	"log"
)

type Block interface {
	Execute()
	AddPrevious(d Previous)
	GetName() string
}

type BlockFunc func(*BlockStruct) error

type BlockStruct struct {
	Name                string
	Type                string
	Previous            Reference
	Next                Reference
	Func                BlockFunc
	Error               error
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
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		if b == nil {
	// 			fmt.Printf(`"Next" of "%s" blockStruc empty
	// `)
	// 		}
	// 		fmt.Printf("panic occurred: %s\n", err)
	// 	}
	// }()
	err := b.Func(b)
	if err != nil {
		b.Error = err
		if b.ReferenceAfterError != nil {
			b.ReferenceAfterError.Execute(b)
			return
		}
		log.Fatal(err)
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

func (a *Algoritam) NewBlock(previous, next Reference, name string, f BlockFunc, rae Reference) (*BlockStruct, error) {
	if name == "" {
		err := fmt.Errorf("cannot create block with empty string name")
		return nil, err
	}

	block := &BlockStruct{
		Name:                name,
		Previous:            previous,
		Func:                f,
		Next:                next,
		ReferenceAfterError: rae,
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
		return nil, err
	}
	return block, a.add(block)
}

// func (r *BlockStruct) FirstPreviousBlockResult() interface{} {
// 	for {
// 		block, ok := r.Previous.(*BlockStruct)
// 		if ok {
// 			return block.Result
// 		}
// 		romboid, ok := r.Previous.(*Romboid)
// 		if ok {
// 			return romboid.FirstPreviousBlockResult()
// 		}
// 	}
// }
