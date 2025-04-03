package canais

// rootDir is a constant that holds the relative path to the "canais" directory
// within the "internal" directory. This path is used as a base directory for
// various operations related to the "canais" module in the project.
const (
	rootDir = "internal/canais"
)

// entendendoCanais represents the topic "Entendendo Canais",
// which introduces the concept of channels in Go, explaining
// their purpose and how they facilitate communication between
// goroutines.
const (
	entendendoCanais                  string = "Entendendo Canais"
	canaisDirecionaisUtilizandoCanais string = "Canais direcionais & utilizando canais"
	rangeEClose                       string = "Range e Close"
	selectStatement                   string = "Select"
	commaOkExpression                 string = "A expressão comma ok"
	convergencia                      string = "Convergência"
	divergencia                       string = "Divergência"
	context                           string = "Context"
)

// flag constants are used to define command-line flags for various topics
// related to channels in Go. Each flag corresponds to a specific topic,
// allowing users to access information or examples related to that topic
// when running the program.
const (
	flagEntendendoCanais                  string = "--entendendo-canais"
	flagCanaisDirecionaisUtilizandoCanais string = "--canais-direcionais-utilizando-canais"
	flagRangeEClose                       string = "--range-e-close"
	flagSelectStatement                   string = "--select"
	flagCommaOkExpression                 string = "--a-expressao-comma-ok"
	flagConvergencia                      string = "--convergencia"
	flagConvergenciaExample               string = "--convergencia --example"
	flagConvergenciaExampleChanString     string = "--convergencia --example --chan-string"
	flagDivergencia                       string = "--divergencia"
	flagDivergenciaExample                string = "--divergencia --example"
	flagDivergenciaExampleWithFunc        string = "--divergencia --example --with-func"
	flagContext                           string = "--context"
)

// description constants provide a brief explanation of each topic related to channels.
// These descriptions are used to inform users about the content and purpose of each topic
// when they access the program's help or menu options.
const (
	descEntendendoCanais                  string = "Entendendo Canais"
	descCanaisDirecionaisUtilizandoCanais string = "Canais direcionais & utilizando canais"
	descRangeEClose                       string = "Range e Close"
	descSelectStatement                   string = "Select"
	descCommaOkExpression                 string = "A expressão comma ok"
	descConvergencia                      string = "Convergência"
	descConvergenciaExample               string = "Convergência - Exemplo"
	descConvergenciaExampleChanString     string = "Convergência de Strings - Exemplo"
	descDivergencia                       string = "Divergência"
	descDivergenciaExample                string = "Divergência - Exemplo"
	descDivergenciaExampleWithFunc        string = "Divergência - Exemplo com Função"
	descContext                           string = "Context"
)

var (
	topics = []string{
		entendendoCanais,
		canaisDirecionaisUtilizandoCanais,
		rangeEClose,
		selectStatement,
		commaOkExpression,
		convergencia,
		divergencia,
		context,
	}
)
