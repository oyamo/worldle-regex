package parser

import (
	"bytes"
	"html/template"
	"log"
)

func ParseTemplate(name string, data interface{}) []byte {
	tmp, err := template.ParseFiles(name)
	if err != nil {
		log.Println(err)
		return []byte("Error: 500")
	}
	var buf bytes.Buffer
	err = tmp.Execute(&buf, data)
	if err != nil {
		log.Println(err)
		return []byte("Error: 500")
	}
	return buf.Bytes()
}
