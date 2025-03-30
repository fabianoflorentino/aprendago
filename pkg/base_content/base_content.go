// Package basecontent provides foundational structures and functionalities
// that can be extended or embedded by other types to enable shared behavior
// and common operations. It includes utilities for generating menu options,
// formatting sections, and creating section instances, making it a versatile
// component for building modular and reusable Go applications.
package basecontent

import (
	"github.com/fabianoflorentino/aprendago/pkg/format"
	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/section"
)

// MenuOptions defines an interface for handling menu options.
// It provides a method to generate a list of menu options based on a flag
// and a corresponding function to execute for the selected option.
//
// Menu takes two parameters:
// - optionFlag: A string representing the flag or identifier for the menu option.
// - optionFunc: A function to be executed when the menu option is selected.
//
// The method returns a slice of format.MenuOptions, representing the available menu options.
type MenuOptions interface {
	Menu(optionFlag string, optionFunc func()) []format.MenuOptions
}

// BaseContent represents a foundational structure or entity that can be
// extended or embedded by other types to provide shared functionality
// or common behavior.
type BaseContent struct{}

// New creates and returns a new instance of BaseContent.
// This function initializes a BaseContent struct with default values.
func New() *BaseContent {
	return &BaseContent{}
}

// Menu generates a list of menu options based on the provided flags and corresponding functions.
// Each menu option is represented as a format.MenuOptions object.
//
// Parameters:
//   - optionFlags: A slice of strings representing the flags or labels for the menu options.
//   - optionFuncs: A slice of functions corresponding to each menu option, defining the actions to be executed.
//
// Returns:
//   - A slice of format.MenuOptions containing the generated menu options.
func (b *BaseContent) Menu(optionFlags []string, optionFuncs []func()) []format.MenuOptions {
	menuOptions := []format.MenuOptions{}

	for i := 0; i < len(optionFlags) && i < len(optionFuncs); i++ {
		menuOptions = append(menuOptions, format.MenuOptions{Options: optionFlags[i], ExecFunc: optionFuncs[i]})
	}

	return menuOptions
}

// SectionFormat formats a section with the given title and applies it to the BaseContent.
// It takes a title string and a pointer to a Section object as parameters.
// The function modifies the BaseContent based on the provided section details.
func (b *BaseContent) SectionFormat(title string, section *section.Section) {
	section.Format(title)
}

// SectionFactory creates and returns a new Section instance based on the provided root directory.
// It utilizes the BaseContent receiver to configure or initialize the Section as needed.
//
// Parameters:
//   - rootDir: The root directory path used to initialize the Section.
//
// Returns:
//   - A pointer to a Section instance configured with the given root directory.
func (b *BaseContent) SectionFactory(rootDir string) *section.Section {
	section, err := section.New(rootDir)
	if err != nil {
		logger.New("Error creating section: ", err.Error()).Register()
		return nil
	}

	return section
}
