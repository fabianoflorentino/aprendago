// Package aplicacoes provides a collection of constants and utilities
// for managing application-specific configurations, command-line flags,
// and descriptions. It serves as a central repository for defining
// application names, their functionalities, and associated metadata
// to streamline the development and execution of various Go applications.
package aplicacoes

// rootDir represents the relative path to the "aplicacoes" directory within the internal package.
// It is used as a base directory for various application-specific operations and configurations.
const (
	rootDir = "internal/aplicacoes"
)

// Constants for application names and their respective flags and descriptions.
// These constants are used to identify different applications and their functionalities.
const (
	DocumentacaoJSON string = "Documentação JSON"
	JSONMarshal      string = "JSON marshal (ordenação)"
	JSONUnmarshal    string = "JSON unmarshal (desornação)"
	InterfaceWriter  string = "A interface Writer"
	PacoteSort       string = "O pacote sort"
	CustomizandoSort string = "Customizando sort"
	Bcrypt           string = "bcrypt"
)

// Constants for command-line flags and their descriptions.
// These flags are used to control the behavior of the applications and provide additional options to the user.
const (
	flagDocumentacaoJSON                     string = "--documentacao-json"
	flagDocumentacaoJSONExampleJSONMarshal   string = "--documentacao-json --example --json-marshal"
	flagDocumentacaoJSONExampleJSONUnmarshal string = "--documentacao-json --example --json-unmarshal"
	flagDocumentacaoJSONExampleJSONEncoder   string = "--documentacao-json --example --json-encoder"
	flagJSONMarshal                          string = "--json-marshal"
	flagJSONUnmarshal                        string = "--json-unmarshal"
	flagInterfaceWriter                      string = "--interface-writer"
	flagPacoteSort                           string = "--pacote-sort"
	flagPacoteSortExample                    string = "--pacote-sort --example"
	flagCustomizandoSort                     string = "--customizando-sort"
	flagCustomizandoSortExample              string = "--customizando-sort --example"
	flagBcrypt                               string = "--bcrypt"
)

// Constants for descriptions of the applications and their functionalities.
// These descriptions provide additional context and information about what each application does.
const (
	descDocumentacaoJSON                     string = "Descreve como documentar um pacote em Go"
	descDocumentacaoJSONExampleJSONMarshal   string = "Exemplo de como ordenar um JSON"
	descDocumentacaoJSONExampleJSONUnmarshal string = "Exemplo de como desordenar um JSON"
	descDocumentacaoJSONExampleJSONEncoder   string = "Exemplo de como usar o encoder JSON"
	descJSONMarshal                          string = "Descreve o pacote json.Marshal"
	descJSONUnmarshal                        string = "Descreve o pacote json.Unmarshal"
	descInterfaceWriter                      string = "Descreve o que é a interface Writer"
	descPacoteSort                           string = "Descreve o pacote sort"
	descPacoteSortExample                    string = "Exemplo de como usar o pacote sort"
	descCustomizandoSort                     string = "Descreve como customizar o pacote sort"
	descCustomizandoSortExample              string = "Exemplo de como customizar o pacote sort"
	descBcrypt                               string = "Descreve o pacote bcrypt"
)
