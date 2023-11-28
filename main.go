package main

import (
	"fmt"
	"log"
	"os/exec"
)


func main() {
	x := 0
	for {
		//out1, _ := exec.Command("cd /var/test/").Output()
		//out2, _ := exec.Command("mv 163535-1608224541.pdf /var/test/").Output()
		//out3, _ := exec.Command("cd /var/test/").Output()
		out, err := exec.Command("convert", " -density 300", "f -quality 100", "/var/test/163535-1608224541.pdf", fmt.Sprintf("/var/test/pic-%d.jpg", x)).Output()
		if err != nil {
			log.Println(err.Error())
		}
	//	op1 := string(out1[:])
	//	log.Println(op1)
		//op2 := string(out2[:])
		//log.Println(op2)
		//op3 := string(out3[:])
		//log.Println(op3)
		output := string(out[:])
		log.Println(output)
		x++
	}
}
