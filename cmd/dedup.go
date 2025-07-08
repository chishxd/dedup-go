package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/chishxd/dedup-go/internal/hasher"
	"github.com/chishxd/dedup-go/internal/scanner"
)

func main(){
    var dirFlag = flag.String("path", ".", "Path to the directory to dedup")
    flag.Parse()
	files, err := scanner.WalkFiles(*dirFlag)
	if err != nil{
		log.Fatal(err)
	}
	//Declaring hash Map
	m:= make(map[string][]string)
	// Iterating through each file, generating Hash for them and adding them to map
	for _, f := range files{
		hash, err := hasher.HashFiles(f)
		if err != nil {
			log.Println("Error hashing: ", f, err)
			continue
		} 
		m[hash] = append(m[hash], f)
	}
	for hash, files := range m{
		if len(files) > 1{
			fmt.Println(hash)
			fmt.Println(files)
		}
	}

}