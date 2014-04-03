package main

import "fmt"
import "flag"
import "os"
import "boom"


func main() {
  flag.Usage = func() {
    fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
    fmt.Println("  ", os.Args[0], "[options]", "[search term(s)]", "\n")
    fmt.Println("Options:")
    flag.PrintDefaults()
    fmt.Println("")
  }

  var storage_path string
  flag.StringVar(&storage_path, "storage_path", os.Getenv("HOME") + "/.boom", "boom storage path")

  flag.Parse()

  arguments := flag.Args()
  data := boom.ParseBoomDataFile(storage_path)

  if len(arguments) == 0 {
    for _, name := range boom.ListNames(data) {
      fmt.Println(name)
    }
  } else if len(arguments) == 1 {
    for _, name := range boom.FetchListSnippetNamesFor(data, arguments[0]) {
      fmt.Println(name)
    }
  } else if len(arguments) == 2 {
    fmt.Println(boom.FetchSnippet(data, arguments[0], arguments[1]))
  }
}
