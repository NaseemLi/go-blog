package captchaapi

import (
	"goblog/common/res"
	"goblog/global"
	"image/color"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type CaptchaApi struct {
}

type CaptchaResponse struct {
	CaptchaID string `json:"captchaID"`
	Captcha   string `json:"captcha"`
}

func (CaptchaApi) CaptchaView(c *gin.Context) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		// 验证码图片的高度，以像素为单位
		Height: 60,
		// 验证码图片的宽度，以像素为单位
		Width: 200,
		// 验证码图片中随机噪点的数量，0 表示没有噪点
		NoiseCount: 1,
		// 控制显示在验证码图片中的线条：2 | 4 表示曲线 + 点线
		// 可选值：1=直线，2=曲线，4=点线，8=虚线，16=中空直线，32=中空曲线
		ShowLineOptions: 2 | 4,
		// 验证码的长度（字符数）
		Length: 4,
		// 验证码的字符源，这里使用数字 + 小写字母
		Source: "1234567890",
		// 验证码背景颜色（RGBA）
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		// 使用的字体文件
		//Fonts: []string{"wqy-microhei.ttc"},
	}
	// 将配置赋值给 driverString
	driverString = captchaConfig
	// 将字体文件转换为可用格式，赋值给 driver
	driver = driverString.ConvertFonts()
	// 使用 driver 和 stores 创建新的验证码实例
	captcha := base64Captcha.NewCaptcha(driver, global.CaptchaStore)
	// 生成验证码，返回 lid（唯一ID）、lb64s（base64图像）、lerr（错误信息）
	lid, lb64s, _, err := captcha.Generate()
	if err != nil {
		res.FailWithMsg("土木验证码生成失败", c)
		logrus.Error(err)
	}
	// 从验证码存储器中获取验证码值
	res.OkWithData(CaptchaResponse{
		CaptchaID: lid,
		Captcha:   lb64s,
	}, c)
}
