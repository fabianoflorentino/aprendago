package entities

import (
	"fmt"
	"reflect"

	"github.com/fabianoflorentino/aprendago/internal/resolutions"
)

type HelpMe struct {
	Flag string `yaml:"flag,omitempty"`
	Span string `yaml:"span"`
}

type QuestionOption struct {
	Identifier string `yaml:"identifier"`
	Span       string `yaml:"span"`
}

type Question struct {
	Text        string           `yaml:"text"`
	RightAnswer string           `yaml:"right_answer"`
	Options     []QuestionOption `yaml:"options"`
}

type Topic interface {
	GetFlag() string
	GetHelpMes() []HelpMe
}

type DefaultTopic struct {
	Title   string   `yaml:"title"`
	Content string   `yaml:"content"`
	HelpMes []HelpMe `yaml:"help_me"`
}

func (dt DefaultTopic) GetFlag() string {
	return ""
}

func (dt DefaultTopic) GetHelpMes() []HelpMe {
	return dt.HelpMes
}

type Outline struct {
	*DefaultTopic `yaml:",inline"`
	Flag          string `yaml:"flag"`
}

type Exercise struct {
	*DefaultTopic `yaml:",inline"`
	Level         int                    `yaml:"level"`
	Number        int                    `yaml:"number"`
	resolution    resolutions.Resolution `yaml:"-"`
}

type Exam struct {
	*DefaultTopic `yaml:",inline"`
	Level         int        `yaml:"level"`
	Questions     []Question `yaml:"questions"`
}

type Chapter struct {
	Title     string     `yaml:"title"`
	Outlines  []Outline  `yaml:"outlines"`
	Exercises []Exercise `yaml:"exercises"`
	Exam      Exam       `yaml:"exam"`
}

func (o Outline) GetFlag() string {
	return o.Flag
}

func (e Exercise) GetFlag() string {
	return fmt.Sprintf("--exercicio=%d --nivel=%d", e.Number, e.Level)
}

func (e Exercise) ExecuteResolution() {
	resolution := reflect.ValueOf(e.resolution)
	methodName := fmt.Sprintf("Nivel%dExercicio%d", e.Level, e.Number)
	method := resolution.MethodByName(methodName)
	method.Call([]reflect.Value{})
}

func (e Exam) GetFlag() string {
	return fmt.Sprintf("--prova --nivel=%d", e.Level)
}

func (c Chapter) Topics() []Topic {
	topics := []Topic{}
	for _, o := range c.Outlines {
		topics = append(topics, o)
	}
	for _, e := range c.Exercises {
		topics = append(topics, e)
	}
	if c.Exam.DefaultTopic != nil {
		topics = append(topics, c.Exam)
	}
	return topics
}

func (c Chapter) SerializeHelpMes(t Topic) []map[string]string {
	helpMes := []map[string]string{}

	for _, helpMe := range t.GetHelpMes() {
		helpMes = append(helpMes, map[string]string{
			"flag": t.GetFlag() + " " + helpMe.Flag,
			"span": helpMe.Span,
		})
	}

	return helpMes
}