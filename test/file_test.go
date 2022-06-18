package test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

//定义文件分片大小
const chunkSize = 1024 * 1024 * 1 // 1MB

//文件分片
func TestFileFarg(t *testing.T) {
	fileInfo, err := os.Stat("yyyyy.mp4")
	if err != nil {
		t.Error(err)
	}
	//计算文件分片数量
	var chunkNum = int(fileInfo.Size()/chunkSize) + 1
	myFile, err := os.OpenFile("yyyyy.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Error(err)
	}
	//创建数组存放文件分片
	bytes := make([]byte, chunkSize)
	//循环读取文件分片
	for i := 0; i < chunkNum; i++ {
		//指定读取文件分片的起始位置
		myFile.Seek(int64(i*chunkSize), 0)
		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			bytes = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		//将文件分片读取到bytes中
		myFile.Read(bytes)
		//将文件分片写入到文件分片文件中
		f, err := os.OpenFile("./"+strconv.Itoa(i)+"."+"chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		//f, err := os.OpenFile("./"+strconv.Itoa(i)+path.Ext("yyyyy.mp4"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Error(err)
		}
		f.Write(bytes)
		f.Close()
	}
	myFile.Close()
}

//分片文件合并
func TestFileMerge(t *testing.T) {
	//创建一个文件分片汇总文件
	myFile, err := os.OpenFile("yyyyy2.mp4", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	//计算文件分片数量
	fileInfo, err := os.Stat("yyyyy.mp4")
	if err != nil {
		t.Error(err)
	}
	var chunkNum = int(fileInfo.Size()/chunkSize) + 1
	//循环读取文件分片
	for i := 0; i < chunkNum; i++ {
		//读取文件分片
		f, err := os.OpenFile("./"+strconv.Itoa(i)+"."+"chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Error(err)
		}
		all, err := ioutil.ReadAll(f)
		if err != nil {
			t.Error(err)
		}
		//将文件分片写入到文件分片汇总文件中
		myFile.Write(all)
		f.Close()
	}
	myFile.Close()
}

//校验文件一致性
func TestFileCheck(t *testing.T) {
	//获取第一个文件的信息
	file1, _ := os.OpenFile("yyyyy.mp4", os.O_RDONLY, 0666)
	byteFile1, _ := ioutil.ReadAll(file1)
	//获取第二个文件的信息
	file2, _ := os.OpenFile("yyyyy2.mp4", os.O_RDONLY, 0666)
	byteFile2, _ := ioutil.ReadAll(file2)
	//校验文件一致性
	if len(byteFile1) != len(byteFile2) {
		t.Error("文件不一致")
		return
	}
	if fmt.Sprintf("%x", md5.Sum(byteFile1)) != fmt.Sprintf("%x", md5.Sum(byteFile2)) {
		t.Error("文件不一致")
	} else {
		t.Log("文件一致")
	}

}
