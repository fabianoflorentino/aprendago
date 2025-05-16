// Package documentacao provides utilities and helper functions for displaying
// documentation-related information and guidance, particularly for Go projects.
// It includes features to assist users in understanding documentation flags,
// commands, and best practices for writing Go documentation.
package documentacao

import (
	base "github.com/fabianoflorentino/aprendago/pkg/base_content"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// Help exibe informações de ajuda relacionadas ao capítulo de documentação,
// incluindo descrições de flags e comandos úteis para documentação em Go.
func Help() {
	h := []format.HelpMe{
		{Flag: flagIntroducaoDocumentacao, Description: descIntroducaoDocumentacao},
		{Flag: flagGoDoc, Description: descGoDoc},
		{Flag: flagCommandGoDoc, Description: descCommandGoDoc},
		{Flag: flagPkgGoDev, Description: descPkgGoDev},
		{Flag: flagEscrevendoDocumentacao, Description: descEscrevendoDocumentacao},
	}

	b := base.New()
	b.HelpMe("Capítulo 25: Documentação", h)
}
