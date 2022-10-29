package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Command(args []string) {
	switch args[0] {
	case "serve":
		serve()
		break
	case "migrations":
		migrations()
		break
	}

}

func serve() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load .env file")
	}
	fmt.Println("Get .Env Value . . .")
	fmt.Println("=========================")
	fmt.Println("Server started on port " + os.Getenv("PORT"))

	restAPI := InitializeRestAPI()
	http.ListenAndServe(":"+os.Getenv("PORT"), restAPI)
}
func migrations() {

}
