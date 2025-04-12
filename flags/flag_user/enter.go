package flaguser

import (
	"fmt"
	"goblog/global"
	"goblog/models"
	"goblog/models/enum"
	"goblog/utils/pwd"
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/term"
)

type FlagUser struct {
}

func (FlagUser) Create() {
	var role enum.RoleType
	fmt.Println("选择角色:		1.超级管理员		2.普通用户		3.访客")
	_, err := fmt.Scan(&role)
	if err != nil {
		logrus.Errorf("输入错误 %s", err)
		return
	}
	if !(role == 1 || role == 2 || role == 3) {
		logrus.Errorf("输入角色错误 %s", err)
		return
	}

	var username string
	fmt.Println("请输入用户名")
	fmt.Scan(&username)
	//查用户名是否存在
	var model models.UserModel
	err = global.DB.Take(&model, "username + ?", username).Error
	if err == nil {
		logrus.Errorf("此用户名已存在")
		return
	}
	fmt.Println("请输入密码:")
	password, err := term.ReadPassword(int(os.Stdin.Fd())) // 隐藏输入
	if err != nil {
		fmt.Println("读取密码时出错:", err)
		return
	}
	fmt.Println("请再次输入密码:")
	repassword, err := term.ReadPassword(int(os.Stdin.Fd())) // 隐藏输入
	if err != nil {
		fmt.Println("读取密码时出错:", err)
		return
	}
	if string(password) != string(repassword) {
		fmt.Println("两次密码不一致")
	}

	hashPwd, _ := pwd.GenerateFromPassword(string(password))
	err = global.DB.Create(&models.UserModel{
		Username:       username,
		Nickname:       "admin",
		RegisterSource: enum.RegisterTerminalSourceType,
		Password:       hashPwd,
		Role:           role,
	}).Error
	if err != nil {
		fmt.Println("创建用户失败", err)
	}
	logrus.Infof("用户创建成功")
}
