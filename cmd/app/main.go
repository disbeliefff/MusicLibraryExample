package main

import (
	"MusicLibrary/api/server"
	"MusicLibrary/cfg"
	"fmt"
)

func main() {
	conf, err := cfg.LoadConfig("./cfg/app.env")
	if err != nil {
		panic(err)
	}

	fmt.Println(conf)

	s := server.NewServer(conf)

	err = s.Start()
	if err != nil {
		fmt.Printf("cannot start server: %v", err)
	}
}

//func initLogger() *slog.Logger {
//
//}
