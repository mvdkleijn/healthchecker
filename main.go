/*
	Healthchecker - A simple, healthcheck utility.
    Copyright (C) 2023  Martijn van der Kleijn

	This file is part of Healthchecker.

    Healthchecker is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    Healthchecker is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with Healthchecker.  If not, see <http://www.gnu.org/licenses/>.
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
