package main

import (
	"os"
	"os/signal"
	"log"
	"errors"
	"net/http"
	"context"
	"syscall"
	"time"
	"github.com/julienschmidt/httprouter"
)

func newRouter() *httprouter.Router {
	mux := httprouter.New()

	mux.GET("/youtube/channel/stats", getChannelStats())

	return mux

}


// getChannelStats defines the handler function for the "/youtube/channel/stats" endpoint
// It writes a response with the text "response!" to the HTTP response writer. 

func getChannelStats() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Write([]byte("response!"))
	}
}


func main() {
	// Create a new HTTP server.
	srv := &http.Server{
		Addr: ":10101",		// Listen on port 10101.
		Handler: newRouter(),	// Use the router returned by newRouter()
	}

	idleConnsClosed := make(chan struct{})

	// Start a goroutine to handle server shutdown
	go func() {
		// Listen for OS interrupt signals (e.g. Ctrl+C)
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		log.Println("service interrupt received")i

		// Create a context with a timeout of 60 seconds for server shutdown.
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()


		// Shut down the server gracefully	
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("http server shutdown error: %v", err)
		}

		log.Println("shutdown complete")

		close(idleConnsClosed)

	}()

	log.Printf("Starting server on port 10101")

	// Start the HTTP server and listen for incoming requests.
	if err := srv.ListenAndServe(); err != nil {
		// Check if server was closed intentionally.
		if !errors.Is(err, http.ErrServerClosed){
			log.Fatalf("fatal http server failed to start: %v", err)
		
		}
	
	}

	<-idleConnsClosed
	log.Println("Service Stop")



}
