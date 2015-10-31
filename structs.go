package wahanthok

type Meaning struct {
	L1 string	// Meaning in first language
	L2 string	// Meaning in second language
}

// Wa denotes a word in a language with its meanings in another language.
type Wa struct {
	Word string
	Meanings []Meaning
}

// A collection of meanings
type Was map[string]Wa

// Represent the search page for templating
type SearchPage struct {
	BackgroundImage string
	W Wa
	Msg string	// not found or error message
}