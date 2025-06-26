package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var intVar int
	var int8Var int8
	var int16Var int16
	var int32Var int32
	var runeVar rune
	var int64Var int64

	var uintVar uint
	var uint8Var uint8
	var uint16Var uint16
	var uint32Var uint32
	var uint64Var uint64
	var uintptrVar uintptr

	var float32Var float32
	var float64Var float64

	var complex64Var complex64
	var complex128Var complex128

	var byteVar byte
	var stringVar string = "Hello, World!"

	var boolVar bool

	fmt.Println("Размеры типов данных (в байтах):")
	fmt.Println("Целые числа:")
	fmt.Println("int:", unsafe.Sizeof(intVar))
	fmt.Println("int8:", unsafe.Sizeof(int8Var))
	fmt.Println("int16:", unsafe.Sizeof(int16Var))
	fmt.Println("int32:", unsafe.Sizeof(int32Var))
	fmt.Println("rune:", unsafe.Sizeof(runeVar))
	fmt.Println("int64:", unsafe.Sizeof(int64Var))
	fmt.Println("\nБеззнаковые целые:")
	fmt.Println("uint:", unsafe.Sizeof(uintVar))
	fmt.Println("uint8:", unsafe.Sizeof(uint8Var))
	fmt.Println("uint16:", unsafe.Sizeof(uint16Var))
	fmt.Println("uint32:", unsafe.Sizeof(uint32Var))
	fmt.Println("uint64:", unsafe.Sizeof(uint64Var))
	fmt.Println("uintptr:", unsafe.Sizeof(uintptrVar))
	fmt.Println("\nЧисла с плавающей точкой:")
	fmt.Println("float32:", unsafe.Sizeof(float32Var))
	fmt.Println("float64:", unsafe.Sizeof(float64Var))
	fmt.Println("\nКомплексные числа:")
	fmt.Println("complex64:", unsafe.Sizeof(complex64Var))
	fmt.Println("complex128:", unsafe.Sizeof(complex128Var))
	fmt.Println("\nБайт и строка:")
	fmt.Println("byte:", unsafe.Sizeof(byteVar))
	fmt.Println("string:", unsafe.Sizeof(stringVar), "(структура строки)")
	fmt.Println("\nЛогический тип:")
	fmt.Println("bool:", unsafe.Sizeof(boolVar))
}
