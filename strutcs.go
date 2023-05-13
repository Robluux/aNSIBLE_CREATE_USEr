
package alpha

import "net/http"

// Options are the options passed to create a client
type Options struct {
	Function   string `json:"function"`
	Symbol     string `json:"symbol"`
	OutputSize string `json:"outputsize"`
	DataType   string `json:"datatype"`
	APIKey     string `json:"apikey"`
}

// Client is the struct returned and used for API calls