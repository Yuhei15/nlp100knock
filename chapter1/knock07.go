package chapter1

import (
	"flag"
	"fmt"
)

func StringsJoin(x, y, z string) string {

	flag.StringVar(&x, "x", x, "output string")
	flag.StringVar(&y, "y", y, "output string")
	flag.StringVar(&z, "z", z, "output string")
	flag.Parse()
	fmt.Println(x + "の時" + y + "は" + z)
	return fmt.Sprint(x + "の時" + y + "は" + z)
}
