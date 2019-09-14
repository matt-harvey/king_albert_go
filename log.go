package main

import (
	"log"
	"os"
)

var Log *log.Logger = log.New(os.Stdout, "king_albert_go ", log.Lshortfile|log.LstdFlags)
