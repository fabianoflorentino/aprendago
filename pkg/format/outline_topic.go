package format

import (
	"fmt"
	"strings"
)

type OutlineContent struct {
	Title   string
	Content string
}

func (outlineContent *OutlineContent) AddHeader(content string) {
	outlineContent.Content += strings.ToUpper(content) + "\n"
}

func (outlineContent *OutlineContent) AddText(content string) {
	outlineContent.Content += content + "\n"
}

func (outlineContent *OutlineContent) StartList(callback func(*List)) {
	list := List{level: 0}
	callback(&list)
	outlineContent.Content += list.render()
}

type listRendering interface {
	render() string
}

type List struct {
	level int
	items []listRendering
}

func (currentList List) render() string {
	var content string

	for _, listItem := range currentList.items {
		content += listItem.render()
	}

	return content
}

func (currentList *List) StartSubList(callback func(*List)) {
	childList := List{level: currentList.level + 1}
	currentList.items = append(currentList.items, &childList)
	callback(&childList)
}

func (list *List) AddItem(content string) {
	itemContent := itemContent{content, list}
	list.items = append(list.items, &itemContent)
}

type itemContent struct {
	content string
	list    *List
}

func (itemContent itemContent) render() string {
	tabs := strings.Repeat("  ", itemContent.list.level)
	return tabs + "- " + itemContent.content + "\n"
}

func FormatOutlineTopic(topic OutlineContent) {
	fmt.Printf("%v \n %v\n", topic.Title, topic.Content)
}

func StartSection(callback func(*OutlineContent)) {
	section := OutlineContent{}
	callback(&section)
	fmt.Println(section.Content)
}
