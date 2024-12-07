package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
)

type DayResults struct {
	Day       int
	Part1     Result
	Part2     Result
	TotalTime string
}

var results []DayResults

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Advent of Code 2024 Results</title>
    <style>
        body { font-family: monospace; padding: 20px; max-width: 500px; margin-left: auto; margin-right: auto; text-align: center; }
        table { border-collapse: collapse; width: 100%; }
        th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
        th { background-color: #f2f2f2; }
    </style>
</head>
<body>
    <h1>Advent of Code 2024 Results</h1>
    <table>
        <tr>
            <th>Day</th>
            <th>Part 1</th>
            <th>Part 1 Time</th>
            <th>Part 2</th>
            <th>Part 2 Time</th>
            <th>Total Time</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.Day}}</td>
            <td>{{.Part1.Value}}</td>
            <td>{{.Part1.Duration}}</td>
            <td>{{.Part2.Value}}</td>
            <td>{{.Part2.Duration}}</td>
            <td>{{.TotalTime}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`

func RunAllSolutions() []DayResults {
	results = []DayResults{}
	// Find and run all solutions
	for day := 1; day <= 25; day++ {
		dayResults, err := RunDay(day)
		if err == nil {
			results = append(results, dayResults)
		}

	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Day < results[j].Day
	})
	return results
}

func RunDay(day int) (DayResults, error) {
	solution, ok := GetSolution(day)
	if !ok {
		return DayResults{}, fmt.Errorf("no solution found for day %d", day)
	}

	input, err := ReadInputFile(day)
	if err != nil {
		return DayResults{}, fmt.Errorf("failed to read input for day %d: %v", day, err)
	}

	return RunAndStore(day, solution, input), nil
}

func StartServer(port int) {
	tmpl := template.Must(template.New("results").Parse(htmlTemplate))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		results := RunAllSolutions()
		tmpl.Execute(w, results)
	})

	fmt.Printf("Server starting on http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
