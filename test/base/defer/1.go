package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"io"
	"log"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {

	var dst *os.File
	var src *os.File

	defer func() {
		if err := dst.Close(); err != nil {
			fmt.Println("Close dst file error:", err)
		}
	}()

	defer func() {
		if err := src.Close(); err != nil {
			fmt.Println("Close dst file error:", err)
		}
	}()

	if src, err = os.Open(srcName); err != nil {
		return
	}

	if dst, err = os.Create(dstName); err != nil {
		return
	}

	if written, err = io.Copy(dst, src); err != nil {
		fmt.Println("Copy file error:", err)
	}

	return
}

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatalf("error opening device: %v", err)
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	window := gocv.NewWindow("webcamwindow")
	defer window.Close()

	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Println("Unable to read from the webcam")
			continue
		}

		window.IMShow(img)
		window.WaitKey(50)
	}
}
