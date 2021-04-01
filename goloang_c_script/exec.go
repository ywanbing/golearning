package goloang_c_script

/*
   #include<stdlib.h>
   int  execCmd(char* name)
   {
   	return system(name);
   }
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func ExecCmd(name string, args ...string) error {
	arg := ""
	for _, s := range args {
		arg += s + " "
	}

	cmd := C.CString(name + " " + arg)
	defer C.free(unsafe.Pointer(cmd))

	code := C.execCmd(cmd)
	if code != 0 {
		return fmt.Errorf("exec cmd failed , %v", cmd)
	}
	return nil
}
