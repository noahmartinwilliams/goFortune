package goFortune

import "testing"

func TestFortune(t *testing.T) {
	c := Fortune()

	retc := make(chan string)
	request := FortuneRequest{FortuneOpts:"./test", Retc:retc}

	c <- request
	output := <-retc

	if output != "hello, world\n" && output != "goodbye, world\n" {
		t.Errorf("Fortune did not output a valid response. Got: " + output)
	}
}
