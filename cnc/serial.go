package cnc

import (
        "github.com/tarm/goserial"
        "log"
        "io"
        "bufio"
        "fmt"
)


// this defines the communication on the serial port to the arduino

type SerialRW struct {
	reads chan string
	writes chan string
	connection io.ReadWriteCloser
}

func (this *SerialRW) Write(command string) {
	this.writes <- command
}

func (this *SerialRW) Connect(path string, baud int) error {
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
    s, err := serial.OpenPort(c)
    if err != nil {
        return err
    }

    reader := bufio.NewReader(s)
    //reader goroutine
    go func() {
    	for {
	        line, err := reader.ReadString('\n')
	        if err != nil {
	        	log.Fatal(err)
	        }
	        log.Printf("%s%s%s", BLUE, line, GREEN)
	        this.reads <- line
	    }
    }()
    
    //writer goroutine
    go func() {
    	for {
    		select {
    		case cmd := <- this.writes :
    			if cmd == "quit" {
    				log.Println("Quit called")
    				return
    			}
    			_, err = s.Write([]byte(fmt.Sprintf("%s\n",cmd)))
			    if err != nil {
		            log.Fatal(err)
		            return
		    	}
    		}
    	}
    }()
}

