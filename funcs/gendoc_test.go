package funcs

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
)

func toString(s string) string {
	return strings.ToLower(strings.Replace(strings.Replace(s, "SINGLE_", "", 1), "LIST_", "[]", 1))
}

func TestGenFuncList(t *testing.T) {
	decoders := map[string]bool{
		"decDouble":   true,
		"decFloat":    true,
		"decInt64":    true,
		"decUint64":   true,
		"decInt32":    true,
		"decFixed64":  true,
		"decFixed32":  true,
		"decBool":     true,
		"decString":   true,
		"decBytes":    true,
		"decUint32":   true,
		"decEnum":     true,
		"decSFixed32": true,
		"decSFixed64": true,
		"decSInt32":   true,
		"decSInt64":   true,
		"exists":      true,
	}
	funcs := []string{}
	for name, us := range funcsMap.nameToUniq {
		for _, u := range us {
			f := funcsMap.uniqToFunc[u]
			ins := make([]string, len(f.In))
			for i, in := range f.In {
				ins[i] = toString(in.String())
			}
			if decoders[name] {
				ins = append(ins, "bytes")
			}
			funcs = append(funcs, fmt.Sprintf("func %v(%v) %v", name, strings.Join(ins, ","), toString(f.Out.String())))
		}
	}
	sort.Strings(funcs)
	fmt.Fprintf(os.Stderr, "%v\n", strings.Join(funcs, "\n"))
}
