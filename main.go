package main

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	fileroot := "/users/donmennerich/Projects/dmnyu/dlts-finding-aids-ead-sample-set-1/ead-files"

	for _,r := range []string{"archives", "tamwag", "fales"} {
		repoRoot := filepath.Join(fileroot, r)
		root, err := os.Open(repoRoot)
		if err != nil {
			panic(err)
		}

		files, err := root.Readdir(-1)
		for _,file := range files {
			path := filepath.Join(repoRoot, file.Name())

			out, err := exec.Command("xmllint", "--format", path).Output()
			if err != nil {
				panic(err)
			}
			targetFile, err := os.Create(filepath.Join(r, file.Name()))
			if err != nil {
				panic(err)
			}
			writer := bufio.NewWriter(targetFile)
			writer.Write(out)
			writer.Flush()
		}
	}
}
