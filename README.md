
# SkibidiLang Installation Guide

## Windows

1. **Double-click** on the `install.bat` file.
2. **Restart** your terminal (Command Prompt or PowerShell).
3. If the terminal still doesn't recognize `skibidi` as a command:
   - Manually move `skibidi.exe` to a folder already in your system PATH (e.g., `C:\Windows` or any other folder added to your environment variables).
   - Or add the SkibidiLang folder to your system environment variables.

## macOS & Linux

1. Open a terminal.
2. Run the following command:

   ```bash
   chmod +x install.sh && ./install.sh
   ```

3. Restart your terminal.
4. If `skibidi` is still not recognized:
   - Manually move the `skibidi` binary to `/usr/local/bin` using:

     ```bash
     sudo mv skibidi /usr/local/bin
     ```

   - Make sure `/usr/local/bin` is in your `$PATH`.



# 📚 SkibidiLang Language Documentation

Welcome to the official documentation for **SkibidiLang**, a fun and expressive programming language created using Go. This guide explains the syntax, keywords, and behaviors in SkibidiLang.

---

## 🔤 Keywords

| Keyword  | Purpose                       |
| -------- | ----------------------------- |
| `rizz`   | Declare a variable            |
| `gyatt`  | Loop (`while`)                |
| `ohio`   | Print to console              |
| `bussin` | `else` block                  |
| `if`     | Start conditional block       |
| `true`   | Boolean true                  |
| `false`  | Boolean false                 |
| `null`   | Null literal                  |
| `and`    | Logical AND                   |
| `or`     | Logical OR                    |
| `not`    | Logical NOT                   |
| `fn`     | Function declaration (future) |

---

## 📄 Variables

### Declare a Variable

```skibidi
rizz name = "Shahid";
rizz age = 20;
rizz isCool = true;
```

---

## ➕ Operators

* Arithmetic: `+`, `-`, `*`, `/`
* Comparison: `==`, `!=`, `<`, `<=`, `>`, `>=`
* Logical: `and`, `or`, `not`

---

## 🔁 Loops

### While Loop

```skibidi
rizz x = 0;
gyatt (x < 5) {
  ohio x;
  x = x + 1;
}
```

---

## 🧠 Conditionals

```skibidi
rizz score = 85;

if (score > 90) {
  ohio "You got an A+";
} bussin {
  ohio "Try harder next time.";
}
```

---

## 🗣️ Output

```skibidi
ohio "Hello, World!";
ohio 3 + 4;
```

---

## 🤖 Boolean Logic

```skibidi
rizz isSkibidi = true;
rizz isSigma = false;

if (isSkibidi and not isSigma) {
  ohio "You are skibidi but not sigma.";
}
```

---

## 🧪 REPL Mode

Start interactive shell:

```bash
skibidi
```

You can type SkibidiLang code directly and get output instantly.

---

## 🔧 Future Work

* Function declarations
* Better error reporting
* File I/O support

---

## 🧑‍💻 Created by

Shahid Amin & Team

MIT License
