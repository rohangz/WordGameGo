package main

import ("fmt"
		"log"
		"os"
		"bufio"
		"net")
func main() {
	conn,err:=net.Dial("tcp","localhost:3000")
	if err!=nil {
		log.Fatalln(err)
	}
	reader:=bufio.NewReader(os.Stdin)
	for{
	fmt.Print("Enter Text ")
	text,_:=reader.ReadString('\n')
	_,err = conn.Write([]byte(text))
	if err!=nil {
		log.Fatalln(err)
	}
	text=text[:len(text)-1]
	fmt.Print(text,len(text))
	if text=="exit" {
		break
	}

	}
}