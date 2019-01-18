package base

import (
	"github.com/zhaojigang/gojvm/rtda"
)

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextpc := pc + offset
	frame.SetNextPC(nextpc)
}
