package mymod

import "fmt"
import "github.com/Leoyoungxh/innermod"

const (
	haha = "hahaha"
	vava = "vavava"
)

func Print() {
	fmt.Println("Version 1.3.0, require innermod 1.0.0")
	innermod.Innerfun()
}
