# Monkeylang Interpreter

## Overview

Monkeylang Interpreter is an interpreter for a JavaScript-like programming language written in [Go](https://golang.org/). It features a lexer, abstract syntax tree (AST), parser, evaluator, and a REPL (Read-Eval-Print Loop), supporting key language features such as higher-order functions, closures, and various data structures. This project is designed to follow the principles outlined in Thorsten Ball's book _"Writing an Interpreter in Go"_.

## Features

- **Lexer and Parser**: Converts code into a structured format for evaluation.
- **Abstract Syntax Tree (AST)**: Represents code structure.
- **Evaluator**: Executes code, supporting higher-order functions, closures, arrays, strings, and hashmaps.
- **REPL**: Interactive environment to write and test Monkeylang code.

## Run Locally

### With Docker

1. **Clone the repository**

   ```bash
   git clone https://github.com/ta-02/c-like-interpreter.git
   cd c-like-interpreter
   ```

2. **Build the Docker image**

   ```bash
   docker build -t monkeylang-interpreter .
   ```

3. **Run the Docker image**

   ```bash
   docker run --rm -it -v $(pwd):/app monkeylang-interpreter
   ```

   This will start the REPL inside the Docker container.

### Without Docker

1. **Clone the repository**

   ```bash
   git clone https://github.com/ta-02/c-like-interpreter.git
   cd c-like-interpreter
   ```

2. **Build the project**

   ```bash
   go build -o monkeylang main.go
   ```

3. **Run the REPL**

   ```bash
   ./monkeylang
   ```
