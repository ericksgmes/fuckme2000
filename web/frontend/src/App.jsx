import { useState } from "react";
import { useEffect } from "react";
import Editor from "@monaco-editor/react";
import { useMonaco } from "@monaco-editor/react";


function App() {
  const [code, setCode] = useState("let x = 3;\nlet y = x + 4;\nprint(y);");
  const [output, setOutput] = useState("");

  const runCode = async () => {
    const response = await fetch("http://localhost:3001/run", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ code }),
    });

    const text = await response.text();
    setOutput(text);
  };

  const monaco = useMonaco();

  useEffect(() => {
  if (monaco) {
    monaco.languages.register({ id: "f2000" });

    monaco.languages.setMonarchTokensProvider("f2000", {
      keywords: ["let", "print"],
      operators: ["+", "=", ";"],
      tokenizer: {
        root: [
          [/[a-zA-Z_]\w*/, {
            cases: {
              "@keywords": "keyword",
              "@default": "identifier",
            },
          }],
          [/[0-9]+/, "number"],
          [/[=+;]/, "operator"],
          [/[()]/, "delimiter"],
          [/\s+/, "white"],
        ],
      },
    });

    monaco.editor.defineTheme("f2000-dark", {
      base: "vs-dark",
      inherit: true,
      rules: [
        { token: "keyword", foreground: "ff79c6" },
        { token: "number", foreground: "bd93f9" },
        { token: "identifier", foreground: "f8f8f2" },
        { token: "operator", foreground: "8be9fd" },
        { token: "delimiter", foreground: "50fa7b" },
      ],
      colors: {},
    });
  }
}, [monaco]);

  return (
    <div className="flex flex-col h-screen bg-zinc-900 text-white p-4">
      <h1 className="text-2xl font-bold mb-4">fuckme2000 Web IDE</h1>
  
      <div className="flex flex-1 overflow-hidden gap-4">
        {/* Left side: Editor */}
        <div className="flex-1 rounded border border-zinc-700 overflow-hidden min-w-0">
        <Editor
          height="100%"
          language="f2000"
          theme="f2000-dark"
          value={code}
          onChange={(val) => setCode(val)}
        />
        </div>
  
        {/* Right side: Output */}
        <div className="w-1/3 bg-zinc-800 p-4 rounded overflow-auto">
          <h2 className="text-lg font-semibold mb-2">Output</h2>
          <pre className="whitespace-pre-wrap text-green-400">{output}</pre>
        </div>
      </div>
  
      {/* Run Button */}
      <div className="mt-4">
        <button
          onClick={runCode}
          className="bg-pink-600 hover:bg-pink-700 transition px-4 py-2 rounded text-white"
        >
          Run
        </button>
      </div>
    </div>
  );  
}

export default App;
