package main

import "fmt"

func main() {
  fmt.Println("<!DOCTYPE html>")
  fmt.Println("<html>")
  fmt.Println("<head></head>")
  fmt.Println("<body>")
  fmt.Println("Hello, world")
  fmt.Println("</body>")
  fmt.Println("</html>")


  // This file can be re-directed into an existing .html file and write
  // your programs output to a HTML document
  // use go install
  // then call the folder name  and redirect into a html file
  // [FolderName] > index.html
  // Opening the html file, you will see, "Hello, world"


}
