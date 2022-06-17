package test

import (
	"bytes"
	"cloud_disk/core/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)


func TestXormTest(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:adminroot@tcp(127.0.0.1:3306)/Tables?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil{
		t.Fatal(err)
	}

	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	b := new(bytes.Buffer)
	err = json.Indent(b, marshal, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b.String())
}