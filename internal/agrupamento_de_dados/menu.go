package agrupamento_de_dados

import (
	// "go/format" removed as it does not contain MenuOptions

	base "github.com/fabianoflorentino/aprendago/internal/base_content"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func Menu(menu []string) []format.MenuOptions {
	m := base.New()

	newFlag := []string{
		flagArray,
		flagSliceLiteralComposta,
		flagSliceForRange,
		flagSliceFatiandoOuDeletando,
		flagSliceAnexando,
		flagSliceMake,
		flagSliceMultiDimensional,
		flagSliceSurpresaArraySubjacente,
		flagMapsIntroducao,
		flagMapsRangeEDeletando,
	}

	newExecFunc := []func(){
		func() { m.SectionFormat(array, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(sliceLiteralComposta, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(sliceForRange, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(sliceFatiandoOuDeletando, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(sliceAnexando, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(sliceMake, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(sliceMultiDimensional, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(sliceSurpresaArraySubjacente, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(mapsIntroducao, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(mapsRangeEDeletando, m.SectionFactory(rootDir)) },
	}

	return m.Menu(newFlag, newExecFunc)
}
