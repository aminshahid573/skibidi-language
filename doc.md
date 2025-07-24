# 🚽 Skibidi Programming Language Documentation

---

## Table of Contents
1. [Introduction](#introduction)
2. [Running Skibidi](#running-skibidi)
3. [Lexical Elements](#lexical-elements)
4. [Variables & Scope](#variables--scope)
5. [Data Types](#data-types)
6. [Operators](#operators)
7. [Statements](#statements)
8. [Functions](#functions)
9. [Built-in Functions](#built-in-functions)
10. [Interactive Mode (REPL)](#interactive-mode-repl)
11. [Error Handling](#error-handling)
12. [Examples](#examples)
13. [Meme Vibe & Philosophy](#meme-vibe--philosophy)

---

## 1. Introduction

**Skibidi** is a meme-powered, beginner-friendly programming language with a unique, playful syntax. It supports variables, arithmetic, logic, loops, functions, block scoping, and more—all with a Skibidi twist.

---

## 2. Running Skibidi

### Run a Skibidi Program
- **Windows:**
  ```
  skibidi.exe run myfile.skibidi
  ```
- **Linux/Mac:**
  ```
  ./skibidi run myfile.skibidi
  ```

### Start Interactive Mode (REPL)
- **Windows:**
  ```
  skibidi.exe -i
  ```
- **Linux/Mac:**
  ```
  ./skibidi -i
  ```

---

## 3. Lexical Elements

### Keywords
| Keyword   | Meaning/Usage                |
|-----------|-----------------------------|
| skibidi   | Variable declaration        |
| rizz      | Assignment operator         |
| cap       | If statement                |
| nocap     | Else statement              |
| bussin    | While loop                  |
| gyatt     | Print statement             |
| ohio      | End of statement            |
| sigma     | Function definition         |
| beta      | Function call               |
| alpha     | Return from function        |
| gyatfor   | For loop                    |
| input     | Read input from user        |
| true      | Boolean true                |
| false     | Boolean false               |
| bruh      | Single-line comment         |

### Identifiers
- Names for variables and functions.
- Start with a letter or underscore, followed by letters, digits, or underscores.

### Literals
- **Numbers:** `42`, `3.14`, `-7`
- **Strings:** `"hello world"`
- **Booleans:** `true`, `false`

### Comments
- Single-line comments start with `bruh` and continue to the end of the line.
  ```skibidi
  bruh this is a comment
  ```

---

## 4. Variables & Scope

### Declaration
```skibidi
skibidi x rizz 5 ohio
skibidi name rizz "skibidi" ohio
```

### Assignment
```skibidi
x rizz 10 ohio
name rizz "sigma" ohio
```

### Scope
- Variables declared inside `{ ... }` are local to that block (including function bodies and control structures).
- Variables declared outside are global.

---

## 5. Data Types

- **Numbers:** Floating-point (e.g., `42`, `3.14`, `-7`)
- **Strings:** Double-quoted, e.g., `"hello world"`
- **Booleans:** `true`, `false`

---

## 6. Operators

### Arithmetic
| Operator | Meaning         | Example         |
|----------|----------------|----------------|
| +        | Addition       | `a + b`        |
| -        | Subtraction    | `a - b`        |
| *        | Multiplication | `a * b`        |
| /        | Division       | `a / b`        |
| %        | Modulo         | `a % b`        |
| -        | Unary minus    | `-a`           |

### Comparison
| Operator | Meaning           | Example         |
|----------|-------------------|----------------|
| ==       | Equal             | `a == b`       |
| <        | Less than         | `a < b`        |
| >        | Greater than      | `a > b`        |
| <=       | Less or equal     | `a <= b`       |
| >=       | Greater or equal  | `a >= b`       |

### Logical
| Operator | Meaning           | Example         |
|----------|-------------------|----------------|
| &&       | Logical AND       | `a && b`       |
| \|\|     | Logical OR        | `a || b`       |

---

## 7. Statements

### Print Statement
```skibidi
gyatt "Hello, world!" ohio
gyatt x ohio
gyatt "Value: " + x ohio
```

### If/Else Statement
```skibidi
cap (x > 5) {
    gyatt "x is big" ohio
}
nocap {
    gyatt "x is small" ohio
}
```

### While Loop
```skibidi
bussin (x < 10) {
    gyatt x ohio
    x rizz x + 1 ohio
}
```

### For Loop
```skibidi
gyatfor (skibidi i rizz 0; i < 5; i rizz i + 1) {
    gyatt i ohio
}
```
- The for-loop header uses semicolons to separate initialization, condition, and post-expression.
- No `ohio` is needed inside the parentheses.

### Input Statement
```skibidi
gyatt "Enter your name:" ohio
skibidi name rizz input ohio
gyatt "Hello, " + name + "!" ohio
```
- `input` reads a line from the user as a string.

---

## 8. Functions

### Definition
```skibidi
sigma add(a, b) {
    alpha a + b ohio
}
```

### Calling
```skibidi
skibidi result rizz beta add(10, 32) ohio
gyatt result ohio
```
- Use `beta` as an expression to get the return value.

### Return
```skibidi
alpha value ohio
```
- Returns `value` from the function.

### Scope in Functions
- Function parameters and variables declared inside the function are local to that function.

---

## 9. Built-in Functions

| Function | Usage             | Description                        |
|----------|-------------------|------------------------------------|
| len      | `len(s)`          | Length of string `s`               |
| abs      | `abs(x)`          | Absolute value of number `x`       |
| str      | `str(x)`          | Converts number `x` to string      |

**Example:**
```skibidi
skibidi s rizz "skibidi" ohio
gyatt len(s) ohio
gyatt abs(-42) ohio
gyatt str(123.45) ohio
```

---

## 10. Interactive Mode (REPL)

### Starting the REPL
- Run `skibidi -i` (or `skibidi.exe -i` on Windows).

### Features
- **Single-line and multi-line input:** Enter statements or multi-line blocks (functions, loops, etc.).
- **Expression evaluation:** Enter expressions (like `2 + 2`) and see the result.
- **REPL commands:**
  - `:help` — Show help for REPL commands.
  - `:vars` — List all current variables and their values.
  - `:funcs` — List all defined functions.
  - `:exit` — Exit the interactive mode.
- **Error recovery:** Errors are printed, but the REPL keeps running.
- **Session state:** Variables and functions persist for the whole session.
- **Automatic `ohio`:** If you forget to end a statement with `ohio`, the REPL adds it for you (unless you’re typing a block).

### Example Session
```
🚽 Skibidi Interactive Mode v2.0 🚽
Type :help for commands. Type 'exit' or :exit to quit.
skibidi> skibidi x rizz 10 ohio
skibidi> x * 2
20
skibidi> sigma double(n) {
... alpha n * 2 ohio
... }
skibidi> beta double(7) ohio
14
skibidi> :vars
Variables:
  x = 10
skibidi> :funcs
Functions:
  double
skibidi> :exit
Goodbye! Stay sigma! 🗿
```

---

## 11. Error Handling
- Errors are reported with a Skibidi-style message and the line number.
- Common errors:
  - Undefined variable or function
  - Wrong number/type of arguments
  - Unexpected token (syntax error)
  - Division by zero
- The REPL does not exit on error; you can keep coding.

---

## 12. Examples

### Hello World
```skibidi
gyatt "Hello, Skibidi!" ohio
```

### Factorial Function
```skibidi
sigma factorial(n) {
    skibidi result rizz 1 ohio
    gyatfor (skibidi i rizz 1; i <= n; i rizz i + 1) {
        result rizz result * i ohio
    }
    alpha result ohio
}

skibidi num rizz 5 ohio
skibidi fact rizz beta factorial(num) ohio
gyatt "Factorial of " + num + " is " + fact ohio
```

### FizzBuzz
```skibidi
gyatfor (skibidi i rizz 1; i <= 20; i rizz i + 1) {
    cap ((i % 3 == 0) && (i % 5 == 0)) {
        gyatt "FizzBuzz" ohio
    }
    nocap {
        cap (i % 3 == 0) {
            gyatt "Fizz" ohio
        }
        nocap {
            cap (i % 5 == 0) {
                gyatt "Buzz" ohio
            }
            nocap {
                gyatt i ohio
            }
        }
    }
}
```

---

## 13. Meme Vibe & Philosophy
- **Meme-first:** All keywords and error messages are meme-inspired for maximum fun.
- **Beginner-friendly:** Simple, readable syntax.
- **Expressive:** Supports real programming constructs (functions, loops, scope, etc.).
- **Playful:** Encourages creativity and meme-coding.

---

# 🚽 Stay Skibidi!  
If you want to extend the language, add more memes, or need more examples, just ask! 