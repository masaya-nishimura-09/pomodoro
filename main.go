package main

import (
	"fmt"
	"os"

    "github.com/masaya-nishimura-09/pomodoro/notification"
    "github.com/masaya-nishimura-09/pomodoro/timer"
    "github.com/masaya-nishimura-09/pomodoro/model"
)

func main() {
    session := model.Session{}

    focus:=  model.Timer{
        Type: 0,
        Time: 1, 
        Data: getAsciiData("ascii/gopher_focus.txt"), 
    }
    shortBreak:=  model.Timer{ 
        Type: 1,
        Time: 1, 
        Data: getAsciiData("ascii/gopher_short_break.txt"), 
    }
    longBreak:= model.Timer{ 
        Type: 2,
        Time: 15, 
        Data: getAsciiData("ascii/gopher_long_break.txt"), 
    }

    for {
        session.AddCount()

        switch session.Count {
        case 1, 3, 5:
            timer.Timer(focus)
            notification.ShortBreak()
        case 2, 4, 6:
            timer.Timer(shortBreak)
            notification.Focus()
        case 7:
            timer.Timer(focus)
            notification.LongBreak()
        case 8:
            timer.Timer(longBreak)
            notification.Focus()
        }
    }
}

func getAsciiData(filePath string) []byte {
    f, textOpenErr := os.Open(filePath)
    if textOpenErr != nil {
        fmt.Println(textOpenErr)
        fmt.Println("fail to open file")
    }

    data := make([]byte, 4000)
    _, fileReadErr := f.Read(data)
    if fileReadErr != nil {
		fmt.Println(fileReadErr)
		fmt.Println("fail to read file")
	}

	return data
}
