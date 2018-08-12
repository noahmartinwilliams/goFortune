package goFortune

import "log"
import . "os/exec"
import "io"

type FortuneRequest struct {
	FortuneOpts string
	Retc chan string
}

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
