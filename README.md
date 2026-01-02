# Expense Tracker CLI

Command-line tool for recording expenses in a local JSON file. Built in Go with Cobra.

## Features

- Add, list, update, and delete expenses
- Summaries for all expenses or a specific month
- Configurable storage file path

## Installation

Build from source:

```bash
go build -o expense-tracker .
```

## Usage

```bash
./expense-tracker add --description "Coffee" --amount 5
./expense-tracker list
./expense-tracker summary
./expense-tracker summary --month 3
./expense-tracker update --id 1 --amount 6 --description "Coffee and bagel"
./expense-tracker delete --id 1
```

Run `./expense-tracker --help` to see all commands and flags.

## Data Storage

By default, expenses are stored at the user config directory in:

```
expense-tracker/expenses.json
```

You can override the path with the `--file` flag:

```bash
./expense-tracker --file /path/to/expenses.json list
```

Amounts are stored as integers and printed with a `$` prefix.

## Data Format

The file is a JSON array of expense objects with these fields:

- `id` (integer)
- `amount` (integer)
- `description` (string)
- `date` (RFC3339 timestamp)

[Link to Project Idea](https://roadmap.sh/projects/expense-tracker)