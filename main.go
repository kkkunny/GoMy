package main

import (
	myimage "GoMy/image"
	"image"
)

func main() {
	temp := image.NewRGBA(image.Rect(0, 0, 100, 100))
	img := myimage.Convert(temp)
	if err := img.Save("./result.bmp"); err != nil {
		panic(err)
	}
}
