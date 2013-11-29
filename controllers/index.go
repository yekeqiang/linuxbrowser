package controllers
 
import (
    "github.com/fhbzyc/linuxbrowser/libs/dir"
    "github.com/fhbzyc/linuxbrowser/libs/linux"
    //  "fmt"
    "github.com/astaxie/beego"
    //  "os"
    //  "path/filepath"
    "net/url"
    "strconv"
    "strings"
 
//  "syscall"
//  "time"
)
 
type IndexController struct {
    beego.Controller
}
 
type fileType struct {
    dir.SysFile
    Fsize  string
    Ftime  string
    Fgname string
    Funame string
}
 
func (this *IndexController) Get() {
 
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
 
}
 
func lastDir(path string) string {
 
    runeURL := []rune(path)
    length := len(runeURL)
    if length > 0 {
        if string(runeURL[length-1]) == "/" {
 
            path = string(runeURL[0 : length-1])
        }
    }
 
    array := strings.Split(path, "/")
    array[len(array)-1] = ""
 
    return strings.Join(array, "/")
}
 
func browser(urlQuery string) (RequestURI string, dirs []fileType, files []fileType, err error) {
 
    DS := "/"
    //--------------------路径导航--------------------//
 
    filesArray, err1 := dir.ReadDir(urlQuery)
 
    if err1 != nil {
        newUrl, err2 := url.QueryUnescape(urlQuery)
        if err2 != nil {
            err = err2
            return
        } else {
            filesArray, err2 = dir.ReadDir(newUrl)
            if err2 != nil {
                err = err2
                return
            } else {
                RequestURI = newUrl
            }
        }
    } else {
        RequestURI = urlQuery
    }
 
    //结尾加上 '/'
    //      runeURL := []rune(RequestURI)
    //      length := len(runeURL)
    if !strings.HasSuffix(RequestURI, DS) {
        RequestURI += DS
    }
 
    for _, val := range filesArray {
        time := val.Ftime.Format("2006-01-02 03:04:05")
        temp := val.Fsys
        fgname := linux.GetGname(int(temp.Gid))
        funame := linux.GetUname(int(temp.Uid))
 
        if val.Ftype == 0 {
            dirs = append(dirs, fileType{SysFile: val, Fsize: "", Ftime: time, Fgname: fgname, Funame: funame})
        } else {
            files = append(files, fileType{SysFile: val, Fsize: strconv.FormatFloat(float64(val.Fsize)/(1024*1024), 'f', 8, 64), Ftime: time, Fgname: fgname, Funame: funame})
        }
    }
 
    return
}
