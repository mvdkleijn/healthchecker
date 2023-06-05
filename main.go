/*
	Copyright (C) 2023  Martijn van der Kleijn

	This file is part of the healthchecker application.

    This Source Code Form is subject to the terms of the Mozilla Public
  	License, v. 2.0. If a copy of the MPL was not distributed with this
  	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Ensure that a URL is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <URL>")
		os.Exit(1)
	}

	// Get the URL from command-line argument
	url := os.Args[1]

	// Make a HEAD request to the provided URL
	resp, err := http.Head(url)
	if err != nil {
		fmt.Printf("Error making HEAD request: %s\n", err)
		os.Exit(1)
	}

	// Check the HTTP status code
	if resp.StatusCode == http.StatusOK {
		fmt.Println("HTTP status code 200 OK")
		os.Exit(0)
	} else {
		fmt.Printf("HTTP status code: %d\n", resp.StatusCode)
		os.Exit(1)
	}
}
