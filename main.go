package main

import (
	"fmt"
	"strings"
	"time"

	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting clipboard flattener")
	// loop that runs forever
	for {
		// get clipboard contents
		contents := clipboard.Read(clipboard.FmtText)
		// check if clipboard contents are empty
		if contents == nil {
			// if so, skip this iteration
			continue
		}
		// convert clipboard contents to strings
		content_string := string(contents)
		// check if clipboard contents are empty
		if content_string != "" {
			// check if clipboard contents are already flattened
			if strings.Contains(content_string, "\n") {
				// if not, flatten them
				content_string = strings.Replace(content_string, "\r", " ", -1)
				content_string = strings.Replace(content_string, "\n", " ", -1)
				content_string = strings.Replace(content_string, "  ", " ", -1)
				// write flattened contents back to clipboard
				clipboard.Write(clipboard.FmtText, []byte(content_string))
				fmt.Println("Flattened clipboard contents")
			}
		}
		time.Sleep(1 * time.Second)
	}
}
