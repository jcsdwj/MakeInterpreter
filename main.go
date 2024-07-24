/**
  @author: 伍萬
  @since: 2023/11/28
  @desc: //TODO
**/

package main

import (
	"fmt"
	"lexer/repl"
	"os"
	user2 "os/user"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands!\n")
	repl.Start(os.Stdin, os.Stdout)
}

// 输入
// let add = fn(x,y){ x+y; };
