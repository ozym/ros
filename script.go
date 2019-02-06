package ros

import (
	"strings"
)

func ParseSystemScriptSource(source string) string {
	var lines []string
	for _, l := range strings.Split(source, "\\\n") {
		if strings.HasSuffix(l, "}\"") {
			l = l + ";"
		}
		lines = append(lines, strings.Join(strings.Fields(l), " "))
	}
	return strings.Replace(strings.Join(lines, " "), `; \n`, `;\n`, -1)
}

func PostSystemScriptSource(source string) string {
	source = strings.TrimSpace(source)
	source = strings.Replace(source, `\\n`, "\n", -1)
	source = strings.Join(strings.Split(source, `\n`), " ")
	source = strings.Join(strings.FieldsFunc(source, func(c rune) bool { return c == ' ' }), " ")
	source = strings.Replace(source, "\n", `\n`, -1)
	source = strings.Replace(source, `\"`, "\"", -1)
	source = strings.Replace(source, `\$`, "$", -1)
	return source
}
