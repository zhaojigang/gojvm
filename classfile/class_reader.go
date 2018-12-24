package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

// u1 - 无符号1个字节（8位）
func (self *ClassReader) readUint8() uint8 {
	value := self.data[0]     // 读取第一个字节
	self.data = self.data[1:] // 原本的data去除被读取的字节
	return value
}

// u2 - 无符号2个字节（16位）
func (self *ClassReader) readUint16() uint16 {
	value := binary.BigEndian.Uint16(self.data) // 从data中读取16位（2个字节）
	self.data = self.data[2:]                   // 原本的data去除被读取的字节
	return value
}

// u4 - 无符号4个字节（32位）
func (self *ClassReader) readUint32() uint32 {
	value := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return value
}

// u8 - 无符号8个字节（64位）
func (self *ClassReader) readUint64() uint64 {
	value := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return value
}

// 读取interface表：
// 表头：eg. u2 interface_count
// 表项：eg. u2 interface[interface_count]
// 读取 exception_index_table 表
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16() // 表头：表的大小
	s := make([]uint16, n) // 创建存储表数据（表项）的容器
	for i := range s {
		s[i] = self.readUint16() // 读取表数据
	}
	return s
}

// 读取指定字节个数据
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
