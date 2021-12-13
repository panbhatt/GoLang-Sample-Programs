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
	limits := make(chan bool, workers )
	pairs := make(chan pair, workers)
	
	results := make(chan results)

	go collectHashes(pairs, results)

	wg.Add(1)
	err := walkDir(dir, pairs,wg, limits)

	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()

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

func walkDir(dir string, pairs chan<- pair,wg *sync.WaitGroup, limits chan bool)  error {

defer wg.Done()

visit := func(p string, fi os.FileInfo, err error) error {
		if fi.Mode().IsDir() &&  p != dir {
				wg.Add(1)
				go walkDir(p, pairs,wg, limits)
				return filepath.SkipDir
		}

		if fi.Mode().IsRegular() && fi.Size() > 0 {
			wg.Add(1)
			go processFile(p, pairs, wg, limits)
		}

		return nil
}

limits <- true
defer func() {
	<-limits
}()

return filepath.Walk(dir, visit)

}

func processFile(path string, pairs chan<- pair, wg *sync.WaitGroup, limits chan bool){

	defer wg.Done()

	limits <- true

	defer func() {
		<-limits
	}()

	pairs <- hashFile(path)
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