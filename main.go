package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/ogier/pflag"
)

var (
	user string
	// emoji string
)

func main() {
	// parse flag
	pflag.Parse()
	if pflag.NFlag() == 0 {
		printUsage()
	}

	users := strings.Split(user, ",")
	fmt.Printf("Search user(s): %s\n", users)
	for _, u := range users {
		result := getUsers(u)
		color.Cyan(`Username: %s`, result.Login)
		color.Blue(`Name: %s`, result.Name)
		color.Green(`Company: %s`, result.Company)
		color.HiMagenta(`Bio: %s`, result.Bio)
		color.Yellow(`Location: %s`, result.Location)
		color.Yellow(`Blog: %s`, result.Blog)
		fmt.Println()
	}
}

func init() {
	pflag.StringVarP(&user, "user", "u", "", "Search Users")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	pflag.PrintDefaults()
	os.Exit(1)
}
