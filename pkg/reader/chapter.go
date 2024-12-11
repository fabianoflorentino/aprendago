package reader

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/fabianoflorentino/aprendago/internal/entities"
	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"gopkg.in/yaml.v3"
)

const (
	CHAPTER_ROOT_DIR    = "internal/outlines"
	CHAPTER_ROOT_NAME   = "capitulo"
	CHAPTER_YML_FORMAT  = ".yml"
	CHAPTER_YAML_FORMAT = ".yaml"
)

// ReadChapterFile é utilizada para ler o arquivo YML do capítulo e retornar uma estrutura de Chapter
func ReadChapterFile(chapterNumber int) (entities.Chapter, error) {
	var docPath string
	var chapter entities.Chapter

	docPathYML := filepath.Join(CHAPTER_ROOT_DIR, CHAPTER_ROOT_NAME+strconv.Itoa(chapterNumber)+CHAPTER_YML_FORMAT)
	docPathYAML := filepath.Join(CHAPTER_ROOT_DIR, CHAPTER_ROOT_NAME+strconv.Itoa(chapterNumber)+CHAPTER_YAML_FORMAT)

	if _, err := os.Stat(docPathYML); err == nil {
		docPath = docPathYML
	} else if _, err := os.Stat(docPathYAML); err == nil {
		docPath = docPathYAML
	} else {
		logger.Log("File %s or %s not found", docPathYML, docPathYAML)
		return entities.Chapter{}, os.ErrNotExist
	}

	data, err := os.ReadFile(docPath)
	if err != nil {
		return entities.Chapter{}, err
	}

	err = yaml.Unmarshal(data, &chapter)
	if err != nil {
		return entities.Chapter{}, err
	}

	return chapter, nil
}
