import { useState, useEffect, useRef } from "react";
import styles from "./console.module.css";

export default function Console({ output, sendInput, clearConsole }) {
  const [inputValue, setInputValue] = useState("");
  const terminalRef = useRef();

  useEffect(() => {
    terminalRef.current.scrollTop = terminalRef.current.scrollHeight;
  }, [output]);

  const handleInputChange = (e) => {
    setInputValue(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const trimmedInput = inputValue.trim();
    const trimmedInputLowerCase = trimmedInput.toLowerCase();
    if (trimmedInput === "") {
      return;
    }

    switch (trimmedInputLowerCase) {
      case "clear":
        clearConsole();
        break;
      default:
        sendInput(trimmedInput + "\r");
    }

    setInputValue("");
  };

  return (
    <div className={styles.container}>
      <div className={styles.console} ref={terminalRef}>
        {output.map((line, index) => (
          <div key={index}>{line}</div>
        ))}
        <form onSubmit={handleSubmit}>
          <input
            className={styles.console__input}
            value={inputValue}
            onChange={handleInputChange}
          />
        </form>
      </div>
    </div>
  );
}
