package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGenerate(t *testing.T) {
	err := generate("./test_data/in.mp4", "/tmp/out_"+fmt.Sprint(rand.Int())+".jpg")
	if err != nil {
		t.Error(err)
	}
}
