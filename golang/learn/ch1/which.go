package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//The logical work here will be divided into
// 1) reading the input argument, which is the name of the executable file that the utility will be searching for
// 2) reading the PATH environment variable, splitting it, and iterating over the directories of the PATH variable.
// 3) looking for the desired binary file in these directories and determining whether it can be found or not, whether
//it is a regular file, and whether it is an executable file

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Please provide an argument!")
		return
	}

	execFile := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	// fmt.Println(execFile)

	for _, directory := range pathSplit {

		fullPath := filepath.Join(directory, execFile)
		fileInfo, err := os.Stat(fullPath)

		if err == nil {
			// something like -rwxr-xr-x
			mode := fileInfo.Mode()

			//Is it a regular file and is it executable?
			if mode.IsRegular() && mode&0111 != 0 {
				fmt.Println(fullPath)
				return
			}
			// fmt.Println(directory, fileInfo, err, mode)
		}

	}
}
