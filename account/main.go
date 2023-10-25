package main

import (
    "context"
    "os"
    "os/signal"
    "syscall"
    "time"
    "net/http"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jj-attaq/api-todo/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnDB()
}

func main() {
    log.Println("Starting server...")
    router := gin.Default()

    router.GET("/api/account", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "hello": "there",
        })
    })

	// Graceful shutdown
	port := os.Getenv("PORT")
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", server.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Printf("\nShutting down server...")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
