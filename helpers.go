
package alpha

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bradfitz/slice"
)

// NewClient returns a new Alpha API Client
func NewClient(opts Options, client *http.Client) Client {
	c := Client{}
	c.base = APIURL
	c.Options = opts
	c.Client = client
	return c
}

// Get gets API endpoint based on options passed
func (c Client) Get() (ResponsePayload, error) {
	var pload APIPayload
	rpload := ResponsePayload{}
	resp, err := c.Client.Get(fmt.Sprintf("%s?function=%s&symbol=%s&outputsize=%s&apikey=%s", APIURL, c.Options.Function, c.Options.Symbol, c.Options.OutputSize, c.Options.APIKey))
	if err != nil {
		return rpload, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return rpload, err
	}
	err = json.Unmarshal(b, &pload)
	if err != nil {