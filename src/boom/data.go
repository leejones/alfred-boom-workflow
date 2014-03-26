package boom

import "io/ioutil"
import "encoding/json"

type Snippet map[string]string

type List map[string][]Snippet

type ListOfLists map[string][]List

type Lists []List

func ParseBoomDataFile(storage_path string) Lists {
  boom_data, _ := ioutil.ReadFile(storage_path)
  var list_of_lists ListOfLists
  json.Unmarshal(boom_data, &list_of_lists)
  return lists(list_of_lists)
}

func ListNames(lists Lists) []string {
  names := make([]string, len(lists))
  for index, list := range lists {
    for list_name, _ := range list {
      names[index] = list_name
    }
  }
  return names
}

func FetchListSnippetNamesFor(lists Lists, passed_list_name string) []string {
  snippets := fetchSnippetsForList(lists, passed_list_name)
  snippet_names := make([]string, len(snippets))
  for index, snippet := range snippets {
    for snippet_name, _ := range snippet {
      snippet_names[index] = snippet_name
    }
  }
  return snippet_names 
}

func FetchSnippet(lists Lists, passed_list_name string, passed_snippet_name string) string {
  for _, snippet := range fetchSnippetsForList(lists, passed_list_name) {
    for snippet_name, snippet_value := range snippet {
      if snippet_name == passed_snippet_name {
        return snippet_value
      }
    }
  }
  return "" 
}

func fetchSnippetsForList(lists Lists, passed_list_name string) []Snippet {
  for _, list := range lists {
    for list_name, snippets := range list {
      if list_name == passed_list_name {
        return snippets
      }
    }
  }
  return make([]Snippet, 0)
}

func lists(lists_of_lists ListOfLists) []List {
  return lists_of_lists["lists"]
}
