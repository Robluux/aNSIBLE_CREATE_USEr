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
			fmt.Printf("Error: %v", err)
			return
		}
		for _, v := range api.TimeSeries {
			fmt.Printf("Symbol: %v\n", val)
			fmt.Printf("Date: %v\n", v.Date)
			fmt