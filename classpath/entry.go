package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 寻找和读取 class 文件
	// 入参：className - class文件的相对路径，eg. 如果要读取 java.lang.Object 类，则className = java/lang/Object.class
	// 返回值：
	// 1. 读取到的class文件内容的[]byte
	// 2. 最终定位到包含className文件的Entry对象
	// 3. 错误信息error
	readClass(className string) ([]byte, Entry, error)
	String() string // toString()
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path) // 从path1/classes和path2路径下查找并读取className文件：java -cp path1/classes:path2
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path) // 从path的所有jar文件中查找并读取className文件：java -cp path/*
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".zip") {
		return newZipEnrty(path) // 从path/lib1.jar下查找并读取className文件：java -cp path/lib1.jar 或者 java -cp path/lib1.zip
	}
	return newDirEntry(path) // 从path下查找并读取className文件：java -cp path
}
