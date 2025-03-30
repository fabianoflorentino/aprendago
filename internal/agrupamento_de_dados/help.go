// Help provides a comprehensive guide to the available flags and their descriptions
// for Chapter 8: "Agrupamento de Dados". This chapter covers topics related to arrays,
// slices, and maps, including their usage, operations, and specific scenarios. The
// function organizes the flags and their corresponding descriptions using the
// format.HelpMe structure and outputs the information using the format.PrintHelpMe
// function.
package agrupamento_de_dados

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// help provides a list of available flags and their descriptions for Chapter 8:
// "Agrupamento de Dados". It displays topics related to arrays, slices, and maps,
// including their usage, operations, and specific scenarios. The function utilizes
// the format.HelpMe structure to organize the flags and their corresponding descriptions,
// and outputs the information using the format.PrintHelpMe function.
func Help() {
	h := []format.HelpMe{
		{Flag: flagArray, Description: descArray},
		{Flag: flagSliceLiteralComposta, Description: descSliceLiteralComposta},
		{Flag: flagSliceForRange, Description: descSliceForRange},
		{Flag: flagSliceFatiandoOuDeletando, Description: descSliceFatiandoOuDeletando},
		{Flag: flagSliceFatiandoOuDeletandoRes, Description: descSliceFatiandoOuDeletandoRes},
		{Flag: flagSliceAnexando, Description: descSliceAnexando},
		{Flag: flagSliceMake, Description: descSliceMake},
		{Flag: flagSliceMultiDimensional, Description: descSliceMultiDimensional},
		{Flag: flagSliceSurpresaArraySubjacente, Description: descSliceSurpresaArraySubjacente},
		{Flag: flagMapsIntroducao, Description: descMapsIntroducao},
		{Flag: flagMapsRangeEDeletando, Description: descMapsRangeEDeletando},
	}

	fmt.Println("Capítulo 8: Agrupamento de Dados")
	format.PrintHelpMe(h)
}
