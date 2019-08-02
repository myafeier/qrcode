package qrcode

import (
	"bytes"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"unicode/utf8"
)

type QrCode struct {
	Font *truetype.Font  //字体文件
	LogoImage image.Image  //logo文件
	Size  int       // 生成的图片宽度
	FontSize float64    //字号
	FontDPI  float64  //字体密度
	FrontColor color.Color  //字体颜色
	BackColor  color.Color   //背景颜色
}

func Init(fontByte,logoByte []byte, size int,fontSize float64,fontDPI float64 ,frontColor string,backColor string)(qrCode *QrCode,err error){
	qrCode=new(QrCode)
	qrCode.Size=size
	qrCode.FontSize=fontSize
	qrCode.FontDPI=fontDPI

	if frontColor!=""{
		qrCode.FrontColor, err = colorful.Hex(frontColor)
		if err!=nil{
			return
		}
	}else{
		qrCode.FrontColor=color.Black
	}

	if frontColor!=""{
		qrCode.BackColor, err = colorful.Hex(backColor)
		if err!=nil{
			return
		}
	}else{
		qrCode.BackColor=color.White
	}
	if size!=0{
		qrCode.Size=size
	}else{
		qrCode.Size=1000
	}

	if fontSize>1{
		qrCode.FontSize=fontSize
	}else{
		qrCode.FontSize=40
	}

	if fontDPI>1{
		qrCode.FontDPI=fontDPI
	}else{
		qrCode.FontDPI=72
	}


	if fontByte!=nil{
		err=qrCode.readFontFromByte(fontByte)
		if err!=nil{
			return
		}
	}

	if logoByte!=nil{
		err=qrCode.readLogoImageFromByte(logoByte)
		if err!=nil{
			return
		}
	}

	return
}


func (self *QrCode)GenerateQrCodeWithAlphaLogoArea(url string)(imgData []byte,err error){

	qrCodeData,err:=self.GenerateQRCode(url)
	if err!=nil{
		return
	}

	qrCodeImage,_,err:=image.Decode(bytes.NewReader(qrCodeData))
	if err!=nil{
		return
	}


	offsetX := (self.Size-200)/2

	newImg := image.NewRGBA(qrCodeImage.Bounds())

	//draw.Draw(newImg,newImg.Bounds(),&image.Uniform{self.BackColor}, image.Point{}, draw.Src)
	draw.Draw(newImg,qrCodeImage.Bounds(),qrCodeImage,image.Point{-1,-1},draw.Src)

	blank := image.NewRGBA(image.Rect(0,0,200,200))

	draw.Draw(newImg, newImg.Bounds(), blank, image.Point{-offsetX,-offsetX}, draw.Src)


	buffer:=new(bytes.Buffer)
	err=png.Encode(buffer,newImg)
	if err!=nil{
		return
	}
	imgData=buffer.Bytes()
	return




}

func (self *QrCode)GenerateQRCodeWithLogoAndTitle(url,title string)(imgData []byte,err error)  {

	qrCodeData,err:=self.GenerateQRCodeWithLogo(url)
	if err!=nil{
		return
	}

	qrCodeImage,_,err:=image.Decode(bytes.NewReader(qrCodeData))
	if err!=nil{
		return
	}



	c:=freetype.NewContext()
	c.SetDPI(self.FontDPI)
	c.SetFont(self.Font)
	c.SetFontSize(self.FontSize)
	titleImage:=image.NewRGBA(image.Rect(0,0,self.Size,c.PointToFixed(self.FontSize).Round()*2))
	c.SetClip(titleImage.Bounds())
	c.SetDst(titleImage)
	c.SetSrc(image.Black)
	startx:=1000-c.PointToFixed(40).Round()*utf8.RuneCountInString(title)-100
	starty:=100-c.PointToFixed(40).Round()+10
	position:=freetype.Pt(startx,starty) //字体出现的位置
	_,err=c.DrawString(title,position)
	if err!=nil{
		return
	}

	newImg := image.NewRGBA(qrCodeImage.Bounds())

	draw.Draw(newImg,newImg.Bounds(),&image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg,newImg.Bounds(),qrCodeImage,image.Point{-1,-1}, draw.Over)
	draw.Draw(newImg,newImg.Bounds(),titleImage,image.Point{-1,-1},draw.Over)
	buffer:=new(bytes.Buffer)
	err=png.Encode(buffer,newImg)
	if err!=nil{
		return
	}
	imgData=buffer.Bytes()
	return


}

func (self *QrCode)GenerateQRCodeWithLogo(url string)(imageData []byte,err error){

	qrCodeData,err:=self.GenerateQRCode(url)
	if err!=nil{
		return
	}

	qrCodeImage,_,err:=image.Decode(bytes.NewReader(qrCodeData))
	if err!=nil{
		return
	}


	offsetX := (self.Size-self.LogoImage.Bounds().Dx())/2

	newImg := image.NewRGBA(qrCodeImage.Bounds())

	draw.Draw(newImg,newImg.Bounds(),&image.Uniform{self.BackColor}, image.Point{}, draw.Src)
	draw.Draw(newImg,qrCodeImage.Bounds(),qrCodeImage,image.Point{-1,-1},draw.Over)
	draw.Draw(newImg, newImg.Bounds(), self.LogoImage, image.Point{-offsetX,-offsetX}, draw.Over)


	buffer:=new(bytes.Buffer)
	err=png.Encode(buffer,newImg)
	if err!=nil{
		return
	}
	imageData=buffer.Bytes()
	return



}

func (self *QrCode)GenerateQRCode(url string)(imageData []byte,err error)  {

	q, e := qrcode.New(url, qrcode.High)
	if e != nil {
		return
	}

	q.ForegroundColor = self.FrontColor
	q.BackgroundColor = self.BackColor

	imageData,err= q.PNG(self.Size)
	return
}



func (self *QrCode)readFontFromByte(fontByte []byte)(err error){
	self.Font,err=freetype.ParseFont(fontByte)
	return
}

func (self *QrCode)readLogoImageFromByte(logoByte []byte)(err error){

	logoImage,_,err:=image.Decode(bytes.NewReader(logoByte))
	if err!=nil{
		return
	}
	self.LogoImage=resize.Resize(300,300,logoImage,resize.Lanczos3)
	return
}


