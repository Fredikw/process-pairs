// Use 'go tun main.go' to run your program
// Use 'go build' to make an executable

package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {

	stamp := 0
	prev := 0
	
	/* 
	Check if a primary exists. Read communication file and check if time_stamp updated 
	
	*/
	for {

		infile, err := ioutil.ReadFile("phoenix_com.txt")

		if err != nil {
			fmt.Println(err)
		}
		time_stamp, err := strconv.Atoi(strings.Split(string(infile), "\n")[1])

		if time_stamp == prev {
			break
		} else {
			prev = time_stamp
		}
		time.Sleep(100 * time.Millisecond)
	}

	/* 
	Spawn backup 
	
	*/
	cmd := exec.Command("./main", ".")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	/* 
	Set count to latest value stored in phoenix_com.txt 
	
	*/
	infile, err := ioutil.ReadFile("phoenix_com.txt")

	if err != nil {
		fmt.Println(err)
	}
	count, err := strconv.Atoi(strings.Split(string(infile), "\n")[0])

	/*
	Update count and stamp, and write to file 
	
	*/
	for {
		stamp = stamp + 1

		if stamp%10 == 0 {
			count = count + 1
			fmt.Println(count)
		}
		s := strconv.Itoa(count) + "\n" + strconv.Itoa(stamp)

		mydata := []byte(s)

		err := ioutil.WriteFile("phoenix_com.txt", mydata, 0777)

		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
