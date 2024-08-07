# Dynamic-JSON-Parser

## Overview
This project provides a dynamic JSON parser written in Go, utilizing the reflection package. The parser can handle JSON structures that are not known at compile-time, making it highly versatile and adaptable to various JSON data formats.

## Features
Dynamic Parsing: Parse JSON data without predefined structures.
Reflection Usage: Leverage Go's reflection capabilities for flexible data handling.
Error Handling: Robust error handling for invalid JSON data.
Extensible: Easily extend the parser for additional data types and formats.

## Getting Started
### Prerequisites
Go 1.16+
### Installation
Clone the repository to your local machine:

```sh
git clone https://github.com/your-username/dynamic-json-parser-go.git
cd dynamic-json-parser-go
```
### Usage
- Import the package in your Go project
- Use the dynamic parser to parse JSON data

### Examples
Here are some example usages of the dynamic JSON parser:

Example 1: Simple JSON
```
{
    "name": "John Doe",
    "age": 30
}
```
Example 2: Nested JSON
```
{
    "name": "John Doe",
    "age": 30,
    "address": {
        "city": "New York",
        "zip": "10001"
    }
}
```
## Contributing
We welcome contributions to enhance the functionality of this project! To contribute:
- Fork the repository.
- Create a new branch.
- Make your changes and commit them.
- Push to your branch and create a pull request.

  Thank You !! :)
