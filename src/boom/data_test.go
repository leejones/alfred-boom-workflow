package boom

import "testing"
import "os"

func Fixture(name string) string {
  cwd, _ := os.Getwd()
  fixture_dir := cwd + "/../../test/fixtures"
  return fixture_dir + "/" + name
}

func TestFileLoad(t *testing.T) {
  path := Fixture("boom")
  data := ParseBoomDataFile(path)
  got := len(data)
  expected := 3
  if  got != expected {
    t.Errorf("Wanted %v, %v", expected, got)
  }
}

func TestListNames(t *testing.T) {
  path := Fixture("boom")
  data := ParseBoomDataFile(path)
  got := ListNames(data)[1]
  expected := "video" 
  if  got != expected {
    t.Errorf("Wanted %v, %v", expected, got)
  }
}

func TestFetchListSnippetNames(t *testing.T) {
  path := Fixture("boom")
  data := ParseBoomDataFile(path)
  got := FetchListSnippetNamesFor(data, "img")[2]
  expected := "make-it-so" 
  if  got != expected {
    t.Errorf("Wanted %v, %v", expected, got)
  }
}

func TestFetchSnippet(t *testing.T) {
  path := Fixture("boom")
  data := ParseBoomDataFile(path)
  got := FetchSnippet(data, "img", "i-got-this")
  expected := "http://dl.dropboxusercontent.com/u/152933/important-documents/i-got-this-01.jpg" 
  if  got != expected {
    t.Errorf("Wanted %v, %v", expected, got)
  }
}
