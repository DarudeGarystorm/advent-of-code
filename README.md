# Advent of Code 2024

This repository contains solutions for the Advent of Code 2024 challenges.

## How to Run

To run the solutions for each day, use the `go run` command followed by the path to the main file. For example, to run the solution for day 1:

```sh
go run day01/main.go
```

Similarly, to run the solution for day 2:

```sh
go run day02/main.go
```

## How to Test

To run the tests, use the `go test ./...` command. This command will automatically find and run all the test files in the project that follow the naming convention `*_test.go`.

Navigate to the root directory of your project and run:

```sh
go test ./...
```

For more detailed output, use the `-v` flag:

```sh
go test -v ./...
```

This will provide verbose output, showing the details of each test that is run.

## Project Structure

- `day01/main.go`: Solution for day 1.
- `day02/main.go`: Solution for day 2.
- `utils/utils.go`: Contains utility functions used across different days.
- `utils/utils_test.go`: Contains tests for the utility functions (in the `utils_test` package).

## Dependencies

This project uses the Go standard library. Ensure you have Go installed on your machine. You can download and install Go from [the official website](https://golang.org/dl/).

## License

This project is licensed under the MIT License.

## Setting Up a New Day

To set up a new day, use the `go run init_day.go <year> <current_day>` command. For example, to set up day 3 for the year 2024:
```sh
go run init_day.go 2024 3
```
This will create the directory and boilerplate code for the next day (day 3 in this example). After running the command, download the input file for the new day and save it as `day03/input.txt`.