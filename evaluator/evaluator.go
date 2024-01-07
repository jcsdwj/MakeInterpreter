/**
  @author: 伍萬
  @since: 2024/1/7
  @desc: //TODO
**/

package evaluator

import (
	"lexer/ast"
	"lexer/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}
