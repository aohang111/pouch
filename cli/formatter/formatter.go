package formatter

import (
	"sort"
	"strings"
)

func PreFormat(format string) string {
	//cut the space
	format = strings.Trim(format, " ")
	//input by cmd of "\t" is "\\t"
	replace := strings.NewReplacer(`\t`, "\t", `\n`, "\n")
	format = replace.Replace(format)

	if format[len(format)-1:] != "\n" {
		format += "\n"
	}
	return format
}
func LabelsToString(labels map[string]string) string {
	var labelstring string
	sorted_keys := make([]string, 0)
	for key := range labels {
		sorted_keys = append(sorted_keys, key)
	}
	sort.Strings(sorted_keys)
	for _, key := range sorted_keys {
		keyValueString := key + "=" + labels[key] + " "
		labelstring += keyValueString
	}
	return labelstring
}
