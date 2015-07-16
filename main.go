package main

import (
	"os"
	"syscall"
	"strconv"
	"time"
	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "KUWASHIMA Yuichiro"
	app.Email = "ykuwashima@gmail.com"
	app.Usage = "viagra [-d minutes] PID"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Action = func(c *cli.Context) {
		var err error
		var min int64
		min, err = strconv.ParseInt(c.String("duration"),10,64)
		if err != nil {
			println("Duration is not number")
		}

		if len(c.Args()) > 0 {
			var proc int
			var preNice int
			proc, err = strconv.Atoi(c.Args()[len(c.Args())-1])
			if err != nil {
				println("Proc no parse error")
				return
			}
			preNice, err = syscall.Getpriority(syscall.PRIO_PROCESS, proc);
			if err != nil {
				println("Unknown PID")
				return
			}
			err = syscall.Setpriority(syscall.PRIO_PROCESS, proc, 20);
			if err != nil {
				println("Setpriority failed")
				return
			}
			println("Power up!")
			timer := time.NewTimer(time.Duration(min) * time.Second)
			<- timer.C
			err = syscall.Setpriority(syscall.PRIO_PROCESS, proc, preNice);
			println("Time is up!")
		}
	}

	app.Run(os.Args)
}
