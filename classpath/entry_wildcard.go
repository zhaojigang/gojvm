package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// WildcardEntry 其实属于 CompositeEntry 的一种，也是从组合路径进行寻找
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // 去掉结尾的*号
	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 虚拟机规范：通配符加载只可以加载baseDir目录下的jar，其子目录下的jar不可以加载
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		// 获取baseDir下的jar文件
		if strings.HasSuffix(path, ".jar") {
			jarEnrty := newZipEnrty(path)
			compositeEntry = append(compositeEntry, jarEnrty)
		}
		return nil
	}
	// 遍历baseDir下的每一个文件，针对每一个文件执行walkFn函数操作
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
