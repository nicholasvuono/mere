package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/nicholasvuono/mere/internal/pkg/repl"
)

const MEERKAT = `

     .''''''''.
   .'          '.
 .'          )  .
''....          .
      ''.      .
        .      .
        .      .
        .       .
      .'        .
    .'          .
   '  .         .
   ' ' .         .
   '.' .         .
       .         .
       .          .
       .          .
    .'            .
   .               '.
    .          .      ' .
     '..       . '.       ' ' ' ' ' . .
   .''         .    ' . . . . . . . ''
   '. . . . . .'
`

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println(MEERKAT)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("THE MERE PROGRAMMING LANGUAGE  |  nothing more than a simple, pure and familiar programming language")
	fmt.Println("____________________________________________________________________________________________________")
	fmt.Println()
	fmt.Printf("Hey %s, let's jump right in!\n", user.Username)
	fmt.Println()
	repl.Start(os.Stdin, os.Stdout)
}
