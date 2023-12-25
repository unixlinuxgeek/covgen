package main

import (
	"log"
	"testing"
)

func TestGenerate(t *testing.T) {
	err := generate("./test_data/in.mp4", "./out11.jpg")
	if err != nil {
		log.Fatal(err)
	}
}
