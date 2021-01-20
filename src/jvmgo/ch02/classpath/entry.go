package classpath

import (
	"os"
	"strings" // 操作字符的工具集
)

const pathListSeparator = string(os.PathListSeparator) // 路径分隔符

// 表示类路径项的接口
// 在 Go 中， 不用显式实现接口，实现只要具有接口的方法即可。（隐式的鸭子类型）
type Entry interface {
	/**
	 * 寻找并加载 class 文件
	 * 
	 * 入参：
	 * 	className: class 文件的相对路径，形如 java/lang/Object.class
	 *
	 * 返回：
	 * 	[]byte：读取到的字节数据
	 * 	Entry：最终定位到 class 文件的 Entry
	 * 	error：错误信息
	 */
	readClass(className string) ([]byte, Entry, error)
	// 返回变量的字符串表示
	String() string
}

// 根据参数创建不同的 Entry
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	} else if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	} else if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || 
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path) // 目录形式的类路径
}
