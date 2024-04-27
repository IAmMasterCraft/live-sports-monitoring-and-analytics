package logger

import (
    "log"
    "os"
)

// A simple wrapper around the standard log package
var (
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
)

func init() {
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Failed to open log file:", err)
    }

    Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    Warning = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
