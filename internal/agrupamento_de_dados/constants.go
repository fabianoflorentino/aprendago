package agrupamento_de_dados

// rootDir represents the relative path to the directory where data grouping topics are stored.
// This constant is used to reference the internal directory structure within the project.
const (
	rootDir = "internal/agrupamento_de_dados"
)

// Array represents a fixed-size sequence of elements of the same type.
// Arrays in Go have a fixed length, and their size is part of their type.
// They are useful when you know the exact number of elements you need to store.
const (
	Array                        string = "Array"
	SliceLiteralComposta         string = "Slice: literal composta"
	SliceForRange                string = "Slice: for range"
	SliceFatiandoOuDeletando     string = "Slice: fatiando ou deletando de uma fatia"
	SliceAnexando                string = "Slice: anexando a uma slice"
	SliceMake                    string = "Slice: make"
	SliceMultiDimensional        string = "Slice: multi dimensional"
	SliceSurpresaArraySubjacente string = "Slice: a surpresa do array subjacente"
	MapsIntroducao               string = "Maps: introdução"
	MapsRangeEDeletando          string = "Maps: range e deletando"
)
