package main

import (
	"ddns/ddns"
	"flag"
	"log"
	"os"
	"time"
)

func main() {
	configPath := flag.String("config", "config.yaml", "config file path")
	logPath := flag.String("log", "ddns.log", "log file path")
	flag.Parse()

	err := ddns.LoadConfig(*configPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	ddns.IsInit()

	logfile, err := os.OpenFile(*logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	defer logfile.Close()

	loger := log.New(logfile, "", log.LstdFlags)

	tasker := time.NewTicker(time.Second * time.Duration(ddns.Config.ScheduledTask))
	for {
		<-tasker.C

		msg, err := ddns.RunDDNS()
		if err != nil {
			loger.Println(err)
		}
		if msg != nil {
			loger.Println(*msg)
		}
	}
}
