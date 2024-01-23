/**
  @author: 伍萬
  @since: 2024/1/24
  @desc: //TODO
**/

package evaluator

import (
	"lexer/ast"
	"lexer/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
