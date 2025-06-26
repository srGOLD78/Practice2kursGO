package main

import "fmt"

func main() {
	var intVar int = -78
	var int8Var int8 = 127
	var int16Var int16 = 32767
	var int32Var int32 = 2147483647
	var runeVar rune = 'Я'
	var int64Var int64 = 9223372036854775807

	var uintVar uint = 78
	var uint8Var uint8 = 255
	var uint16Var uint16 = 65535
	var uint32Var uint32 = 4294967295
	var uint64Var uint64 = 18446744073709551615
	var uintptrVar uintptr = 0xDEADBEEF

	var float32Var float32 = 3.1415926535
	var float64Var float64 = 3.14159265358979323846

	var complex64Var complex64 = 3 + 4i
	var complex128Var complex128 = 3 + 4i

	var byteVar byte = 'A'
	var stringVar string = "Hello, World!"

	var boolVar bool = true

	fmt.Println("Целые числа:")
	fmt.Println("int:", intVar)
	fmt.Println("int8:", int8Var)
	fmt.Println("int16:", int16Var)
	fmt.Println("int32:", int32Var)
	fmt.Println("rune:", runeVar, string(runeVar))
	fmt.Println("int64:", int64Var)
	fmt.Println("\nБеззнаковые целые:")
	fmt.Println("uint:", uintVar)
	fmt.Println("uint8:", uint8Var)
	fmt.Println("uint16:", uint16Var)
	fmt.Println("uint32:", uint32Var)
	fmt.Println("uint64:", uint64Var)
	fmt.Println("uintptr:", uintptrVar)
	fmt.Println("\nЧисла с плавающей точкой:")
	fmt.Println("float32:", float32Var)
	fmt.Println("float64:", float64Var)
	fmt.Println("\nКомплексные числа:")
	fmt.Println("complex64:", complex64Var)
	fmt.Println("complex128:", complex128Var)
	fmt.Println("\nБайт и строка:")
	fmt.Println("byte:", byteVar, string(byteVar))
	fmt.Println("string:", stringVar)
	fmt.Println("\nЛогический тип:")
	fmt.Println("bool:", boolVar)
}


