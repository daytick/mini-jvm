package classpath

import (
	"file/filepath"
	"io/ioutil"
	"errors"
	"archive/zip"
)

type ZipEntry struct {
	absPath string // jar 或 zip 文件的绝对路径
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, ZipEntry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.close() // 确保打开的文件得以关闭
	for i,f := range r.File { // 遍历压缩包内的文件
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer f.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className) // 提示压缩包内不存在名为 className 的文件
}

func (self *ZipEntry) String() string {
	return self.absPath
}
