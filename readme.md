Quiz Game

A simple command line based quiz game, where the questions and answers are read from a CSV file. The game has a time limit, after which it will automatically terminate.
Features

    The game reads questions and answers from a CSV file, with the format of question,answer.
    The user can specify the CSV file and time limit using command line flags.
    The game will keep track of the number of correct answers, and display the final score at the end.

Usage

To run the quiz game, use the following command:

go

go run main.go [flags]

The available flags are:

    -csv: The name of the CSV file. The default value is problems.csv.
    -limit: The time limit for the quiz in seconds. The default value is 5.

Example

Here is an example of how to run the quiz game with a different CSV file and time limit:

go

go run main.go -csv=questions.csv -limit=10

This will run the quiz game with the questions from questions.csv file, and a time limit of 10 seconds.