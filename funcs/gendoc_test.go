package funcs

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
)

func toString(s string) string {
	return strings.ToLower(strings.Replace(strings.Replace(strings.Replace(s, "_BYTES", "_[]byte", 1), "SINGLE_", "", 1), "LIST_", "[]", 1))
}

func TestGenFuncList(t *testing.T) {
	funcs := []string{}
	for name, us := range funcsMap.nameToUniq {
		for _, u := range us {
			f := funcsMap.uniqToFunc[u]
			ins := make([]string, len(f.In))
			for i, in := range f.In {
				ins[i] = toString(in.String())
			}
			funcs = append(funcs, fmt.Sprintf("func %v(%v) %v", name, strings.Join(ins, ","), toString(f.Out.String())))
		}
	}
	sort.Strings(funcs)
	fmt.Fprintf(os.Stderr, "%v\n", strings.Join(funcs, "\n"))
}
