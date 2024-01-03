package hello

import "strings"

func Say2(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}
	return "Hello, " + strings.Join(names, ", ") + "!"
}
