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

func TestFortuneStream(t *testing.T) {
	sc := FortuneStream("./test")
	str := <-sc

	if str != "hello, world\n" && str != "goodbye, world\n" {
		t.Errorf("Error: FortuneStream did not return valid fortune. Got: " + str)
	}

	str = <-sc
	if str != "hello, world\n" && str != "goodbye, world\n" {
		t.Errorf("Error: FortuneStream did not return valid fortune. Got: " + str)
	}
}
