
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
type Client struct {
	Options Options
	Client  *http.Client
	base    string
}

// APIPayload is the response payload from Get
type APIPayload struct {
	MetaData   Meta                       `json:"Meta Data"`
	TimeSeries map[string]TimeSeriesDaily `json:"Time Series (Daily)"`
}

// TimeSeriesDaily is the payload in the TimeSeries map