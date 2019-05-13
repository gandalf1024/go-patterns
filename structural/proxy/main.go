package main

import (
	"fmt"
)

// 代理模式
func main() {
	t, err := NewTerminal("gopher")
	if err != nil {
		panic(err.Error())
	}

	excResp, excErr := t.Execute("say_hi")
	if excErr != nil {
		fmt.Printf("ERROR: %s\n", excErr.Error())
	}

	fmt.Println(excResp)
}

type ITerminal interface {
	Execute(cmd string) (resp string, err error)
}

type GopherTerminal struct {
	User string
}

func (gt *GopherTerminal) Execute(cmd string) (resp string, err error) {
	prefix := fmt.Sprintf("%s@go_term$:", gt.User)

	switch cmd {
	case "say_hi":
		resp = fmt.Sprintf("%s Hi %s", prefix, gt.User)
	case "man":
		resp = fmt.Sprintf("%s Visit 'https://golang.org/doc/' for Golang documentation", prefix)
	default:
		err = fmt.Errorf("%s Unknown command", prefix)
	}

	return
}

type Terminal struct {
	currentUser    string
	gopherTerminal *GopherTerminal
}

//核心代码 - 构建对象
func NewTerminal(user string) (t *Terminal, err error) {
	if user == "" {
		err = fmt.Errorf("User cant be empty")
		return
	}

	if authErr := authorizeUser(user); authErr != nil {
		err = fmt.Errorf("You (%s) are not allowed to use terminal and execute commands", user)
		return
	}

	t = &Terminal{currentUser: user}

	return
}

func (t *Terminal) Execute(command string) (resp string, err error) {
	t.gopherTerminal = &GopherTerminal{User: t.currentUser}

	fmt.Printf("PROXY: Intercepted execution of user (%s), asked command (%s)\n", t.currentUser, command)

	if resp, err = t.gopherTerminal.Execute(command); err != nil {
		err = fmt.Errorf("I know only how to execute commands: say_hi, man")
		return
	}

	return
}

func authorizeUser(user string) (err error) {
	if user != "gopher" {
		err = fmt.Errorf("User %s in black list", user)
		return
	}
	return
}
