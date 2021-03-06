package heap

import "github.com/zhaojigang/gojvm/classfile"

type Field struct {
	ClassMember
	constantValueIndex uint
	slotId             uint
}

// 根据 classFile 创建 字段表
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constantValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}
func (self *Field) ConstantValueIndex() uint {
	return self.constantValueIndex
}

// getter
func (self *Field) SlotId() uint {
	return self.slotId
}
