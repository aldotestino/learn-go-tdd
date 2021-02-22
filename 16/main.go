package main

import (
	svg "16/svg"
	"os"
	"time"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
