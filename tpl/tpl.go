package tpl

import (
	"bytes"
	"html/template"
)

// ParseTemplate 解析模板字符串并返回结果字符串。
// name 为模板名称，templateStr 为模板字符串，data 为模板数据。
func ParseTemplate(name, templateStr string, data interface{}) (string, error) {
	t, err := template.New(name).Parse(templateStr)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// ParseTemplateFile 解析模板文件并返回结果字符串。
func ParseTemplateFile(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
