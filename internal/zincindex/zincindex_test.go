package zincindex

import (
	"strings"
	"testing"
)

// diffMails is a helper function that compares two mails and fails the test if they're different
func diffMails(want Mail, got Mail, t *testing.T) {
	if want.MessageID != got.MessageID {
		t.Errorf("MessageID got '%s' want '%s'", got.MessageID, want.MessageID)
	}
	if want.Date != got.Date {
		t.Errorf("Date got '%s' want '%s'", got.Date, want.Date)
	}
	if want.From != got.From {
		t.Errorf("From got '%s' want '%s'", got.From, want.From)
	}
	if want.To != got.To {
		t.Errorf("To got '%s' want '%s'", got.To, want.To)
	}
	if want.Subject != got.Subject {
		t.Errorf("Subject got '%s' want '%s'", got.Subject, want.Subject)
	}
	if want.MimeVersion != got.MimeVersion {
		t.Errorf("MimeVersion got '%s' want '%s'", got.MimeVersion, want.MimeVersion)
	}
	if want.ContentType != got.ContentType {
		t.Errorf("ContentType got '%s' want '%s'", got.ContentType, want.ContentType)
	}
	if want.TransferEncoding != got.TransferEncoding {
		t.Errorf("TransferEncoding got '%s' want '%s'", got.TransferEncoding, want.TransferEncoding)
	}
	if want.XFrom != got.XFrom {
		t.Errorf("XFrom got '%s' want '%s'", got.XFrom, want.XFrom)
	}
	if want.XTo != got.XTo {
		t.Errorf("XTo got '%s' want '%s'", got.XTo, want.XTo)
	}
	if want.XCC != got.XCC {
		t.Errorf("XCC got '%s' want '%s'", got.XCC, want.XCC)
	}
	if want.XBCC != got.XBCC {
		t.Errorf("XBCC got '%s' want '%s'", got.XBCC, want.XBCC)
	}
	if want.XFolder != got.XFolder {
		t.Errorf("XFolder got '%s' want '%s'", got.XFolder, want.XFolder)
	}
	if want.XOrigin != got.XOrigin {
		t.Errorf("XOrigin got '%s' want '%s'", got.XOrigin, want.XOrigin)
	}
	if want.XFilename != got.XFilename {
		t.Errorf("XFilename got '%s' want '%s'", got.XFilename, want.XFilename)
	}
	if want.Body != got.Body {
		t.Errorf("Body got: \n'%s' \nwant: \n'%s'", got.Body, want.Body)
	}
}

// TestParsingFullMail tests the parsing of one full mail
func TestParsingFullMail(t *testing.T) {
	input := strings.NewReader(`Message-ID: <27713700.1075860877090.JavaMail.evans@thyme>
Date: Fri, 8 Feb 2002 08:44:58 -0800 (PST)
From: announcements.enron@enron.com
To: dl-ga-all_domestic@enron.com
Subject: Rights of Interested Parties Notice
Mime-Version: 1.0
Content-Type: text/plain; charset=us-ascii
Content-Transfer-Encoding: 7bit
X-From: Enron General Announcements </O=ENRON/OU=NA/CN=RECIPIENTS/CN=MBX_ANNCENRON>
X-To: DL-GA-all_domestic </O=ENRON/OU=NA/CN=RECIPIENTS/CN=DL-GA-all_enron_north_america>
X-cc:
X-bcc:
X-Folder: \Kevin_Hyatt_Mar2002\Hyatt, Kevin\Inbox
X-Origin: Hyatt-K
X-FileName: khyatt (Non-Privileged).pst

NOTICE:

Regulations require Enron to meet certain IRS qualification requirements, and to provide the attached information to employees participating in the qualified plans. 

In addition to attaching informational documentation here, we have also sent, via postal mail, the following notices to the names of those employees participating in the plans

1) Interested Parties' rights to Information and Rights to Notice and Comments; 
2) Notices to Interested Parties for the Enron Employee Stock Ownership Plan, 
3) The Enron Savings Plan, and 
4) The Enron Cash Balance Plan


`)

	want := Mail{
		MessageID:        "<27713700.1075860877090.JavaMail.evans@thyme>",
		Date:             "Fri, 8 Feb 2002 08:44:58 -0800 (PST)",
		From:             "announcements.enron@enron.com",
		To:               "dl-ga-all_domestic@enron.com",
		Subject:          "Rights of Interested Parties Notice",
		MimeVersion:      "1.0",
		ContentType:      "text/plain; charset=us-ascii",
		TransferEncoding: "7bit",
		XFrom:            "Enron General Announcements </O=ENRON/OU=NA/CN=RECIPIENTS/CN=MBX_ANNCENRON>",
		XTo:              "DL-GA-all_domestic </O=ENRON/OU=NA/CN=RECIPIENTS/CN=DL-GA-all_enron_north_america>",
		XCC:              "",
		XBCC:             "",
		XFolder:          "\\Kevin_Hyatt_Mar2002\\Hyatt, Kevin\\Inbox",
		XOrigin:          "Hyatt-K",
		XFilename:        "khyatt (Non-Privileged).pst",
		Body: `NOTICE:

Regulations require Enron to meet certain IRS qualification requirements, and to provide the attached information to employees participating in the qualified plans. 

In addition to attaching informational documentation here, we have also sent, via postal mail, the following notices to the names of those employees participating in the plans

1) Interested Parties' rights to Information and Rights to Notice and Comments; 
2) Notices to Interested Parties for the Enron Employee Stock Ownership Plan, 
3) The Enron Savings Plan, and 
4) The Enron Cash Balance Plan


`,
	}

	got, err := ParseMailFromReader(input)

	if err != nil {
		t.Errorf("Got error from parse %s", err.Error())
	}

	diffMails(want, got, t)
}

