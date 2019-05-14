package formatter

import "strings"

func PreFormat(format string) string {
	//cut the space
	format = strings.Trim(format, " ")
	if format[len(format)-2:] != `\n` {
		format += "\n"
	}
	//input by cmd of "\t" is "\\t"
	replace := strings.NewReplacer(`\t`, "\t", `\n`, "\n")
	format = replace.Replace(format)
	return format
}
