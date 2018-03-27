package main

import ("fmt"
		"net"
		"log"
		"strconv"
		)

func main() {
	var started int = 0
	var count int = 0
	var check int = 0

	//turnToGiveWord:=0 //turn of the player who is giving the word 
	//turnToAnswer:=1//turn of the player who is going to answer the word 
	for {

		if started==0 {
			started = 1
			go ListenToConnections(&count)
		}
		if count>0 {
			check = 1
		}
		if check == 1 {
			if count==0 {
				fmt.Println("Time to terminate Connection ")
				break
			}
		}
	}

}

func ListenToConnections(count *int) {
	listener,err:=net.Listen("tcp","localhost:3000")
	if err!=nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	turnToGiveWord:=0
	turnToAnswer:=1

	for 
	{
		if *count == 3 {
			break
		}
		conn,err:= listener.Accept()
		if err!=nil {
			log.Fatalln(err)
		}
		fmt.Println("New Connection Initiated")
		go listenConnection(conn,count,&turnToGiveWord,&turnToAnswer)
		*count=*count+1
		fmt.Println(*count,"users are connected ")
		
	}
	
}
func listenConnection(conn net.Conn,count,turnToGiveWord,turnToAnswer *int) {
	
	index := *count //id 
	fmt.Println(index,"is my index ")
	var packet string = strconv.Itoa(index)+"$"+strconv.Itoa(*turnToGiveWord)+"$"+strconv.Itoa(*turnToAnswer)
	
	fmt.Println("Packet is",packet)
	conn.Write([]byte(packet))
	for {
		buffer:= make([]byte,1400)
		dataSize, err := conn.Read(buffer)
		if err!=nil {
			log.Fatalln(err)
			return 
		}
		data:=buffer[:dataSize-1]

		fmt.Println(string(data))
		//	_,err = conn.Write(data)
		if err!=nil {
			log.Fatalln(err)
		}
		if string(data) == "exit" {
			fmt.Println(string(*count),"Users are Connected\n")
			*count = *count - 1
			break 
		}
		conn.Write([]byte(string(*count)))
		
	}
	
}
