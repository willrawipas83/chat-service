package cmd

import (
	"log"
	"net/http"
	"sync"

	"github.com/Rawipass/chat-service/config"
	https "github.com/Rawipass/chat-service/internal/http"
	"github.com/Rawipass/chat-service/logger"
	"github.com/Rawipass/chat-service/routes"
	"github.com/spf13/cobra"
)

var wg sync.WaitGroup

var ChatCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Chat Service",
	Run: func(cmd *cobra.Command, args []string) {

		// Init Logger
		logger.InitLogger()

		// Init Database
		config.ConnectDatabase()

		http.HandleFunc("/ws", https.WebSocketHandler)

		log.Println("Server started on :8080")
		go func() {
			log.Fatal(http.ListenAndServe(":8080", nil))
		}()

		r := routes.SetupRouter()
		if err := r.Run(":3001"); err != nil {
			log.Fatalf("Could not run server: %v\n", err)
		}

		logger.Logger.Info("chat service is running")

		// Waiting for Component Shut Down
		wg.Wait()

		// Flush Log
		logger.SyncLogger()
	},
}

func init() {
	rootCmd.AddCommand(ChatCmd)
}
