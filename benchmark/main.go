/**
  @author: 伍萬
  @since: 2024/2/14
  @desc: //TODO
**/

package benchmark

import (
	"flag"
	"fmt"
	"lexer/compiler"
	"lexer/evaluator"
	"lexer/lexer"
	"lexer/object"
	"lexer/parser"
	"lexer/vm"
	"time"
)

var engine = flag.String("engine", "vm", "use 'vm' or 'eval'")
var input = `
let fibonacci = fn(x){
	if (x==0){
		0
	}else{
		if (x==1){
			return 1;
		}else{
			fibonacci(x-1)+fibonacci(x-2);
		}
	}
};
fibonacci(35);
`

func main() {
	flag.Parse()

	var duration time.Duration
	var result object.Object

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if *engine == "vm" {
		comp := compiler.New()
		err := comp.Compile(program)

		if err == nil {
			fmt.Printf("compiler error:%s", err)
			return
		}

		machine := vm.New(comp.ByteCode())

		start := time.Now()

		err = machine.Run()
		if err != nil {
			fmt.Printf("vm error:%s", err)
			return
		}

		duration = time.Since(start)
		result = machine.LastPoppedStackElem()
	} else {
		env := object.NewEnvironment()
		start := time.Now()
		result = evaluator.Eval(program, env)
		duration = time.Since(start)
	}

	fmt.Printf("engine=%s,result=%s,duration=%s\n", *engine, result.Inspect(), duration)
}
