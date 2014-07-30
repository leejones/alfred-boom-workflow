package alfred

import "fmt"

type Item struct {
  Argument string
  Valid string
  Autocomplete string
  Title string
  Subtitle string
}

func PresentItemsOpen() {
  fmt.Println("<items>")
}

func PresentItemsClose() {
  fmt.Println("</items>")
}

func PresentItem(item Item) {
  fmt.Printf("<item arg=\"%v\" valid=\"%v\" autocomplete=\"%v\">\n", item.Argument, item.Valid, item.Autocomplete)
  fmt.Printf("<title>%v</title>\n", item.Title)
  fmt.Printf("<subtitle>%v</subtitle>\n", item.Subtitle)
  fmt.Println("</item>")
}

