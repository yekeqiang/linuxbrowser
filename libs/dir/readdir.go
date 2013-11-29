package dir

import (
	//	"fmt"
	"os"
	"strings"
	"syscall"
	"time"
)

type SysFile struct {
	Ftype int
	Fname string
	Fsize int64
	Fmode os.FileMode
	Ftime time.Time
	Fsys  *syscall.Stat_t
}

func ReadDir(path string) (ret []SysFile, err error) {
	array := []SysFile{}
	dir, err := os.Open(path)
	if err != nil {
		return array, err
	}

	files, _ := dir.Readdir(0)
	var ftype int
	for _, f := range files {

		if f.IsDir() {
			ftype = 0
		} else {
			ftype = 1
		}

		if strings.HasPrefix(f.Mode().String(), "L") || strings.HasPrefix(f.Mode().String(), "l") {
			//链接类型文件
			ftype = 2
		}

		array = append(array, SysFile{Ftype: ftype, Fname: f.Name(), Fsize: f.Size(), Fmode: f.Mode(), Ftime: f.ModTime(), Fsys: f.Sys().(*syscall.Stat_t)})
	}

	return array, nil
}
