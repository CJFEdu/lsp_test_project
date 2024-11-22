# LSP Test Project

This is a project created for testing LSP servers.  Specifically, for testing the Multilspy LSP client library found at [Multilspy](https://github.com/microsoft/multilspy).  The idea was to create a simple project that could be replicated in multiple languages and used as a reference for testing the LSP client library.

# Quiz Generator

A flexible command-line quiz application written in Go that supports multiple quiz types and configurations.

## Features

- Multiple question types:
  - Multiple Choice
  - Fill in the Blank
  - True/False
- JSON-based quiz configuration
- Randomizable question order
- Time-limited quizzes
- Interactive menu system
- Random quote generator with programming humor

## Project Structure

```
.
├── go/
│   ├── main.go      # Main program entry point
│   ├── config.go    # Configuration handling
│   ├── menu.go      # Menu system
│   ├── quotes.go    # Quote generator
│   └── quiz.go      # Quiz logic
├── quiz/            # Quiz JSON files
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

1. Clone the repository
2. Navigate to the `go` directory
3. Run the program:
   ```bash
   go run .
   ```

## Usage

The program provides an interactive menu with the following options:

1. List Available Quizzes
2. Start a Quiz
3. Exit


## Quiz Format

Quizzes are stored in JSON format in the `quiz` directory. Each quiz consists of:

- `config.json`: Quiz configuration file
- Question files: Individual JSON files for each question

### Quiz Configuration

Example `config.json`:
```json
{
  "title": "Sample Quiz",
  "timelimit": 300,
  "randomize": true,
  "passingScore": 70
}
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
