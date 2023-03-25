
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