package qrcode

import (
	"testing"
	"io/ioutil"
	"os"
)

var qrCode *QrCode

func init(){
	logoData,err:=ioutil.ReadFile("logo.png")
	fontData,err:=ioutil.ReadFile("zh.ttf")
	if err!=nil{
		return
	}

	qrCode,err=Init(fontData,logoData,0,0,0,"","")
	if err!=nil{
	}
}

func TestQrCode_GenerateQRCode(t *testing.T) {

	pngData,err:=qrCode.GenerateQRCode("test")
	if err!=nil{
		t.Error(err)
	}

	file,err:=os.Create("qrcode.png")
	if err!=nil{
		t.Error(err)
	}
	file.Write(pngData)
}

func TestQrCode_GenerateQRCodeWithLogo(t *testing.T) {
	pngData,err:=qrCode.GenerateQRCodeWithLogo("test")
	if err!=nil{
		t.Error(err)
	}

	file,err:=os.Create("qrcodeWithLogo.png")
	if err!=nil{
		t.Error(err)
	}
	file.Write(pngData)
}


func TestQrCode_GenerateQRCodeWithLogoAndTitle(t *testing.T) {
	pngData,err:=qrCode.GenerateQRCodeWithLogoAndTitle("test","都看到看看地方")
	if err!=nil{
		t.Error(err)
	}

	file,err:=os.Create("qrcodeWithLogoAndTitle.png")
	if err!=nil{
		t.Error(err)
	}
	file.Write(pngData)
}