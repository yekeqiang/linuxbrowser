package controllers

import (
	//	"fmt"
	"github.com/astaxie/beego"
	"os"
	//	"path/filepath"
	"strings"
	//	"math"
	//	"runtime
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
	case "remove":
		this.remove()
	}

	return

}

func (this *OperationController) read() {

	file := this.GetString("file")

	if file != "" {

		f, err := os.OpenFile(file, os.O_RDWR, 0444)

		if err != nil {
			this.jsonEncode(68, err.Error())
		} else {

			info, _ := f.Stat()
			buf := make([]byte, info.Size())
			f.Read(buf)
			this.Data["value"] = string(buf)
			this.TplNames = "operation/edit.html"

		}
		return

	} else {
		this.jsonEncode(81, "")
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

	oldname := this.GetString("oldname")
	newname := this.GetString("newname")

	err := os.Rename(oldname, newname)

	if err != nil {
		this.jsonEncode(114, err.Error())
	}

	this.jsonEncode(0, "")

	return
}

func (this *OperationController) mkdir() {

	dirname := this.GetString("dirname")

	if dirname != "" {
		err := os.Mkdir(dirname, 0664)
		if err == nil {
			this.jsonEncode(0, "")
		} else {
			this.jsonEncode(89, err.Error())
		}
		return
	}

	this.jsonEncode(89, "create dir fail ! ")

	return
}

func (this *OperationController) create() {

	filename := this.GetString("filename")

	f, err := os.Open(filename)
	if err == nil {
		f.Close()
		this.jsonEncode(33, filename+" file exists")
	} else if os.IsNotExist(err) {

		_, err2 := os.Create(filename)
		if err2 == nil {
			this.jsonEncode(0, "")
		} else {
			this.jsonEncode(89, err2.Error())
		}
		return
	}

	this.jsonEncode(89, "create file fail ! ")

	return
}

func (this *OperationController) remove() {

	path := this.GetString("filename")

	if path != "" {
		f, err := os.Open(path)
		if err != nil {
			this.jsonEncode(89, err.Error())
			return
		}

		_, err2 := f.Readdir(0)
		var err3 error
		if err2 == nil {
			err3 = os.RemoveAll(path)
		} else {
			err3 = os.Remove(path)
		}

		if err != nil {
			this.jsonEncode(185, err3.Error())
		} else {
			this.jsonEncode(0, "")
		}
		return
	}
}

func (this *OperationController) jsonEncode(code int, message string) {

	json := []string{strconv.Itoa(code), message}

	this.Data["json"] = &json
	this.ServeJson()
}
