/**
  @author: 伍萬
  @since: 2024/2/10
  @desc: //TODO
**/

package vm

import (
	"lexer/code"
	"lexer/object"
)

type Frame struct {
	fn          *object.CompiledFunction
	ip          int
	basePointer int
	cl          *object.Closure
}

//func NewFrame(fn *object.CompiledFunction, basePointer int) *Frame {
//	return &Frame{fn: fn, ip: -1, basePointer: basePointer}
//}

func NewFrame(cl *object.Closure, basePointer int) *Frame {
	f := &Frame{
		fn:          nil,
		ip:          -1,
		basePointer: basePointer,
		cl:          cl,
	}

	return f
}

func (f *Frame) Instructions() code.Instructions {
	// return f.fn.Instructions
	return f.cl.Fn.Instructions
}
