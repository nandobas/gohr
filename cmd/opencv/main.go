package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image/color"
)

func main() {

	webcam, err := gocv.VideoCaptureDevice(0)

	if err != nil {
		fmt.Println("Error on open webcam")
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Detector Bas 1.0")
	defer window.Close()

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	classifier.Load("haarcascade_frontalface_default.xml")

	for {
		img := gocv.NewMat()
		if ok := webcam.Read(&img); !ok {
			fmt.Println("cannot read webcam")
			return
		}
		if img.Empty() {
			continue
		}

		myFace := classifier.DetectMultiScale(img)
		for _, r := range myFace {
			gocv.Rectangle(&img, r, color.RGBA{255, 0, 0, 0}, 2)
		}

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
