/*
 * @Author: GG
 * @Date: 2022-08-18 17:01:09
 * @LastEditTime: 2022-08-18 18:10:52
 * @LastEditors: GG
 * @Description:login service
 * @FilePath: \golang-demo\blog\golang\service\login.go
 *
 */
package service

import (
	"errors"
	"golang-demo/blog/golang/dao"
	"golang-demo/blog/golang/models"
	"golang-demo/blog/golang/utils"
	"log"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	// 加密
	passwd = utils.Md5Crypt(passwd, "mszlu")
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}

	uid := user.Uid
	//生成tokan, jwt生成令牌 A.B.C
	token, err := utils.Award(&uid)
	if err != nil {
		log.Panicln("jwt token 生成失败:", err)
		return nil, errors.New("jwt token 生成失败")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar

	var result = &models.LoginRes{
		token,
		userInfo,
	}
	return result, nil
}
