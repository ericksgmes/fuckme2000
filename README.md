
# 💀 fuckme2000

A lovingly hand-crafted programming language built from scratch in Go.  
It features a lexer, parser, and a stack-based virtual machine.  
It does math. It prints things. It will scream at you if you mess up.

> Created out of spite, curiosity, and poor life decisions.

---

## ✨ Features

- `let` variable assignment
- Integer literals
- Variable references
- Addition (`+`)
- `print()` function
- Semicolon-terminated statements
- Basic error handling (with lovingly hand-crafted panics)

---

## 🧠 Syntax

### Variable Assignment

```f2000
let x = 5;
let y = x + 3;

Variables are stored in a key-value memory store.
You can assign literal values, or expressions.

---

### Printing Values

```f2000
print(x);
print(5);
print(x + 1);
```

Use `print()` to output values.  
Supports:
- Integer literals
- Variables
- Simple addition inside the print

---

## 🚫 What It Doesn't Do Yet (But Might Someday)

- `if` statements
- `while` loops
- Functions
- Strings (you wish)
- Types
- Any actual error recovery that isn't screaming

---

## 📦 File Example (`example.f2000`)

```f2000
let x = 10;
let y = x + 2;
print(y);
```

Output:

```
12
```

---

## 🧪 Running It

From the root of your project:

```bash
go run cmd/fuckme2000/main.go examples/example.f2000
```

Make sure Go is installed, and your soul is prepared.

---

## 🛠 Structure

```
fuckme2000/
├── cmd/
|   ├── fuckme2000/  # Entry point
├── internal/
│   ├── lexer/      # Tokenizer
│   ├── parser/     # Syntax parser
│   └── vm/         # Virtual machine
├── examples/       # Sample .f2000 files
└── README.md       # You're reading it
```

---

## 🙏 Credits

- Language design: pure emotional chaos
- Lexer: built with anger and string slicing
- VM: stack-based, because registers are scary
- Parser: very hand-rolled, very fragile

---

## 🗿 Philosophy

> If your language doesn’t panic at runtime,  
> is it really a language at all?

---
