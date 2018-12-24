package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string // 存放zip或jar文件的目录（绝对路径）
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 1. 打开zip文件
	reader, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()

	// 2. 遍历zip或jar中的所有文件，找到要搜索的className文件，读取为[]byte
	for _, file := range reader.File {
		if file.Name == className {
			readCloser, err := file.Open();
			if err != nil {
				return nil, nil, err
			}
			defer readCloser.Close()

			data, err := ioutil.ReadAll(readCloser)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}

func newZipEnrty(path string) *ZipEntry {
	// 将path转换成绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absDir}
}