// TestOnlyHeaderSplitNewlineIsCut tests that only the newline that divides the header from the body is cut.
// That is to say that any newlines that are actually part of the body that are at the beginning aren't erroneously removed
func TestOnlyHeaderSplitNewlineIsCut(t *testing.T) {
	input := strings.NewReader(`Date: BogusDate
From: Example@gmail.com



This is the body`)

	want := Mail{
		Date: "BogusDate",
		From: "Example@gmail.com",
		Body: `

This is the body`,
	}

	got, err := ParseMailFromReader(input)

	if err != nil {
		t.Errorf("Got error from parse %s", err.Error())
	}

	diffMails(want, got, t)
}

func TestThrowErrorWhenGivenFaultyHeader(t *testing.T) {
	input := strings.NewReader("FalseHeader: BogusInfo")

	_, err := ParseMailFromReader(input)

	if err == nil {
		t.Errorf("Parse did not fail on faulty header")
	}
}

func TestEmptyFileShouldReturnEmptyMail(t *testing.T) {
	input := strings.NewReader("")

	got, err := ParseMailFromReader(input)

	if err != nil {
		t.Errorf("Got error from parse %s", err.Error())
	}

	want := Mail{}

	diffMails(want, got, t)
}

func TestRecognizeCcAndBccHeader(t *testing.T) {
	input := strings.NewReader("Cc: This is a CC header\nBcc: This is a BCC header")

	got, err := ParseMailFromReader(input)

	if err != nil {
		t.Errorf("Got error from parse %s", err.Error())
	}

	want := Mail{
		Cc:  "This is a CC header",
		Bcc: "This is a BCC header",
	}

	diffMails(want, got, t)
}

func TestHandleMultilineHeader(t *testing.T) {
	input := strings.NewReader("To: Example@example.com, \n\tGriveous@example.com, Harry@example.com")

	got, err := ParseMailFromReader(input)

	if err != nil {
		t.Errorf("Got error from parse %s", err.Error())
	}

	want := Mail{
		To: "Example@example.com, Griveous@example.com, Harry@example.com",
	}

	diffMails(want, got, t)
}

func TestHandleLargeMultilineHeader(t *testing.T) {
	input := strings.NewReader("To: Example@example.com, \n\tGriveous@example.com, \n\tHarry@example.com")

	got, err := ParseMailFromReader(input)

	if err != nil {
		t.Errorf("Got error from parse %s", err.Error())
	}

	want := Mail{
		To: "Example@example.com, Griveous@example.com, Harry@example.com",
	}

	diffMails(want, got, t)
}

func TestHandleMultilineHeaderWithSingleSpace(t *testing.T) {
	input := strings.NewReader("To: Example@example.com, \n Griveous@example.com, Harry@example.com")

	got, err := ParseMailFromReader(input)

	if err != nil {
		t.Errorf("Got error from parse %s", err.Error())
	}

	want := Mail{
		To: "Example@example.com, Griveous@example.com, Harry@example.com",
	}

	diffMails(want, got, t)
}

func TestHandleMultilineHeaderWithMultiSpace(t *testing.T) {
	input := strings.NewReader("To: Example@example.com, \n   Griveous@example.com, Harry@example.com")

	got, err := ParseMailFromReader(input)

	if err != nil {
		t.Errorf("Got error from parse %s", err.Error())
	}

	want := Mail{
		To: "Example@example.com, Griveous@example.com, Harry@example.com",
	}

	diffMails(want, got, t)
}
