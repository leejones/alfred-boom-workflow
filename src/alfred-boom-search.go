package main

import "fmt"
import "flag"
import "os"
import "regexp"
import "strings"
import "boom"
import "alfred"

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

  arguments := strings.Fields(flag.Args()[0])
  data := boom.ParseBoomDataFile(storage_path)

  alfred.PresentItemsOpen()

  if len(arguments) == 1 {
    if (len(boom.FetchListSnippetNamesFor(data, arguments[0])) >= 1) {
      if arguments[0] == flag.Args()[0] {
        item := alfred.Item{
          Argument: arguments[0],
          Valid: "no",
          Autocomplete: arguments[0] + " ",
          Title: arguments[0],
          Subtitle: "Search the " + arguments[0] + " list...",
        }
        alfred.PresentItem(item)
      } else {
        for _, name := range boom.FetchListSnippetNamesFor(data, arguments[0]) {
          snippet := boom.FetchSnippet(data, arguments[0], name)
          item := alfred.Item{
            Argument: arguments[0] + " " + name,
            Valid: "no",
            Autocomplete: arguments[0] + " " + name + " ",
            Title: name,
            Subtitle: "Copy " + snippet + " to your clipboard",
          }
          alfred.PresentItem(item)
        }
      }
    } else {
      for _, name := range boom.ListNames(data) {
        matched, _ := regexp.MatchString(arguments[0], name)
        if matched {
          item := alfred.Item{
            Argument: name,
            Valid: "no",
            Autocomplete: name + " ",
            Title: name,
            Subtitle: "Search the " + name + " list...",
          }
          alfred.PresentItem(item)
        }
      }
    }
  } else if len(arguments) == 2 {
    snippet := boom.FetchSnippet(data, arguments[0], arguments[1])
    if snippet != "" {
      item := alfred.Item{
        Argument: snippet,
        Valid: "yes",
        Autocomplete: arguments[0] + " " + arguments[1] + " ",
        Title: arguments[1],
        Subtitle: "Copy " + snippet + " to your clipboard",
      }
      alfred.PresentItem(item)
    } else {
      for _, name := range boom.FetchListSnippetNamesFor(data, arguments[0]) {
        matched, _ := regexp.MatchString(arguments[1], name)
        if matched {
          snippet := boom.FetchSnippet(data, arguments[0], name)
          item := alfred.Item{
            Argument: name,
            Valid: "no",
            Autocomplete: arguments[0] + " " + name + " ",
            Title: name,
            Subtitle: "Copy " + snippet + " to your clipboard",
          }
          alfred.PresentItem(item)
        }
      }
    }
  } else {
    for _, name := range boom.ListNames(data) {
      item := alfred.Item{
        Argument: name,
        Valid: "no",
        Autocomplete: name + " ",
        Title: name,
        Subtitle: "Search the " + name + " list...",
      }
      alfred.PresentItem(item)
    }
  }

  alfred.PresentItemsClose()
}
