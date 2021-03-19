package mymod

import "fmt"
import "github.com/Leoyoungxh/innermod"

const (
	haha = "haha"
	vava = "vava"
)

func Print() {
	fmt.Println(haha, vava)
	innermod.Innerfun()
}
