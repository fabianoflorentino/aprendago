package agrupamento_de_dados

// rootDir represents the relative path to the directory where data grouping topics are stored.
// This constant is used to reference the internal directory structure within the project.
const (
	rootDir = "internal/agrupamento_de_dados"
)

// The following constants define the names of various topics related to data grouping in Go.
// Each constant corresponds to a specific topic and provides a human-readable name for it.
const (
	array                        string = "Array"
	sliceLiteralComposta         string = "Slice: literal composta"
	sliceForRange                string = "Slice: for range"
	sliceFatiandoOuDeletando     string = "Slice: fatiando ou deletando de uma fatia"
	sliceAnexando                string = "Slice: anexando a uma slice"
	sliceMake                    string = "Slice: make"
	sliceMultiDimensional        string = "Slice: multi dimensional"
	sliceSurpresaArraySubjacente string = "Slice: a surpresa do array subjacente"
	mapsIntroducao               string = "Maps: introdução"
	mapsRangeEDeletando          string = "Maps: range e deletando"
)

// The following constants define command-line flags for various topics related to data grouping in Go.
// Each flag corresponds to a specific topic and is used to trigger functionality or operations
// associated with that topic in the application.
const (
	flagArray                        string = "--array"
	flagSliceLiteralComposta         string = "--slice-literal-composta"
	flagSliceForRange                string = "--slice-for-range"
	flagSliceFatiandoOuDeletando     string = "--slice-fatiando-ou-deletando-de-uma-fatia"
	flagSliceAnexando                string = "--slice-anexando-a-uma-slice"
	flagSliceMake                    string = "--slice-make"
	flagSliceMultiDimensional        string = "--slice-multi-dimensional"
	flagSliceSurpresaArraySubjacente string = "--slice-a-surpresa-do-array-subjacente"
	flagMapsIntroducao               string = "--maps-introducao"
	flagMapsRangeEDeletando          string = "--maps-range-e-deletando"
)

// The following constants provide descriptions for various topics related to data grouping in Go.
// Each constant corresponds to a specific topic and gives a brief explanation of its purpose or content.
const (
	descArray                        string = "Apresenta o tópico Array."
	descSliceLiteralComposta         string = "Apresenta o tópico Slice Literal Composta."
	descSliceForRange                string = "Apresenta o tópico Slice: for range."
	descSliceFatiandoOuDeletando     string = "Apresenta o tópico Slice: fatiando ou deletando de uma fatia."
	descSliceFatiandoOuDeletandoRes  string = "Apresenta a resolução do tópico Slice: fatiando ou deletando de uma fatia."
	descSliceAnexando                string = "Apresenta o tópico Slice: anexando a uma slice."
	descSliceMake                    string = "Apresenta o tópico Slice: Make."
	descSliceMultiDimensional        string = "Apresenta o tópico Slice: Multi Dimensional."
	descSliceSurpresaArraySubjacente string = "Apresenta o tópico Slice: a surpresa do array subjacente."
	descMapsIntroducao               string = "Apresenta o tópico Maps: introdução."
	descMapsRangeEDeletando          string = "Apresenta o tópico Maps: Range e Deletando."
)
