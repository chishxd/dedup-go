package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
)

func main(){
    var dirFlag = flag.String("path", ".", "Path to the directory to scan")
    flag.Parse()
    err := filepath.WalkDir(*dirFlag, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        fmt.Println("Path:", path, "IsDir:", d.IsDir())
        return nil
    })
    if err != nil {
        fmt.Println("Error walking directory:", err)
    }

}
