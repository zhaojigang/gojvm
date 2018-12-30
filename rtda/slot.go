package rtda

type Slot struct {
	num int32   // 基本类型（所有的基本类型都可以使用 int32 来处理）
	ref *Object // 引用类型
}
