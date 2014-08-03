package main

import (
	"bitbucket.org/dennishedegaard/sudoku.go"
	"encoding/json"
	"fmt"
	"github.com/hoisie/web"
	"io/ioutil"
)

func index(ctx *web.Context) string {
	data, err := ioutil.ReadFile("static/index.html")
	if err != nil {
		ctx.Abort(500, "Unable to read index.html")
	}
	return string(data)
}

func solve(ctx *web.Context) string {
	// Get the data from the request.
	input := []byte(ctx.Params["board"])

	// Convert from JSON string to sudoku.Board object.
	board := sudoku.Board{}
	err := json.Unmarshal(input, &board)
	if err != nil {
		fmt.Println(err)
		ctx.Abort(
			500, fmt.Sprintf("Unable to parse board because: %s", err))
		return ""
	}
	fmt.Println("Input board:")
	fmt.Println(board.String())

	// Validate the board.
	_, err = board.IsValid()
	if err != nil {
		ctx.Abort(400, fmt.Sprintf("%s", err))
	}

	// Solve and return the result.
	board = board.Solve()
	fmt.Println("Output board:")
	fmt.Println(board.String())

	// Convert to JSON and return the result.
	output, err := json.Marshal(board)
	if err != nil {
		ctx.Abort(
			500, fmt.Sprintf("Unable to marshal board because: %s", err))
	}
	// return board.String()
	ctx.ContentType("application/json")
	return string(output)
}

func main() {
	web.Get("/", index)
	web.Post("/solve", solve)
	web.Run("0.0.0.0:9999")
}
