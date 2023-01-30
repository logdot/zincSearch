package zincindex

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// ParseMailFromReader takes a reader and parses it into the returned mail object.
func ParseMailFromReader(reader io.Reader) Mail {
	mail := Mail{}

	endHeader := false

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && !endHeader {
			endHeader = true
			continue
		}

		if !endHeader {
			ParseHeaderLine(&mail, scanner.Text())
		} else {
			ParseBodyLine(&mail, scanner.Text())
		}
	}

	return mail
}

// ParseHeaderLine takes a mail and line and parses the header at the given mail into the mail.
// Headers in the file get parsed into the properties of the passed in mail.
//
// If a header appears multiple times in a file, the most 'recent' one is used.
// The header parsing is case-sensitive, so "Content-Type: ..." would get detected as a header, but "content-type: ..." would not.
// All headers are trimmed of any spaces before being parsed.
//
// ParseHeaderLine will return an error if it's given an unrecognized header.
func ParseHeaderLine(mail *Mail, line string) error {
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
	default:
		return errors.New("Unrecognized header " + before)
	}

	return nil
}

// ParseBodyLine takes a mail and a line and appends the line to the body of the line
func ParseBodyLine(mail *Mail, line string) error {
	mail.Body += line + "\n"
	return nil
}
