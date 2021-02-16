package image

import (
	"GoMy/file"
	"errors"
	"fmt"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// 颜色设置
type setColorImage interface {
	Set(x, y int, c color.Color)
}

// 打开图片
func Open(path string) (*Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// 获取文件类型
	ft, err := file.GetFileType(f)
	if err != nil {
		return nil, err
	}
	// 打开图片
	var img image.Image
	switch ft {
	case "jpg":
		img, err = jpeg.Decode(f)
	case "png":
		img, err = png.Decode(f)
	case "gif":
		img, err = gif.Decode(f)
	case "bmp":
		img, err = bmp.Decode(f)
	default:
		return nil, errors.New("can not open image which type is " + ft)
	}
	if err != nil {
		return nil, err
	}
	result := &Image{image: img}
	return result, nil
}

// 创建
func Convert(img image.Image) *Image {
	result := &Image{
		image: img,
	}
	return result
}

// 图片
type Image struct {
	image image.Image // 解码的图片
}

// 获取图片尺寸, width/height
func (this *Image) GetSize() (int, int) {
	point := this.image.Bounds().Size()
	return point.X, point.Y
}

// 获取像素点颜色
func (this *Image) GetColor(x, y int) color.Color {
	return this.image.At(x, y)
}

// 设置像素点颜色
func (this *Image) SetColor(x, y int, clr color.Color) error {
	if value, ok := this.image.(setColorImage); ok {
		value.Set(x, y, clr)
	} else {
		return errors.New("the type of image does not support changing color")
	}
	return nil
}

// 保存
func (this *Image) Save(path string) error {
	// 获取格式
	filename := filepath.Base(path)
	pointIndex := strings.LastIndex(filename, ".")
	if pointIndex < 0 {
		return errors.New("do not define format, please add string eg.'.jpg'")
	}
	format := filename[pointIndex+1:]
	// 根据格式生成文件
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return err
	}
	switch format {
	case "jpg":
		err = jpeg.Encode(f, this.image, &jpeg.Options{Quality: 100})
	case "png":
		err = png.Encode(f, this.image)
	case "gif":
		err = gif.Encode(f, this.image, &gif.Options{NumColors: 256, Quantizer: nil, Drawer: nil})
	case "bmp":
		err = bmp.Encode(f, this.image)
	default:
		if err = os.Remove(path); err == nil {
			err = errors.New(fmt.Sprintf("the type \"%s\" is not supported", format))
		}
	}
	if err != nil {
		return err
	}
	return nil
}
