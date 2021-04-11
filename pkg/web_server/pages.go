package web_server

import (
	"net/http"

	_ "embed"
)

var (
	//go:embed files/login.html
	pageLogin []byte
	//go:embed files/main.html
	pageMain []byte

	//go:embed files/favicon.png
	favicon []byte
	//go:embed files/avatar.png
	avatar []byte
)

func loginPage(w http.ResponseWriter) {
	w.Write(pageLogin)
}

func mainPage(w http.ResponseWriter) {
	w.Write(pageMain)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(favicon)
}

func avatarHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(avatar)
}
