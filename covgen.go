// Capture 1 frame from video
// Захватываем один кадр
// example: ./covgen /var/tmp/in.mp4 /var/tmp/out.jpg
// example: go run ./covgen.go ./in.mp4  ./out.jpg
// out.jpg generated from video
// go run ./covgen.go ./in.mp4  ./out1.jpg
// ./covgen ./in.mp4  ./out2.jpg

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 3 {
		// /var/tmp/in.mp4
		vidName := os.Args[1]
		// /var/tmp/in.jpg
		//imgName := os.Args[2]
		outImg := os.Args[2]

		f, err := os.OpenFile(vidName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 777)
		err = f.Close()
		if err != nil {
			//fmt.Printf("%s File not found!!!\n", vidName)
			log.Fatal(err)
		}

		// ffmpeg -y -ss 00:01:00 -i in.mp4 -frames:v 1 -q:v 2 output.jpg
		app := "./ffmpeg"
		arg0 := "-y"
		arg1 := "-ss"
		arg2 := "00:00:01"
		arg3 := "-i"
		arg4 := vidName
		arg5 := "-vf"
		arg6 := "scale=200:200"
		arg7 := "-frames:v"
		arg8 := "1"
		arg9 := "-q:v"
		arg10 := "2"
		arg11 := outImg

		cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Enter the name of the video file and output image!!!")
	}
}
