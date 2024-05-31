package outline

import "fmt"

type OulineContent struct {
	Title   string
	Content string
}

func FormatOutlineTopic(topic OulineContent) {
	fmt.Printf("%v \n %v\n", topic.Title, topic.Content)
}
