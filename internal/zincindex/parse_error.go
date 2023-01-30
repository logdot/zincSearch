package zincindex

import "strconv"

type ParseError struct {
	LineNumber uint
	Line       string
	Reason     string
}

func (e ParseError) Error() string {
	lineNumberString := strconv.FormatUint(uint64(e.LineNumber), 10)
	return "Error at line " + "\"" + e.Line + "\" at " + lineNumberString + " because of " + e.Reason
}
