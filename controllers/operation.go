package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	//	"path/filepath"
	"strings"
	//	"math"
//	"strconv"
)

type OperationController struct {
	beego.Controller
}

func (this *OperationController) Get() {

	action := this.GetString("action")
	action = strings.Trim(action, " ")

	if action == "" {
		this.jsonEncode("26", "")
		return
	}

	if action == "delete" {
		this.deleteFile()
		return
	}

	if action == "edit" {
		this.editFile()
		return
	}
}

func (this *OperationController) deleteFile() {

	removeFile := this.GetString("file")

	fmt.Println(removeFile)

	if removeFile != "" {
		err := os.Remove(removeFile)
		if err != nil {
			this.jsonEncode("33", "")
		} else {
			this.jsonEncode("0", "")
		}

	} else {
		this.jsonEncode("39", "")
	}
}

func (this *OperationController) editFile() {

	editFile := this.GetString("file")

	if editFile != "" {

		f, err := os.OpenFile(editFile, os.O_RDWR, 0666)

		if err != nil {
			this.jsonEncode("33", "")
		} else {

			info, _ := f.Stat()
			buf := make([]byte, info.Size())
			f.Read(buf)
			this.Data["value"] = string(buf)
			this.TplNames = "operation/edit.html"

		}
		return

	} else {
		this.jsonEncode("39", "")
	}
	return
}

func (this *OperationController) jsonEncode(code string, message string) {

	json := []string{code, message}

	this.Data["json"] = &json
	this.ServeJson()
}
