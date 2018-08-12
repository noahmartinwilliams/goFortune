package goFortune

import "log"
import . "os/exec"
import "io"

// The data type that passes requests for a fortune.
type FortuneRequest struct {
	// The options to pass to fortune for this request
	FortuneOpts string
	// A channel that returns the string from fortune
	Retc chan string
}

// Fortune() returns a channel which can be used to send fortune requests.
func Fortune() chan FortuneRequest {
	retc := make (chan FortuneRequest, 100)
	go func() {
	for {
		inputReq, ok := <-retc
		if !ok {
			return
		}
		cmd := Command("sh", "-c", "fortune " + inputReq.FortuneOpts)
		stdoutPipe ,e := cmd.StdoutPipe()

		if e != nil {
			log.Println(e)
		}

		e = cmd.Start()
		if e != nil {
			log.Println(e)
		}


		str := ""

		buf := make([]byte, 128)
		for {
			size, e := stdoutPipe.Read(buf)
			if e == io.EOF {
				break
			} else if e != nil {
				log.Println(e)
			}
			str = str + string(buf[0:size])
		}
		inputReq.Retc <- str

	}
	} ()
	return retc
}

// FortuneStream takes a string contained space separated arguments to pass to fortune 
// and returns a channel which will continuously return fortunes.
func FortuneStream(args string) chan string {
	retc := make(chan string, 10)
	go func() {
		fortuneReq := FortuneRequest{FortuneOpts:args, Retc:retc}
		reqc := Fortune()
		defer close(reqc)
		defer close(retc)

		for {
			reqc <- fortuneReq
		}
	} ()
	return retc
}
