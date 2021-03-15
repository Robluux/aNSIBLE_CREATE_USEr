# Alphavantage API Lib

### Examples

```
func main() {
	opts := alpha.Options{
			Function:   "TIME_SERIES_DAILY",
			Symbol:     val,
			OutputSize: "compact",

			APIKey: os.Getenv("API_TOKEN"),
		}
		client := alpha.NewClient(opts, &http.Client{})
		api, err := client.Get()
		if err != nil {
	