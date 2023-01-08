package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"zincSearchProject/internal/zincindex"
)

func main() {
	username := os.Getenv("ZincUsername")
	password := os.Getenv("ZincPassword")
	go zincindex.Authenticate("http://127.0.0.1:4080/", "enronmail", username, password).IndexDB(`.\resources\enron_mail_20110402\`)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println("Failed to start profiling server")
		log.Fatal(err)
	}
}
