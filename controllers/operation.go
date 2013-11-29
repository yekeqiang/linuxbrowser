package controllers

import (
	//	"fmt"
	"github.com/astaxie/beego"
	"os"
	//	"path/filepath"
	"strings"
	//	"math"
	//	"runtime"
	"strconv"
)

type OperationController struct {
	beego.Controller
}

func (this *OperationController) Get() {

	action := this.GetString("action")
	action = strings.Trim(action, " ")

	switch action {

	case "":

		RequestURI, dirs, files, err := browser(this.Ctx.Request.RequestURI)
		if err != nil {
			this.Ctx.WriteString("404")
			return
		}

		this.Data["RequestURI"] = RequestURI
		this.Data["parentPath"] = lastDir(RequestURI)
		this.Data["dirs"] = dirs
		this.Data["files"] = files
		this.TplNames = "index/index.html"

	case "read":
		this.read()
	case "copy":
		this.copy()
	case "move":
		this.move()
	case "mkdir":
		this.mkdir()
	case "rename":
		this.rename()
	case "create":
		this.create()
	case "delete":
		this.delete()
	}

	return

}

func (this *OperationController) read() {

	file := this.GetString("file")

	if file != "" {

		f, err := os.OpenFile(file, os.O_RDWR, 0666)

		if err != nil {
			this.jsonEncode(33, "")
		} else {

			info, _ := f.Stat()
			buf := make([]byte, info.Size())
			f.Read(buf)
			this.Data["value"] = string(buf)
			this.TplNames = "operation/edit.html"

		}
		return

	} else {
		this.jsonEncode(39, "")
	}
	return
}

func (this *OperationController) copy() {

	return
}

func (this *OperationController) move() {

	file := this.GetString("file")
	newname := this.GetString("newname")

	err := os.Rename(file, newname)

	if err != nil {
		this.jsonEncode(89, "")
	}

	this.jsonEncode(0, "")

	return
}
func (this *OperationController) rename() {

	file := this.GetString("file")
	newname := this.GetString("newname")

	err := os.Rename(file, newname)

	if err != nil {
		this.jsonEncode(89, "")
	}

	this.jsonEncode(0, "")

	return
}

func (this *OperationController) mkdir() {

	dir := this.GetString("dir")

	f, err := os.Open(dir)

	defer f.Close()

	if err != nil && os.IsNotExist(err) {

		if dir != "" {
			err2 := os.Mkdir(dir, 0664)
			if err2 == nil {
				this.jsonEncode(0, "")
			}
		}
	}

	this.jsonEncode(89, "")

	return
}

func (this *OperationController) create() {

	createFile := this.GetString("file")

	f, err := os.Open(createFile)

	defer f.Close()

	if err != nil && os.IsNotExist(err) {

		if createFile != "" {
			_, err2 := os.Create(createFile)
			if err2 == nil {
				this.jsonEncode(0, "")
			}
		}
	}

	this.jsonEncode(89, "")

	return
}

func (this *OperationController) delete() {

	removeFile := this.GetString("file")

	if removeFile != "" {
		err := os.Remove(removeFile)
		if err != nil {
			this.jsonEncode(62, "")
		} else {
			this.jsonEncode(0, "")
		}

	} else {
		this.jsonEncode(39, "")
	}
}

func (this *OperationController) jsonEncode(code int, message string) {

	json := []string{strconv.Itoa(code), message}

	this.Data["json"] = &json
	this.ServeJson()
}
