package zincindex

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
	Cc               string `json:"cc"`
	Bcc              string `json:"bcc"`
	Body             string `json:"body"`
}
