package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "User Name"

	str := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang = "en">
	<head>
	<meta charset="UTF-8">
	<title>Hello worlg go web</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>`)

	nf, err := os.Create("index.gohtml")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

}
