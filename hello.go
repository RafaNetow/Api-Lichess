package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Configura el manejador de solicitudes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Hola, mundo!")
	})

	// Inicia el servidor en el puerto 8080
	fmt.Println("Servidor iniciado en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}