package main

import (
        "github.com/tarm/goserial"
        "log"
        "bufio"
        "fmt"
)

var GREEN = "\033[92m"
var BLUE = "\033[94m"

func main() {
        c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
        s, err := serial.OpenPort(c)
        if err != nil {
                log.Fatal(err)
        }

        reader := bufio.NewReader(s)
        go func() {
        	for {
		        line, err := reader.ReadString('\n')
		        if err != nil {
		        	log.Fatal(err)
		        }
		        log.Printf("%s%s%s", BLUE, line, GREEN)
		    }
        }()
        
        var cmd string

        for cmd != "quit" {
        	_, err := fmt.Scan(&cmd)
        	if err != nil {
        		log.Fatal(err)
        	}
        	_, err = s.Write([]byte(fmt.Sprintf("%s\n",cmd)))
		    if err != nil {
	            log.Fatal(err)
	    	}
        }
}