package store

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Claim struct {
	ID        uint   `json:"id" `
	ModelName string `json:"model_name"`
	Amount    int    `json:"amount"`
}

func GenerateBarGraphClaims(claims []Claim, models []string) {
	// Implementation for generating bar graph from claims data
	// This is a placeholder function. Actual implementation would depend on the graphing library used.

	// Count claims per type
	counts := make(map[string]int)
	for _, c := range claims {
		counts[c.ModelName]++
	}

	// Prepare chart data
	var xAxis []string
	var yAxis []opts.BarData
	for _, m := range models {
		xAxis = append(xAxis, m)
		yAxis = append(yAxis, opts.BarData{Value: counts[m]})
	}

	// Create bar chart
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Claims by Models",
			Subtitle: "Generated using faker & random",
		}),
	)
	bar.SetXAxis(xAxis).AddSeries("Claims", yAxis)

	// Save to HTML
	f, _ := os.Create("claims.html")
	defer f.Close()
	bar.Render(f)

}
