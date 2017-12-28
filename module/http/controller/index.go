package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/golang/freetype"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"
)

type IndexController struct {
	beego.Controller
}

//首页
func (self *IndexController) Index() {
	self.TplName = "index.html"
}

//临时生成图片
func (self *IndexController) GenImg() {
	txt_left, err := self.GetInt("txt_left", 0)
	if err != nil {
		fmt.Println("参数 txt_left 格式不对")
		return
	}
	txt_top, err := self.GetInt("txt_top", 0)
	if err != nil {
		fmt.Println("参数 txt_top 格式不对")
		return
	}
	txt_size, err := self.GetInt("txt_size", 0)
	if err != nil {
		fmt.Println("参数 txt_size 格式不对")
		return
	}
	txt_val := self.GetString("txt_val")
	img_src := self.GetString("img_src")
	if len(img_src) > 0 {
		img_src = fmt.Sprintf("./view%s", img_src)
	} else {
		img_src = "./view/res/img/doutu.jpg"
	}

	//生成图片
	//读取原始图片
	img_file, err := os.Open(img_src)
	if err != nil {
		fmt.Println("打开图片出错 ", err)
		return
	}
	defer img_file.Close()
	img, err := jpeg.Decode(img_file)
	if err != nil {
		fmt.Println("解码图片时出错 ", err)
		return
	}

	//读取字体
	fontBytes, err := ioutil.ReadFile("./view/res/font/mnkt.ttf")
	if err != nil {
		fmt.Println("读取字体出错 ", err)
		return
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println("转换字体样式出错 ", err)
		return
	}

	des_img := image.NewRGBA(img.Bounds())
	draw.Draw(des_img, img.Bounds(), img, image.ZP, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(float64(txt_size))
	c.SetClip(des_img.Bounds())
	c.SetDst(des_img)
	c.SetSrc(image.Black)

	pt := freetype.Pt(txt_left, txt_top+txt_size-int((float64(txt_size)*float64(5)/float64(48))))
	// pt := freetype.Pt(txt_left, txt_top)
	_, err = c.DrawString(txt_val, pt)
	if err != nil {
		fmt.Println("向图片写文字出错 ", err)
		return
	}

	new_img, err := os.Create("new.jpg")
	if err != nil {
		fmt.Println("建立新图片文件出错 ", err)
		return
	}
	defer new_img.Close()
	err = jpeg.Encode(new_img, des_img, &jpeg.Options{100})
	if err != nil {
		fmt.Println("生成新图片出错 ", err)
		return
	}
	fmt.Println("添加水印图片结束请查看")
}
