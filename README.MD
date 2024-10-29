# Guess-it-2

## Description

**Guess-it-2** is a Go program that reads a series of floating-point numbers from user input and uses linear regression to predict the next value based on the input data. The program calculates the coefficients for the linear regression model and provides a range around the predicted value, allowing users to see potential future values.

## Features

- Accepts a series of floating-point numbers,integers  as input.
- Computes linear regression coefficients (slope and intercept).
- Predicts the next value based on the input data.
- The program calculates the last five values if the data set is huge enough
- The program has a guess margin of 20.0
- Provides a range around the predicted value with adjustable margins.

## Prerequisites

- Go programming language installed (version 1.15 or later).

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/nyagooh/guess-it-2.git
   cd guess-it-2
## Usage

To execute the program:
```bash
./script.sh
```
Input your data in the standard input:
```bash
189
113
```
Expected output:
```bash
282 322
```
## Error Handling
The program runs even in a scenario where the data set has negatives.
```bash
-189
113

```
The expected output will look something like:
```bash
-96 -56
```
The program also hanmdles edge cases where the data set has decimal values
```bash
189.8
113.6

```
The expected output will look something like this:
```bash
283 323
```
The program does not run when you input is only one.The program waits for you to input another data for it to run.
```bash
132
```
## Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.
## Acknowledgments
  - Go Programming Language
  - Random Number Generation Techniques

## [License](LICENSE)
