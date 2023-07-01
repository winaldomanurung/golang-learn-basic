package main

import (
	"fmt"
	"strconv"
)

func main() {
	// INTEGER
	// int8 –128 to 127
	// int16 –32768 to 32767
	// int32 –2147483648 to 2147483647
	// int64 –9223372036854775808 to 9223372036854775807
	// uint8 0 to 255
	// uint16 0 to 65536
	// uint32 0 to 4294967295
	// uint64 0 to 18446744073709551615
	var x int = 10
	x *= 3

	fmt.Println(x)

	//FLOAT
	// float32 3.40282346638528859811704183484516925440e+38 - 1.401298464324817070923729583289916131280e-45
	// float64 1.797693134862315708145274237317043567981e+308 - 4.940656458412465441765687928682213723651e-324
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(d)

	//PARSING
	//With ParseFloat, this 64 tells how many bits of precision to parse.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// For ParseInt, the 0 means infer the base from the string. 64 requires that the result fit in 64 bits.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	e, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(e)

}
