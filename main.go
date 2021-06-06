package main

import (
	"fmt"
	"time"
)

const (
	logInfo = "INFO"
	logWarning = "WARNING"
	logError = "ERROR"
)

type logEntry struct {
	time time.Time
	severity string
	message string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // Using empty struct is similar to using a bool channel but requires no memory. A performance convention.

func main() {
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	doneCh <- struct{}{} // Pass an empty struct to signal end
	time.Sleep(100 * time.Millisecond)
}

func logger() {
	for {
		select { // Blocks until a message is received on one of the channels
		case entry := <-logCh:
			fmt.Printf("%v - [%v] %s\n", entry.time.Format("2020-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
		// NOTE: If you add a default statement it won't block
	}
}
