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
			log.Fatal(e)
		}

		e = cmd.Start()
		if e != nil {
			log.Fatal(e)
		}


		str := ""

		buf := make([]byte, 128)
		for {
			size, e := stdoutPipe.Read(buf)
			if e == io.EOF {
				break
			} else if e != nil {
				log.Fatal(e)
			}
			str = str + string(buf[0:size])
		}
		inputReq.Retc <- str

	}
	} ()
	return retc
}
