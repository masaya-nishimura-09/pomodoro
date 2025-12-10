package main

import (
    "fmt"
    "time"
    "strings"
    "github.com/gosuri/uilive"
)

type Session struct {
}

func main() {
    time := time.Now()
    fmt.Println("Current time is:", timeToString(time))

    for i := 0; i <= 100; i++ {
        fmt.Fprintf(writer, "Downloading.. (%d/%d) GB\n", i, 100)
        time.Sleep(time.Millisecond * 5)
    }

    fmt.Fprintln(writer, "Finished: Downloaded 100GB")
    writer.Stop()
}

func timeToString(t time.Time) string {
    formattedTime := t.Format("2006-01-02 15:04:05")
    currentTimeString := strings.Split(formattedTime, " ")[1]
    return currentTimeString
}

