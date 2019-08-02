package qrcode

import (
	"io/ioutil"
	"os"
	"testing"
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
		panic(err)
	}
}

func TestQrCode_GenerateQRCode(t *testing.T) {
	pngData,err:=qrCode.GenerateQRCode("https://demo.u1200.com/af83b54a95b95ade250e334704ffddad7cb06fc9")
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
	pngData,err:=qrCode.GenerateQRCodeWithLogo("https://demo.u1200.com/af83b54a95b95ade250e334704ffddad7cb06fc9")
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

func TestQrCode_GenerateQrCodeWithAlphaLogoArea(t *testing.T) {
	pngData,err:=qrCode.GenerateQrCodeWithAlphaLogoArea("test")
	if err!=nil{
		t.Error(err)
	}

	file,err:=os.Create("qrcodeWithAlpha.png")
	if err!=nil{
		t.Error(err)
	}
	file.Write(pngData)
}
