# 🚽 Skibidi Programming Language

A meme-powered, beginner-friendly programming language with a unique, playful syntax. Skibidi supports variables, arithmetic, logic, loops, functions, block scoping, built-ins, and more—all with a Skibidi twist.

---

## 🌐 Try Skibidi Online!

Check out our **interactive documentation and playground**:

👉 [skibidilang.netlify.app](https://skibidilang.netlify.app)

- Run Skibidi code in your browser
- Explore interactive docs and examples
- Perfect for learning, experimenting, and sharing

---

## Features
- Meme-inspired keywords and error messages
- Variables, assignment, and block scoping
- Arithmetic, comparison, and logical operators
- If/else, while, and for loops
- Functions (including recursion and higher-order)
- Built-in functions: `len`, `abs`, `str`
- Interactive REPL mode
- Beginner-friendly and fun!

---

## Installation

### Windows
1. **Double-click** on `install.bat` in the repo folder.
2. **Restart your terminal** (Command Prompt or PowerShell).
3. If `skibidi` is still not recognized:
   - **Manually create** a folder at `C:/skibidilang` (if it doesn't exist).
   - **Move** `skibidi.exe` into `C:/skibidilang`.
   - **Add** `C:/skibidilang` to your system PATH environment variable:
     - Open System Properties → Environment Variables → System variables → Path → Edit → New → `C:/skibidilang` → OK.
   - Restart your terminal again.

### Mac/Linux
1. Open a terminal and navigate to the repo folder:
   ```sh
   cd skibidi-language
   chmod +x install.sh
   ./install.sh
   ```
2. If `skibidi` is still not recognized, add the repo folder to your PATH or move the `skibidi` binary to `/usr/local/bin`:
   ```sh
   sudo mv skibidi /usr/local/bin
   ```

---

## Usage

### Run a Skibidi Program
- **Windows:**
  ```sh
  skibidi.exe run myfile.skibidi
  ```
- **Linux/Mac:**
  ```sh
  ./skibidi run myfile.skibidi
  ```

### Start Interactive Mode (REPL)
- **Windows:**
  ```sh
  skibidi.exe -i
  ```
- **Linux/Mac:**
  ```sh
  ./skibidi -i
  ```

---

## Language Syntax

### Variable Declaration & Assignment
```skibidi
skibidi x rizz 5 ohio
x rizz 10 ohio
```

### Print
```skibidi
gyatt "Hello, Skibidi!" ohio
gyatt x ohio
```

### If/Else
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

### Functions
```skibidi
sigma add(a, b) {
    alpha a + b ohio
}
skibidi result rizz beta add(10, 32) ohio
gyatt result ohio
```

### Input
```skibidi
gyatt "Enter your name:" ohio
skibidi name rizz input ohio
gyatt "Hello, " + name + "!" ohio
```

### Comments
```skibidi
bruh this is a comment
```

---

## Built-in Functions
| Function | Usage         | Description                  |
|----------|--------------|------------------------------|
| len      | len(s)       | Length of string `s`         |
| abs      | abs(x)       | Absolute value of number `x` |
| str      | str(x)       | Converts number `x` to string|

---

## Example Program

```skibidi
bruh 🚽 Skibidi Ultimate Feature Test 🚽

gyatt "Welcome to the Skibidi Ultimate Feature Test!" ohio

gyatt "What's your name, Skibidi enjoyer?" ohio
skibidi username rizz input ohio
gyatt "Nice to meet you, " + username + "!" ohio

skibidi a rizz 7 ohio
skibidi b rizz 3 ohio
gyatt "a = " + a + ", b = " + b ohio

skibidi sum rizz a + b ohio
gyatt "Sum: " + sum ohio

sigma factorial(n) {
    cap (n <= 1) {
        alpha 1 ohio
    }
    nocap {
        alpha n * beta factorial(n - 1) ohio
    }
}
skibidi num rizz 5 ohio
skibidi fact rizz beta factorial(num) ohio
gyatt "Factorial of " + num + " is " + fact ohio
```

---

## Interactive Mode (REPL)
- Start with `skibidi -i` or `skibidi.exe -i`.
- Enter statements, expressions, or multi-line blocks.
- Use `:help`, `:vars`, `:funcs`, or `exit` for REPL commands.

---

## Stay Skibidi!
- All keywords and errors are meme-inspired for max fun.
- Beginner-friendly, readable, and playful.
- Extend, remix, and meme on!

If you want syntax highlighting, see the [docs](doc.md) or ask for a VSCode extension! 