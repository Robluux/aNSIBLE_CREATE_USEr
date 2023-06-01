
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
type TimeSeriesDaily struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
	Date   string `json:"date,omitempty"`
}

// Meta contains meta data information for the call
type Meta struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

// ResponsePayload allows for more easily structured data from the API
type ResponsePayload struct {
	MetaData   Meta              `json:"Meta Data"`
	TimeSeries []TimeSeriesDaily `json:"Time Series (Daily)"`
}