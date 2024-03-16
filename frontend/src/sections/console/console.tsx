import { useState, useEffect, useRef } from "react";
import styles from "./console.module.css";
import { Execution, ExecutionI } from "@/types/execution";
import { MessageType } from "@/types/message";
import { VscClearAll } from "react-icons/vsc";

export default function Console({
  output,
  sendInput,
  clearConsole,
  shouldFocus,
  executions,
}: {
  executions: Execution[];
  shouldFocus: boolean;
}) {
  const [inputValue, setInputValue] = useState("");
  const [clearCount, setClearCount] = useState(0);
  const terminalRef = useRef();

  useEffect(() => {
    terminalRef.current.scrollTop = terminalRef.current.scrollHeight;

    if (shouldFocus && terminalRef.current) {
      terminalRef.current.querySelector("input").focus();
    }
  }, [executions, shouldFocus, clearCount]);

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
        clearWrapper();
        break;
      default:
        sendInput(trimmedInput + "\r");
    }

    setInputValue("");
  };

  const clearWrapper = () => {
    clearConsole();
    setClearCount((prev) => prev + 1);
  };

  const getExtraClass = (type: MessageType) => {
    if (type == MessageType.Done) {
      return `${styles.msg__type__Done}`;
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.console} ref={terminalRef}>
        <div className={styles.console__header}>
          <span>Console</span>
          <VscClearAll
            color="white"
            className={styles.console__clear}
            onClick={() => clearWrapper()}
          />
        </div>
        {executions.map((exec, i) => (
          <div className={styles.console__execution} key={i}>
            {exec.Messages.map((msg, j) => (
              <div className={styles.console__exec__msg} key={`${i}-${j}`}>
                <span className={`${getExtraClass(msg.type)}`}>
                  {msg.message}
                </span>
              </div>
            ))}
          </div>
        ))}
        <form onSubmit={handleSubmit}>
          <input
            className={styles.console__input}
            value={inputValue}
            onChange={handleInputChange}
            autoFocus={shouldFocus}
          />
        </form>
      </div>
    </div>
  );
}
