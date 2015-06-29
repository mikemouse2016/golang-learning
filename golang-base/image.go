package main

import (
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

//读取jpeg图像并返回image.Image对象
func ImageJPEG(path string) (image.Image, error) {
	//打开图像文件
	f, fileErr := os.Open(path)
	if fileErr != nil {
		return nil, fileErr
	}
	//退出时关闭文件
	defer f.Close()
	//解码
	j, jErr := jpeg.Decode(f)
	if jErr != nil {
		return nil, jErr
	}
	//返回解码后图片
	return j, nil
}

//读取png图像并返回image.Image对象
func ImagePNG(path string) (image.Image, error) {
	//打开图像文件
	f, fileErr := os.Open(path)
	if fileErr != nil {
		return nil, fileErr
	}
	//退出时关闭文件
	defer f.Close()
	//解码
	j, jErr := png.Decode(f)
	if jErr != nil {
		return nil, jErr
	}
	//返回解码后图片
	return j, nil
}

//读取gif图像并返回image.Image对象
func ImageGIF(path string) (image.Image, error) {
	//打开图像文件
	f, fileErr := os.Open(path)
	if fileErr != nil {
		return nil, fileErr
	}
	//退出时关闭文件
	defer f.Close()
	//解码
	j, jErr := gif.Decode(f)
	if jErr != nil {
		return nil, jErr
	}
	//返回解码后图片
	return j, nil
}

//按照分辨率创建一张空白图片对象
func ImageRGBA(width, height int) *image.RGBA {
	//建立图像，image.Rect(最小x，最小y，最大x， 最小y)
	//func Rect(x0, y0, x1, y1 int) Rectangle
	//返回一个矩形Rectangle{Pt(x0, y0), Pt(x1, y1)}。
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

//将图片绘制到图片
func ImageDrawRGBA(img *image.RGBA, imgcode image.Image, x, y int) {
	//绘制图像
	// image.Point A点的X,Y坐标,轴向右和向下增加{0,0}
	// image.ZP ZP is the zero Point
	// image.Pt Pt is shorthand for Point{X, Y}
	//func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
	draw.Draw(img, img.Bounds(), imgcode, image.Pt(x, y), draw.Over)
}

// JPEG将编码生成图片
// 选择编码参数,质量范围从1到100,更高的是更好 &jpeg.Options{90}
func ImageEncodeJPEG(path string, img image.Image, option int) error {
	//
	f, fileErr := os.Open(path)
	if fileErr != nil {
		return fileErr
	}
	defer f.Close()

	return jpeg.Encode(f, img, &jpeg.Options{option})
}

// PNG将编码生成图片
func ImageEncodePNG(path string, img image.Image) error {
	//func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
	//OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。
	f, fileErr := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	if fileErr != nil {
		return fileErr
	}
	defer f.Close()

	return png.Encode(f, img)
}

// GIF将编码生成图片
func ImageEncodeGIF(path string, img image.Image, opts *gif.Options) error {
	f, fileErr := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if fileErr != nil {
		return fileErr
	}
	defer f.Close()

	return gif.Encode(f, img, opts)
}

func main() {

}
