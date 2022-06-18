package test

import (
	"bytes"
	"cloud_disk/core/define"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestFileUploadByFilepath(t *testing.T) {
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

	key := "cloud-disk/laiMeiYun2.jpg"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./img/laimeiyun.jpeg", nil,
	)
	if err != nil {
		panic(err)
	}
}

// TestFileUploadByReader 测试文件上传初始化
func TestFileUploadByReader(t *testing.T) {
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
	key := "cloud-disk/laiMeiYun3.jpg"
	// 可选opt,如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		t.Fatal(err)
	}
	UploadID := v.UploadID //1655554550820baffb73097907b15e9e435cea9efd663aa66af7e311e94ce46a1bc99def8e
	fmt.Println(UploadID)
}

//分片上传
func TestPartUpload(t *testing.T) {
	UploadID := "1655554550820baffb73097907b15e9e435cea9efd663aa66af7e311e94ce46a1bc99def8e"
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
	key := "cloud-disk/laiMeiYun3.jpg"
	file, err2 := os.ReadFile("3.chunk")
	if err2 != nil {
		t.Fatal(err2)
	}
	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 1, bytes.NewReader(file), nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	PartETag := resp.Header.Get("ETag") //0fe522b2a32d98024c8d56b205901d5d
	fmt.Println(PartETag)
}

// 完成分片上传
func TestCompleteUpload(t *testing.T) {
	UploadID := "1655554550820baffb73097907b15e9e435cea9efd663aa66af7e311e94ce46a1bc99def8e"
	PartETag := "0fe522b2a32d98024c8d56b205901d5d"
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
	key := "cloud-disk/laiMeiYun3.jpg"

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 1, ETag: PartETag},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		panic(err)
	}
}
