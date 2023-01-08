package zincindex

import (
	"bufio"
	"os"
	"strings"
)

type Mail struct {
	MessageID        string `json:"message-id"`
	Date             string `json:"date"`
	From             string `json:"from"`
	To               string `json:"to"`
	Subject          string `json:"subject"`
	MimeVersion      string `json:"mime-version"`
	ContentType      string `json:"content-type"`
	TransferEncoding string `json:"transfer-encoding"`
	XFrom            string `json:"x-from"`
	XTo              string `json:"x-to"`
	XCC              string `json:"x-cc"`
	XBCC             string `json:"x-bcc"`
	XFolder          string `json:"x-folder"`
	XOrigin          string `json:"x-origin"`
	XFilename        string `json:"x-filename"`
	Body             string `json:"body"`
}

// ParseMailFromFile takes a file and returns a mail object.
//
// To parse the file it uses ParseLine
func ParseMailFromFile(file *os.File) Mail {
	mail := Mail{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ParseLine(&mail, scanner.Text())
	}

	return mail
}

// ParseLine is the parser used by ParseMailFromFile.
// It takes the contents of a mail line by line.
//
// Headers in the file get parsed into the properties of the returned mail.
// Anything that's not a header is appended into the body of the mail.
//
// If a header appears multiple times in a file, the most 'recent' one is used.
// The header parsing is case-sensitive, so "Content-Type: ..." would get detected as a header, but "content-type: ..." would not.
// All headers are trimmed of any spaces before being parsed, but body lines are left as is.
func ParseLine(mail *Mail, line string) {
	before, after, _ := strings.Cut(strings.Trim(line, " "), ":")
	after = strings.Trim(after, " ")

	switch before {
	case "Message-ID":
		mail.MessageID = after
	case "Date":
		mail.Date = after
	case "From":
		mail.From = after
	case "To":
		mail.To = after
	case "Subject":
		mail.Subject = after
	case "Mime-Version":
		mail.MimeVersion = after
	case "Content-Type":
		mail.ContentType = after
	case "Content-Transfer-Encoding":
		mail.TransferEncoding = after
	case "X-From":
		mail.XFrom = after
	case "X-To":
		mail.XTo = after
	case "X-cc":
		mail.XCC = after
	case "X-bcc":
		mail.XBCC = after
	case "X-Folder":
		mail.XFolder = after
	case "X-Origin":
		mail.XOrigin = after
	case "X-FileName":
		mail.XFilename = after
	}
}
