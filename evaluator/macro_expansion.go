/**
  @author: 伍萬
  @since: 2024/1/26
  @desc: //TODO
**/

package evaluator

import (
	"lexer/ast"
	"lexer/object"
)

func DefineMacros(program *ast.Program, env *object.Environment) {
	definitions := []int{}

	for i, statement := range program.Statements {
		if isMacroDefinition(statement) {
			addMacro(statement, env)
			definitions = append(definitions, i)
		}
	}

	for i := len(definitions) - 1; i >= 0; i = i - 1 {
		definitionIndex := definitions[i]
		program.Statements = append(program.Statements[:definitionIndex], program.Statements[definitionIndex+1:]...)
	}
}

func addMacro(statement ast.Statement, env *object.Environment) {

}

func isMacroDefinition(statement ast.Statement) bool {
	return true
}
