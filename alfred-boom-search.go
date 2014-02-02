package main

import "fmt"
import "flag"
import "os"
import "io/ioutil"
import "encoding/json"


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

  boom_data, _ := ioutil.ReadFile(storage_path)

  type Snippet map[string]string

  type List map[string][]Snippet

  type ListOfLists map[string][]List

  var m ListOfLists
  json.Unmarshal(boom_data, &m)

  for _, list := range m["lists"] {
    for list_name, snippets := range list {
      fmt.Println(list_name)
      for _, snippet := range snippets {
        for snippet_name, snippet_value := range snippet {
          fmt.Println("  ", snippet_name)
          fmt.Println("  ", "  ", snippet_value)
        }
      }
    }
  }

  fmt.Println("Arguments:")
  for _, argument := range arguments {
    fmt.Println(" ", argument)
  }

  fmt.Println("Storage Path:")
  fmt.Println(" ", storage_path)
}
