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
			fmt.Printf("High: %v\n", v.High)
			fmt.Printf("Low: %v\n", v.Low)
			fmt.Printf("Close: %v\n", v.Close)
			fmt.Printf("Open: %v\n", v.Open)
			fmt.Printf("Volume: %v\n\n", v.Volume)
		}