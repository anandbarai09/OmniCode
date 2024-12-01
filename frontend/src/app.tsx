import { useState } from 'preact/hooks'
import './app.css'

export function App() {
  const [theme, setTheme] = useState("light");

  return (
    <div className={`editor-container ${theme}`}>
      {/* Top Toolbar */}
      <header className="toolbar">
        <button onClick={() => alert("Run code!")}>Run</button>
        <button onClick={() => setTheme(theme === "light" ? "dark" : "light")}>
          Toggle Theme
        </button>
        <span className="status-bar">File: main.js | UTF-8</span>
      </header>

      {/* Sidebar */}
      <aside className="sidebar">
        <h2>File</h2>
        <ul>
          <li>src</li>
          <li>public</li>
        </ul>
      </aside>

      {/* Main Code Editor */}
      <main className="code-editor">
        <textarea placeholder="Start coding here..."></textarea>
      </main>

      {/* Bottom Panel */}
      <footer className="footer-panel">
        <p>Terminal - Ready</p>
      </footer>
    </div>
  );
}
