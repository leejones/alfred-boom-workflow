package main

import "fmt"
import "flag"
import "os"
import "regexp"
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

  fmt.Println("<items>")
  if len(arguments) == 0 {
    for _, name := range boom.ListNames(data) {
      fmt.Printf("<item arg=\"%v\" valid=\"no\" autocomplete=\"%v\">\n", name, name)
      fmt.Printf("<title>%v</title>\n",name)
      fmt.Println("</item>")
    }
  } else if len(arguments) == 1 {
    if (len(boom.FetchListSnippetNamesFor(data, arguments[0])) >= 1) {
      for _, name := range boom.FetchListSnippetNamesFor(data, arguments[0]) {
        fmt.Println(name)
      }
    } else {
      for _, name := range boom.ListNames(data) {
        matched, _ := regexp.MatchString(arguments[0], name)
        if matched {
          fmt.Println("Partial match:", name)
        }
      }
    }
  } else if len(arguments) == 2 {
    snippet := boom.FetchSnippet(data, arguments[0], arguments[1])
    if snippet != "" {
      fmt.Println(snippet)
    } else {
      for _, name := range boom.FetchListSnippetNamesFor(data, arguments[0]) {
        matched, _ := regexp.MatchString(arguments[1], name)
        if matched {
          snippet := boom.FetchSnippet(data, arguments[0], name)
          fmt.Println("Partial match:", name)
          fmt.Println("Snippet:", snippet)
        }
      }
    }
  }
  fmt.Println("</items>")
}
