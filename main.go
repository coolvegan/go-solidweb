package main

import (
	"flag"
)

func main() {
	ip_and_port := ":9999"
	app := NewAppContext()
	flag.StringVar(&ip_and_port, "h", ":9999", "host and port eg.: 127.0.0.1:9999")
	flag.Parse()

	app.run(ip_and_port)
}
