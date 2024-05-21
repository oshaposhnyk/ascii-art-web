package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	host := flag.String("h", "localhost", "Server host")
	port := flag.Int("p", 8080, "Server port")
	flag.Parse()

	r := gin.Default()
	r.GET("/", rootHandler)
	r.POST("/ascii", convertStringHandler)
	err := r.Run(fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
