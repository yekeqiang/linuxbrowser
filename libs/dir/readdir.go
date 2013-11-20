package dir

import (
	//	"fmt"
	"os"
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

	for _, f := range files {
		if f.IsDir() {
			array = append(array, SysFile{Ftype: 0, Fname: f.Name(), Fsize: f.Size(), Fmode: f.Mode(), Ftime: f.ModTime(), Fsys: f.Sys().(*syscall.Stat_t)})
		} else {
			array = append(array, SysFile{Ftype: 1, Fname: f.Name(), Fsize: f.Size(), Fmode: f.Mode(), Ftime: f.ModTime(), Fsys: f.Sys().(*syscall.Stat_t)})
		}
	}

	return array, nil
}
