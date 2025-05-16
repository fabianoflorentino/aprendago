package documentacao

// rootDir represents the root directory for documentation files within the project.
const (
	rootDir = "internal/documentacao"
)

const (
	introducao    string = "Introdução"
	goDoc         string = "go doc"
	commandGoDoc  string = "godoc"
	pkgGoDev      string = "https://pkg.go.dev/"
	escrevendoDoc string = "Escrevendo documentação"
)

const (
	flagIntroducaoDocumentacao string = "--introducao-documentacao"
	flagGoDoc                  string = "--go-doc"
	flagCommandGoDoc           string = "--godoc"
	flagPkgGoDev               string = "--pkg-go-dev"
	flagEscrevendoDocumentacao string = "--escrevendo-documentacao"
)

const (
	descIntroducaoDocumentacao string = "Introdução"
	descGoDoc                  string = "go doc"
	descCommandGoDoc           string = "godoc"
	descPkgGoDev               string = "https://pkg.go.dev/"
	descEscrevendoDocumentacao string = "Escrevendo documentação"
)

var (
	topics = []string{
		introducao,
		goDoc,
		commandGoDoc,
		pkgGoDev,
		escrevendoDoc,
	}
)
