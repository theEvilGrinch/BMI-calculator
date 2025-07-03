# BMI Calculator

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)](LICENSE)

[![Support me on Boosty](https://img.shields.io/badge/Boosty-Support%20me-%23f15f2c?style=for-the-badge)](https://boosty.to/theEvilGrinch/donate)
[![Donate](https://img.shields.io/badge/Donate-%23702ff4?style=for-the-badge)](https://yoomoney.ru/to/410016288289737)

This is a simple command-line Body Mass Index (BMI) calculator written in Go. This project serves as a "Hello, World!" to the Go programming language, demonstrating basic concepts like variable declaration, user input, mathematical calculations, and conditional logic with `switch`.

## Table of Contents

- [How It Works](#how-it-works)
- [Prerequisites](#prerequisites)
- [How to Run](#how-to-run)
  - [Option 1: Using `go run`](#option-1-using-go-run)
  - [Option 2: Building the Executable](#option-2-building-the-executable)
  - [Example](#example)
- [Contributing](#contributing)
- [License](#license)

## How It Works

The program performs the following steps:
1.  Prompts the user to enter their height in meters.
2.  Prompts the user to enter their weight in kilograms.
3.  Validates user input to ensure it is a positive number, handling potential errors from non-numeric entries or non-positive values.
4.  Calculates the BMI using the standard formula.
5.  Rounds the result to the nearest whole number.
6.  Displays the calculated BMI along with a corresponding weight status category based on the World Health Organization (WHO) classification.

## Prerequisites

- [Go](https://go.dev/doc/install) (version 1.21 or later) installed on your system.
- [Git](https://git-scm.com/downloads) installed on your system.

## How to Run

Clone the repository:
```bash
  git clone https://github.com/theEvilGrinch/BMI-calculator.git
 ```
There are two ways to run the application:

### Option 1: Using `go run`

This command compiles and runs the program without creating a permanent executable.

1. Navigate to the project directory:
    ```bash
    cd BMI-calculator
    ```
2. Run the application:
    ```bash
    go run main.go
    ```

### Option 2: Building the Executable

This command compiles the source code into a single executable file that you can run directly.

1. Build the executable:
    ```bash
    go build
    ```
    This will create an executable file named `BMI-calculator` (on Linux/macOS) or `BMI-calculator.exe` (on Windows).
2. Run the executable:

   Linux/macOS:
    ```bash
    ./BMI-calculator
    ```
   Windows:
    ```bash
      BMI-calculator.exe
    ```

### Example

```
	Body Mass Index Calculator
Enter your height in meters (e.g., 1.8): 1.8
Enter your weight in kilograms: 75
Your Body Mass Index: 23 (normal weight)
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Maintained by [@theEvilGrinch](https://github.com/theEvilGrinch)
