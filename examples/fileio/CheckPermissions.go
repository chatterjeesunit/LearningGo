package main

import (
	"fmt"
	"os"
)

type PermissionBits uint32

const (
	UserRead PermissionBits = 1 << (8 - iota)
	UserWrite
	UserExecute
	GroupRead
	GroupWrite
	GroupExecute
	OtherRead
	OtherWrite
	OtherExecute
)

func main() {
	file := "/Users/in-sunit.chatterjee/learn/GitHub/go/src/LearningGo/examples/fileio/demo.txt"
	info,_ := os.Stat(file)
	mode := info.Mode()

	fmt.Printf("mode is %s and perm is %d \n" ,  mode, mode.Perm())

	perm := PermissionBits(os.FileMode(int(0077)))

	fmt.Printf("Read Permissions : User = %v, Group = %v, Others = %v\n",
		perm.hasPermissions(UserRead), perm.hasPermissions(GroupRead), perm.hasPermissions(OtherRead))

}

func (p *PermissionBits) hasPermissions(permission PermissionBits) bool {
	return *p & permission != 0
}