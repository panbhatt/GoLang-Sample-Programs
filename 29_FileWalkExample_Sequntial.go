// Problem

/**
we need to iterate through the folder and determien which two files are
same. E.g two files are same if they same hash. i.e. binary content is same, no matter what is the
name of the file and path is . 

*/ 

// THIS IS SEQUENTIAL SOLUTION

package main

import (
	"fmt"
	"crypto/md5"
	"io"
	"log"
	"os"
	"path/filepath"

)

type pair struct {
	hash, path string
}

type fileList []string

type results map[string]fileList


func main() {

	if(len(os.Args)  < 2)  {
		log.Fatal(" Missing arguments, Please provide directory Name. ")
	}

	hashes, err := searchTree(os.Args[1])
	
	// for _, files := range hashes {
	// 	fmt.Println( len(files), files)
	// }

	if err== nil {
		for hash, files := range hashes {

			if len(files) > 1 {
				// Print the file name. 
				fmt.Println(hash, " -> " , len(files))

				for _, file := range files {
					fmt.Println("  ", file)
				}
			} 

		}
	} else{
		log.Fatal(" An Error occured, ", err)
	}

}



// This function will search the direct of directories
func searchTree(dir string) (results, error) {

	hashes := make(results)

	err :=	filepath.Walk(dir, func(p string, fi os.FileInfo, err error) error {

		if(fi.Mode().IsRegular() && fi.Size() > 0) {

			h := hashFile(p)
			hashes[h.hash] = append(hashes[h.hash], h.path)

		}

		return nil

	})

	return hashes, err



}

// THis function is going to return the Hash of the file with its path
func hashFile(path string) pair {

	fl, err := os.Open(path)

	if err!= nil {
		log.Fatal("Unable to read the file" , err)
	}

	hashStream := md5.New()

	if _,err := io.Copy(hashStream, fl); err != nil {
		log.Fatal( "AN Error occured, while reading the File",fl )
	}

	filePath := pair{
		hash : fmt.Sprintf("%x", hashStream.Sum(nil)),
		path : path, 
	}

	//fmt.Printf("%v\n", filePath)

	return filePath

}

