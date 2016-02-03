package convert

import (
	"fmt"
	"github.com/katydid/katydid/expr/compose"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/interp"
	"github.com/katydid/katydid/serialize"
	"io"
)

func Eval(refs relapse.RefLookup, parser serialize.Parser) bool {
	converter := newConverter(refs)
	e := &eval{converter}
	start := e.converter.sconvert(refs["main"])
	final := e.Interpret(start, parser)
	return interp.Nullable(refs, e.converter.getPattern(final.current))
}

type eval struct {
	converter *converter
}

func (this *eval) Interpret(current *state, tree serialize.Parser) *state {
	for {
		fmt.Printf("current = %v\n", toStr(current, this.converter.patterns))
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		for _, t := range current.trans {
			f, err := compose.NewBoolFunc(t.value)
			if err != nil {
				panic(err)
			}
			e, err := f.Eval(tree)
			if err != nil {
				continue
			}
			if !e {
				continue
			}
			fmt.Printf("evaled %v\n", funcs.Sprint(t.value))
			down := this.converter.sconvert(this.converter.getPattern(t.down))
			fmt.Printf("down = %v\n", this.converter.patterns[down.current])
			up := down
			if tree.IsLeaf() {
				fmt.Printf("leaf\n")
			} else {
				fmt.Printf("Down\n")
				tree.Down()
				up = this.Interpret(down, tree)
				fmt.Printf("Up\n")
				tree.Up()
			}
			fmt.Printf("up = %v\n", this.converter.patterns[up.current])
			for _, cup := range t.ups {
				if cup.bot == up.current {
					current = this.converter.sconvert(this.converter.getPattern(cup.top))
					break
				}
			}
			break
		}
	}
	return current
}
