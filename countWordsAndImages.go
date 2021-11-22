package main

import (
	"fmt"
	"os"
	"bytes"
	"strings"
	"golang.org/x/net/html"
)

var raw = `
<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
	  <img src="xxx.jpg" width="104" height="142">
  </body>
</html>`


func main(){

	doc, err :=html.Parse(bytes.NewReader([]byte(raw)))

	if err != nil{ 
		fmt.Fprintf(os.Stderr, "Parse Failed %s", err)
		os.Exit(-1)
	}



	fmt.Println("HTML Successfully Read")
	
	words, pics := countWordAndImages(doc)

	fmt.Printf("WORDS =  %d PICS = %d", words, pics)


}


func countWordAndImages(doc *html.Node) (int,int) {

	var words, images int  

	visit(doc, &words, &images)
	return words, images


}

func visit(n *html.Node, words *int, pics *int)  {

	
	switch n.Type {
	case html.TextNode:
		if s := strings.TrimSpace(n.Data); len(s) > 0 {
			*words += len(strings.Fields(s))
		}

	case html.ElementNode:
		switch n.Data {
		case "img":
			*pics++
		}
	}

	// then visit all the children using recursion

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, pics)
	}


}