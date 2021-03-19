package mymod

import "fmt"
import "github.com/Leoyoungxh/innermod"

const (
	haha = "hahaha"
	vava = "vavava"
)

func Print() {
	fmt.Println(haha, vava)
	innermod.Innerfun()
}
