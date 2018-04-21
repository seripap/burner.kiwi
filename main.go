package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/apex/gateway"
	"github.com/gorilla/context"
	"github.com/haydenwoodhead/burnerkiwi/server"
)

func main() {
	useLambda, err := strconv.ParseBool(os.Getenv("LAMBDA"))

	if err != nil {
		log.Fatalf("Failed to parse LAMBDA env var. Err = %v", err)
	}

	key := os.Getenv("KEY")

	if strings.Compare(key, "") == 0 {
		log.Fatalf("Env var key cannot be empty")
	}

	websiteURL := os.Getenv("WEBSITE_URL")

	if strings.Compare(websiteURL, "") == 0 {
		log.Fatalf("Env var WEBSITE_URL cannot be empty")
	}

	staticURL := os.Getenv("STATIC_URL")

	if strings.Compare(websiteURL, "") == 0 {
		log.Fatalf("Env var STATIC_URL cannot be empty")
	}

	mgKey := os.Getenv("MG_KEY")

	if strings.Compare(mgKey, "") == 0 {
		log.Fatalf("Env var MG_KEY cannot be empty")
	}

	mgDomain := os.Getenv("MG_DOMAIN")

	if strings.Compare(mgKey, "") == 0 {
		log.Fatalf("Env var MG_KEY cannot be empty")
	}

	var developing bool

	debuggingENV := os.Getenv("DEBUGGING")

	if strings.Compare(mgKey, "") == 0 {
		developing = false
	} else {
		developing, err = strconv.ParseBool(debuggingENV)

		if err != nil {
			log.Fatalf("Failed to parse debugging: %v", err)
		}
	}

	s, err := server.NewServer(key, websiteURL, staticURL, mgDomain, mgKey, []string{"rogerin.space"}, developing)

	if err != nil {
		log.Fatalf("Failed to setup new server: %v", err)
	}

	if useLambda {
		log.Fatal(gateway.ListenAndServe("", context.ClearHandler(s.Router))) // wrap mux in ClearHandler as per docs to prevent leaking memory
	} else {
		log.Fatal(http.ListenAndServe(":8080", context.ClearHandler(s.Router)))
	}
}
