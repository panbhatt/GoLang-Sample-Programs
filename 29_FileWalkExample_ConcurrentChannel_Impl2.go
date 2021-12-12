// This is the concurrent channel implementaiton.

package main

import (
	"fmt"
	"crypto/md5"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type pair struct {
	hash, path string
}

type fileList []string

type results map[string]fileList

func main(){

	if(len(os.Args)  < 2)  {
		log.Fatal(" Missing arguments, Please provide directory Name. ")
	}
	dir := os.Args[1]

	wg := new (sync.WaitGroup)

	workers := 2 * runtime.GOMAXPROCS(0)
	fmt.Println(" NO OF WORKERS = ", workers)
	paths := make(chan string)
	pairs := make(chan pair)
	done := make(chan bool)
	results := make(chan results)

	for i:=0;i<workers;i++{
		go processFiles(paths,pairs,done)
	}

	go collectHashes(pairs, results)

	wg.Add(1)
	err := walkDir(dir, paths,wg)

	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	// closing the paths channel to make we are done with the work
	close(paths)

	// watit for all the workers to be done
	for i:=0 ; i< workers ; i++{
		<- done
	}

	// When the workers are done we need to close the Pairs channel too
	close(pairs)

	hashes := <- results

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

func walkDir(dir string, paths chan<- string,wg *sync.WaitGroup)  error {

defer wg.Done()

visit := func(p string, fi os.FileInfo, err error) error {
		if fi.Mode().IsDir() &&  p != dir {
				wg.Add(1)
				go walkDir(p, paths,wg)
				return filepath.SkipDir
		}

		if fi.Mode().IsRegular() && fi.Size() > 0 {
			paths <- p
		}

		return nil
} 

return filepath.Walk(dir, visit)

}


func processFiles(paths <-chan string, pairs chan<- pair, done chan<-bool) {

	for pth := range paths {
		 pairs <- hashFile(pth)
	}

	done <- true

}

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



func collectHashes(pairs <-chan pair, result chan<- results) {

	hashes := make(results)

	for p := range pairs {
		hashes[p.hash] = append(hashes[p.hash], p.path)
	}

	result <- hashes

}