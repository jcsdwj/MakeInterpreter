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
	fn *object.CompiledFunction
	ip int
}

func NewFrame(fn *object.CompiledFunction) *Frame {
	return &Frame{fn: fn, ip: -1}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
