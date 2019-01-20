package heap

import "github.com/zhaojigang/gojvm/classfile"

// 方法符号引用
type MethodRef struct {
	MemberRef
	method *Method // 缓存解析后的方法指针
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
