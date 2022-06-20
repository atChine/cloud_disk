package helper

import (
	"bytes"
	"cloud_disk/core/define"
	"context"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// Md5 值
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GenerateToken 生成token
func GenerateToken(id uint64, identity, name string, second int) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
	}
	//token := jwt.NewWithClaims(jwt.SigningMethodES256, uc) //SigningMethodES256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc) // 换成 SigningMethodHS256
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//AnalyzeToken 解析Token
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claim, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claim.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, nil
}

// MailCodeSend 封装 发送邮箱的方法
func MailCodeSend(mail, code string) error {
	e := email.NewEmail()
	e.From = "AtGao <at911477183@163.com>"
	e.To = []string{mail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", define.EmailAccount, define.EmailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil
}

// RandCode 生成随机数
func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

// UUID  获取UUID
func UUID() string {
	return uuid.NewV4().String()
}

// CosUpload 上传文件到腾讯云
func CosUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SECRETID,
			SecretKey: define.SECRETKEY,
		},
	})

	file, fileHeader, err := r.FormFile("file")
	key := "cloud-disk/" + UUID() + path.Ext(fileHeader.Filename) // 取后缀

	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	if err != nil {
		panic(err)
	}
	return define.CosBucket + "/" + key, nil
}

// CosInitPart 文件分片上传初始化
func CosInitPart(ext string) (string, string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: os.Getenv(define.SECRETID),
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv(define.SECRETKEY),
		},
	})
	key := "cloud-disk/" + UUID() + ext
	// 可选opt,如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		return "", "", err
	}
	return key, v.UploadID, nil
}

// CosPartUpload 上传文件分片
func CosPartUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: os.Getenv(define.SECRETID),
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv(define.SECRETKEY),
		},
	})
	key := r.PostForm.Get("key")
	uploadID := r.PostForm.Get("upload_id")
	file, _, err := r.FormFile("file")
	partNumber, err := strconv.Atoi(r.PostForm.Get("part_number"))
	if err != nil {
		return "", err
	}
	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, file)
	if err != nil {
		return "", err
	}
	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, uploadID, partNumber, bytes.NewReader(buffer.Bytes()), nil,
	)
	if err != nil {
		return "", err
	}
	return strings.Trim(resp.Header.Get("ETag"), "\""), nil
}

// CosCompleteUpload 完成文件分片上传
func CosCompleteUpload(key, uploadID string, cosObject []cos.Object) error {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: os.Getenv(define.SECRETID),
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv(define.SECRETKEY),
		},
	})
	// opt可选
	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cosObject...)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, uploadID, opt,
	)
	if err != nil {
		panic(err)
	}
	return nil
}
