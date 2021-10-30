package utils

import (
	"encoding/xml"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func XmlParse(filePath string, value interface{}) (error) {

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0755)
	if err != nil {
		LogError("%s:%s,%s", filePath, "待解析的文件打开失败", err)
		return errors.New("待解析的文件打开失败")
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		LogError("%s:%s,%s", filePath, "待解析的文件读取失败", err)
		return errors.New("待解析的文件读取失败")
	}
	err = xml.Unmarshal(data, value)
	if err != nil {
		LogError("%s:%s,%s", filePath, "待解析的文件格式转换失败", err)
		return errors.New("待解析的文件格式转换失败")
	}
	return nil
}
