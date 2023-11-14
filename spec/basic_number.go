package spec

import (
	"fmt"
)

func NumberDefineOnly() {
	// 仅声明
	// 有符号
	var intDefineOnly int
	var int8DefineOnly int8
	var int16DefineOnly int16
	var int32DefineOnly int32
	var int64DefineOnly int64
	// 无符号
	var uintDefineOnly uint
	var uint8DefineOnly uint8
	var uint16DefineOnly uint16
	var uint32DefineOnly uint32
	var uint64DefineOnly uint64

	fmt.Println("NumberDefineOnly",
		intDefineOnly, int8DefineOnly, int16DefineOnly, int32DefineOnly, int64DefineOnly,
		uintDefineOnly, uint8DefineOnly, uint16DefineOnly, uint32DefineOnly, uint64DefineOnly,
	)
}

func NumberDefineInit() {
	// 声明并初始化
	// 有符号
	var intDefineInit int = 10
	var int8DefineInit int8 = 10
	var int16DefineInit int16 = 10
	var int32DefineInit int32 = 10
	var int64DefineInit int64 = 10
	// 无符号
	var uintDefineInit uint = 10
	var uint8DefineInit uint8 = 10
	var uint16DefineInit uint16 = 10
	var uint32DefineInit uint32 = 10
	var uint64DefineInit uint64 = 10

	fmt.Println("NumberDefineInit",
		intDefineInit, int8DefineInit, int16DefineInit, int32DefineInit, int64DefineInit,
		uintDefineInit, uint8DefineInit, uint16DefineInit, uint32DefineInit, uint64DefineInit,
	)
}

func TypeInference() {
	// 方式一
	var numberVar = 10
	// 方式二
	numberShort := 20

	fmt.Println("TypeInference", numberVar, numberShort)
}

func Complex() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // "(‐5+10i)"
	fmt.Println(real(x * y))         // "‐5"
	fmt.Println(imag(x * y))         // "10"
}
