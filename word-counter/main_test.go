package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
 b := bytes.NewBufferString("word 1 word2 word3\nline2\nline3 word1")

 exp := 3

 res := count(b, true)
 
 if res != exp {
	 t.Errorf("Expected %d, got %d intead. \n", exp, res)
 }
}