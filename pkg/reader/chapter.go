package reader

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"gopkg.in/yaml.v2"
)

type Topic struct {
	Flag    string `yaml:"flag"`
	Title   string `yaml:"title"`
	HelpMe  string `yaml:"help_me"`
	Outline string `yaml:"outline"`
}

type Chapter struct {
	Title  string  `yaml:"title"`
	Topics []Topic `yaml:"topics"`
}

const (
	CHAPTER_ROOT_DIR    = "./internal/outlines"
	CHAPTER_ROOT_NAME   = "capitulo"
	CHAPTER_YML_FORMAT  = ".yml"
	CHAPTER_YAML_FORMAT = ".yaml"
)

// ReadChapterFile é utilizada para ler o arquivo YML do capítulo e retornar uma estrutura de Chapter
func ReadChapterFile(chapterNumber int) (Chapter, error) {
	var docPath string
	var chapter Chapter

	docPathYML := filepath.Join(CHAPTER_ROOT_DIR, CHAPTER_ROOT_NAME+strconv.Itoa(chapterNumber)+CHAPTER_YML_FORMAT)
	docPathYAML := filepath.Join(CHAPTER_ROOT_DIR, CHAPTER_ROOT_NAME+strconv.Itoa(chapterNumber)+CHAPTER_YAML_FORMAT)

	if _, err := os.Stat(docPathYML); err == nil {
		docPath = docPathYML
	} else if _, err := os.Stat(docPathYAML); err == nil {
		docPath = docPathYAML
	} else {
		logger.Log("File %s or %s not found", docPathYML, docPathYAML)
		return Chapter{}, os.ErrNotExist
	}

	data, err := os.ReadFile(docPath)
	if err != nil {
		return Chapter{}, err
	}

	err = yaml.Unmarshal(data, &chapter)
	if err != nil {
		return Chapter{}, err
	}

	return chapter, nil
}
