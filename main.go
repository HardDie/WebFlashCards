package main

import (
	"log"
	"net/http"

	boltWrapper "github.com/HardDie/WebFlashCards/pkg/bolt_wrapper"
	webServer "github.com/HardDie/WebFlashCards/pkg/web_server"
)

const (
	DBPath = "/tmp/mytmp/test.db"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	// Open DB with words
	bw, err := boltWrapper.NewBoltWrapperOpen(DBPath)
	if err != nil {
		if err == boltWrapper.ErrorDBNotExists {
			log.Println("Can't find exists DB, try to create new one:", DBPath)
			bw, err = boltWrapper.NewBoltWrapperCreate(DBPath)
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("DB", DBPath, "was successfully created")
	} else {
		log.Println("DB", DBPath, "was opened")
	}
	_ = bw

	
	http.Handle("/", webServer.NewLoginPageHandler())
	http.Handle("/favicon.ico", webServer.NewFaviconHandler())
	http.Handle("/avatar.png", webServer.NewAvatarHandler())

	log.Println("Start serving on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
