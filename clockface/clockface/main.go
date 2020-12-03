package main

import (
	"os"
	"time"
	"go-with-tests/Clockface"
)

func main() {
    t := time.Now()
    clockface.SVGWriter(os.Stdout, t)   
}
