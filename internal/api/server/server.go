package server

import (
	"fmt"
	"net/http"

	"wtg_desktop/internal/config"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola desde Wails backend API")
}

func InitServer() {
	port := config.GetServerPort()

	http.HandleFunc("/api/hello", helloHandler)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
		if err != nil {
			panic("Servidor falló: " + err.Error())
		}
	}()
}
