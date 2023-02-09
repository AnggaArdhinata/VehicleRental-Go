package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AnggaArdhinata/backend-go/src/routers"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "server",
	Short: "For running api server",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	mainRoute, err := routers.NewApp()
	if err != nil {
		return err
	}
	handler := cors.AllowAll().Handler(mainRoute)

	var addrs string = "localhost:8080"
	if port := os.Getenv("APP_PORT"); port != "" {
		addrs = port
	}

	fmt.Println("== Welcome to Vehicle Rental Back-End Golang App ==")
	log.Println("app running on port", addrs)
	return http.ListenAndServe(addrs, handler)

}
