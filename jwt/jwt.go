/*
 * @Author: GG
 * @Date: 2022-09-16 10:56:24
 * @LastEditTime: 2022-09-16 15:04:12
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\jwt\jwt.go
 *
 */
package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 自定义信息，token包含的信息结构体
type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var MySecret []byte
var refresTokenType = "refres"

const TOKEN_EXPIRES_AT = time.Minute * 2
const REFRES_TOKEN_EXPIRES_AT = time.Hour * 2

func init() {
	// jwtKey = []byte(os.Getenv("JWT_SECRET"))
	MySecret = []byte("secret")
}

// 获取token
func GenToken(username string, password string, tokenType string) (string, error) {
	// 创建自定义信息
	var expiresAt int64
	if tokenType == refresTokenType {
		expiresAt = time.Now().Add(REFRES_TOKEN_EXPIRES_AT).Unix()
	} else {
		expiresAt = time.Now().Add(TOKEN_EXPIRES_AT).Unix()
	}
	c := MyClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,  // 有效期2小时，返回时间戳
			Issuer:    "jwt-demo", //签发人
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	// 注意这个地方一定要是字节切片不能是字符串

	tokenStr, err := token.SignedString(MySecret)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}

	return tokenStr, nil
}

// 解析token
func ParseToken(tokenStr string) (*jwt.Token, *MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, nil, err
	}
	// valid 是否有效
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return token, claims, nil
	}
	return nil, nil, errors.New("invalid token")
}

// 解析token  refresToken
func ParseToken1(tokenStr string, refresTokenStr string) (string, *MyClaims, error) {
	token, _ := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	// if err != nil {
	// 	return "", nil, err
	// }

	// valid 是否有效,返回token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		fmt.Printf("refresTokenStr: %v\n", refresTokenStr)
		fmt.Println("旧的token")
		return token.Raw, claims, nil
	}

	// token失效，检查refreToken，有效则返回新token
	refresToken, err := jwt.ParseWithClaims(refresTokenStr, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	if err != nil {
		return "", nil, err
	}

	// refresToken 是否有效
	if claims, ok := refresToken.Claims.(*MyClaims); ok && refresToken.Valid {
		tokenStr, _ := GenToken(claims.Username, claims.Password, "")
		fmt.Printf("refresTokenStr: %v\n", refresTokenStr)
		fmt.Println("新的token")
		return tokenStr, claims, nil
	}
	return "", nil, errors.New("invalid token")
}
func main() {
	// s, err := GenToken("ghz", "123", "")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("s: %v\n", s)
	// r, err := GenToken("ghz", "123", refresTokenType)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("r: %v\n", r)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImdoeiIsInBhc3N3b3JkIjoiMTIzIiwiZXhwIjoxNjYzMzExOTQzLCJpc3MiOiJqd3QtZGVtbyJ9.zts8X1KkvpJvhWOODapYlLq7zncvPU_iagNAyscuxk0"
	rToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImdoeiIsInBhc3N3b3JkIjoiMTIzIiwiZXhwIjoxNjYzMzE4NzY2LCJpc3MiOiJqd3QtZGVtbyJ9.Aeot9k9JXrn2BzFej5rXcQlFND2hFtTyaG_T-u5f2n0"
	newToken, mc, err := ParseToken1(token, rToken)
	if err != nil {
		panic(err)
	}
	fmt.Printf("newToken: %v\n", newToken)
	fmt.Printf("mc.Password: %v\n", mc.Password)
	fmt.Printf("mc.Username: %v\n", mc.Username)
}
