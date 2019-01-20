package heap

import "github.com/zhaojigang/gojvm/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method // 缓存解析后的方法指针
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
