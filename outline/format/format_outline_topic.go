package outline

import "fmt"

type OutlineContent struct {
	Title   string
	Content string
}

func FormatOutlineTopic(topic OutlineContent) {
	fmt.Printf("%v \n %v\n", topic.Title, topic.Content)
}
