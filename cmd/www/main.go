package main

import (
	"github.com/labstack/echo/v4"
	//"github.com/bmizerany/pat"
)

func main() {
	// Initialize a router and add the path and handler for the homepage.
	//mux := pat.New()
	//mux.Get("/:locale", http.HandlerFunc(handleHome))
	mux := echo.New()
	mux.GET("/:local", handleEcho)

	// Start the HTTP server using the router.
	//log.Println("starting server on :4018...")
	//err := http.ListenAndServe(":4018", mux)
	//log.Fatal(err)
	mux.Logger.Fatal(mux.Start(":4018"))
}