
package alpha

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bradfitz/slice"
)

// NewClient returns a new Alpha API Client