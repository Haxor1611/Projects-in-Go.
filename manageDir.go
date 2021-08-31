package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	for {
		printOptions()
		var choice int
		fmt.Scanln(&choice)
		if choice == 1 {
			deleteDir()
		} else if choice == 2 {
			createDir()
		} else if choice == 3 {
			moveDir()
		}
	}
}
func deleteDir() {
	clearCon()
	var dir string
	fmt.Println("Enter dir: ")
	fmt.Scanln(&dir)
	err := os.RemoveAll(dir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Deleted.")
}

func printOptions() {
	fmt.Println("1. Delete a directory \n2. Create a directory \n3. Move directory")
}
func clearCon() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func createDir() {
	clearCon()
	var name string
	fmt.Println("Name: ")
	fmt.Scanln(&name)
	err := os.Mkdir(name, 0755)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Created.")
}
func moveDir() {
	clearCon()
	var oldLoc string
	var newLoc string
	fmt.Println("Old Location: ")
	fmt.Scanln(&oldLoc)
	fmt.Println("New Location: ")
	fmt.Scanln(&newLoc)
	err := os.Rename(oldLoc, newLoc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Moved.")
}
