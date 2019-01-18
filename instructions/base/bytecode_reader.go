package base

type BytecodeReader struct {
	code []byte // 字节码
	pc   int    // 记录读取到了哪个字节
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) ReadUint8() uint8 {
	byte := self.code[self.pc] // 读取一个字节
	self.pc++
	return byte
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

func (self *BytecodeReader) ReadUint16() uint16 {
	// 连续读取两个字节
	byte1 := uint16(self.ReadUint8()) // 高8位
	byte2 := uint16(self.ReadUint8()) // 低8位

	return (byte1 << 8) | byte2
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *BytecodeReader) ReadInt32() int32 {
	// 连续读取4个字节
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())

	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// Getter
func (self *BytecodeReader) PC() int {
	return self.pc
}
