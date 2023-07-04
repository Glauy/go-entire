package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/wailsapp/wails/v2/pkg/commands/bindings"
)

// workingDirectory
// projectDir
func main() {

	// Get the directory of this file
	_, filename, _, _ := runtime.Caller(0)
	cmdDirectory := filepath.Dir(filename)

	projectDir := filepath.Join(cmdDirectory, "..", "..")

	options := bindings.Options{
		ProjectDirectory: projectDir,
		GoModTidy:        true,
	}
	bindings.GenerateBindings(options)

	s, err := bindings.GenerateBindings(options)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}
