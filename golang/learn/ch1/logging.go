package main

import (
	"log"
	"log/syslog"
	"os"
)

/*
All UNIX systems have their own log files for writing logging information that comes from running
servers and programs. Usually, most system log files of a UNIX system can be found
under the /var/log directory.

The UNIX logging service has support for two properties named *logging level* and *logging facility*.

The logging level is a value that specifies the severity of the log entry
while logging facility is like a category used for logging information

 - logging level:  debug, info, notice, warning, err, crit, alert, and emerg
 - logging facility: auth, authpriv, cron, daemon, kern, lpr, mail, mark, news, syslog, user, UUCP, local0, local1, local2, local3, local4, local5, local6, or local7

 Any log sent must specify it's logging facility else it will be lost.

 The Go log package writes to standard error (as it's logging facility) but this can be changed using log. SetOutput()
*/
func main() {
	// Writing to the main system log file (/var/logs/system.log) is as easy as calling and not standard error
	//"systemLog.go" is the custom text describing the log
	// appears thus Sep 15 23:10:42 Toyosis-MacBook-Pro systemLog.go[95272]: 2021/09/15 23:10:42 Everything is fine!
	_, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")

	if err != nil {
		log.Println(err)
		return
	} else {
		// log.SetOutput(sysLog)
		log.Println("Everything is fine!")
	}

	fatapPanicLogs()

	log.Println("The end of main function")
}

func fatapPanicLogs() {
	if len(os.Args) != 1 {
		log.Fatal("Fatal: Hello World!")
	}

	log.Panic("Panic: Hello World!")
}
