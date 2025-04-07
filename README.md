# 🐐 fuckme2000

A toy programming language, custom VM, and web-based IDE — built from scratch with Go, React, and unreasonable determination.  
Also: it's goat-themed now. You're welcome.

---

## ✨ Live Demo

- **IDE** → [frontend](https://fme2000.vercel.app/)  
- **Backend API** → [backend](https://fuckme2000-1.onrender.com)

---

## 🐐 Language Features

- `let` declarations
- Integer math with `+`
- `print()` for output
- Stack-based virtual machine execution
- Minimal syntax, real results

Try this in the IDE:

```f2000
let x = 2;
let y = x + 3;
print(y + 1);
```

---

## ⚙️ Stack Overview

| Layer | Tech |
|-------|------|
| Language | Go |
| Web API | Fiber |
| Frontend | React + Vite |
| Editor | Monaco |
| Styling | TailwindCSS |
| Containers | Docker |
| Hosting | Render (API) + Vercel (IDE) |

---

## 🔧 Local Development

```bash
# Backend
cd web/server
go run main.go

# Frontend
cd web/frontend
yarn install
yarn dev
```

Or use:

```bash
docker-compose up --build
```

---

## 📦 API Reference

**POST** `/run`

```json
{
  "code": "let x = 1; print(x + 2);"
}
```

Returns:
```
3
```

---

## 🧑‍💻 Author

Built by [@ericksgmes](https://github.com/ericksgmes)  
A fullstack learning project focused on compiler design, system architecture, and building ridiculous things with real tools.

---

## 🐐 The Official GOAT

```txt
             ,__
             (oo)____
             (__)    )\
                ||--|| *
       fuckme2000 — the GOAT language
```

---

## 📜 License

MIT — fork it, run it, rename it, summon it.

> 🐐 Not production-ready. Just goat-ready.
---
