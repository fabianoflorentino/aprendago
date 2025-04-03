// Package concorrencia provides utilities and constants for exploring concurrency-related
// topics in Go. It includes definitions for topic names, flags, and descriptions
// that can be used to structure and reference concurrency concepts consistently
// throughout the project. This package is designed to facilitate learning and
// experimentation with concurrency mechanisms such as goroutines, waitgroups, mutexes,
// and atomic operations.
package concorrencia

// rootDir represents the relative path to the directory where concurrency-related
// topics are stored within the project. This constant is used to reference the
// directory in a consistent manner throughout the codebase.
const (
	rootDir = "internal/concorrencia"
)

// The following constants define the names of various concurrency-related topics.
// These constants are used to identify and reference specific sections of the
// concurrency content.
const (
	concorrenciaVsParalelismo  string = "Concorrência vs Paralelismo"
	goroutinesWaitgroups       string = "Goroutines & WaitGroups"
	discussaoCondicaoDeCorrida string = "Discussão: Condição de corrida"
	condicaoDeCorrida          string = "Condição de corrida"
	mutex                      string = "Mutex"
	atomic                     string = "Atomic"
)

// The following constants define the command-line flags that can be used to
// execute specific concurrency-related topics. These flags are used to
// provide a user-friendly interface for selecting and executing different
// sections of the concurrency content.
const (
	flagConcorrenciaVsParalelismo  string = "--concorrencia-vs-paralelismo"
	flagGoroutinesWaitgroups       string = "--goroutines-waitgroups"
	flagDiscussaoCondicaoDeCorrida string = "--discussao-condicao-de-corrida"
	flagCondicaoDeCorrida          string = "--condicao-de-corrida"
	flagMutex                      string = "--mutex"
	flagAtomic                     string = "--atomic"
)

// The following constants define the descriptions for each concurrency-related
// topic. These descriptions provide additional context and information about
// the content covered in each section. They are used to enhance the user
// experience by providing clear explanations of what each topic entails.
const (
	descConcorrenciaVsParalelismo  string = "Apresenta a diferença entre concorrência e paralelismo."
	descGoroutinesWaitgroups       string = "Apresenta o uso de goroutines e waitgroups."
	descDiscussaoCondicaoDeCorrida string = "Apresenta uma discussão sobre condição de corrida."
	descCondicaoDeCorrida          string = "Apresenta o conceito de condição de corrida."
	descMutex                      string = "Apresenta o uso de mutex."
	descAtomic                     string = "Apresenta o uso de atomic."
)

// The following slice contains the names of all concurrency-related topics.
// This slice is used to iterate over the topics and execute the corresponding
// sections when the user selects a specific topic.
var (
	topics = []string{
		concorrenciaVsParalelismo,
		goroutinesWaitgroups,
		discussaoCondicaoDeCorrida,
		condicaoDeCorrida,
		mutex,
		atomic,
	}
)
