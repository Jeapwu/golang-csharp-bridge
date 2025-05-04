package main

/*
#cgo linux LDFLAGS: -L. -ldl
#cgo darwin LDFLAGS: -L. -ldl
#include "bridge.h"
*/
import "C"
import (
    "fmt"
    "os"
    "path/filepath"
    "runtime"
    "unsafe"
)

func getLibName(base string) string {
    switch runtime.GOOS {
    case "windows":
        return base + ".dll"
    case "darwin":
        return "lib" + base + ".dylib"
    default: // linux
        return "lib" + base + ".so"
    }
}

func callAdd(libPath string, a, b int) int {
    cpath := C.CString(libPath)
    defer C.free(unsafe.Pointer(cpath))

    handle := C.load_library(cpath)
    if handle == nil {
        fmt.Printf("Failed to load: %s\n", libPath)
        os.Exit(1)
    }
    defer C.close_library(handle)

    cname := C.CString("add")
    defer C.free(unsafe.Pointer(cname))

    symbol := C.load_symbol(handle, cname)
    if symbol == nil {
        fmt.Println("Failed to find symbol 'add'")
        os.Exit(1)
    }

	addFunc := (C.add_func)(symbol)
	result := C.call_add(addFunc, C.int(3), C.int(5))
	return int(result)
}

func main() {
    basePath := "." // 可根据需要更改

    lib1 := filepath.Join(basePath, getLibName("CalculatorV1"))
    lib2 := filepath.Join(basePath, getLibName("CalculatorV2"))

    r1 := callAdd(lib1, 3, 5)
    fmt.Println("V1: 3 + 5 =", r1)

    r2 := callAdd(lib2, 3, 5)
    fmt.Println("V2: 3 + 5 =", r2)
}