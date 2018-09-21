package main

import (
	"log"
	"os"
)

var Log *log.Logger = log.New(os.Stdout, "alberto ", log.Lshortfile|log.LstdFlags)
