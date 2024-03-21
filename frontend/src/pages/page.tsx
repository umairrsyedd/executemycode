"use client";

import Navbar from "@/sections/navbar/navbar";
import Editor from "@/sections/editor/editor";
import Console from "@/sections/console/console";
import Notepad from "@/sections/notepad/notepad";

import styles from "./page.module.css";
import {
  DefaultLanguage,
  LanguageName,
  LocalStoragePrefix,
  sampleCodeMap,
} from "@/types/languages";
import { useEffect, useRef, useState } from "react";
import z, { Orientation } from "@/components/resizable/resizable_container";
import { ThemeContext, Themes } from "@/context/theme";
import { useLocalStorage } from "@uidotdev/usehooks";
import { useCustomWebSocket } from "@/hooks/useWebSocket";
import { ProgramState } from "@/types/program";
import { SocketState } from "@/types/socket";
import StatusBar from "@/sections/statusbar/statusbar";
import ExecutionManager from "@/types/execution";
import { Message } from "@/types/message";
import { useLocalCode } from "@/hooks/useLocalCode";
import ResizableContainer from "@/components/resizable/resizable_container";

export default function Page() {
  const [currentTheme, setTheme] = useLocalStorage("theme", Themes.Dark);
  const { getCode, saveCodeToLocalStorage, clearLocalStorageCode } =
    useLocalCode();
  const [code, setCode] = useState(getCode(DefaultLanguage));
  const [currentLanguage, setCurrentLanguage] = useState(DefaultLanguage);
  const [programState, setProgramState] = useState(ProgramState.Idle);
  const [socketState, setSocketState] = useState(SocketState.Connecting);
  const [execManager, setExecManager] = useState(new ExecutionManager());
  const consoleInputRef = useRef();

  const onOutput = (output: Message) => execManager.AddMessage(output);

  const onError = (error: Message) => execManager.AddMessage(error);

  const onDone = (message: Message) => {
    execManager.AddMessage(message);
    setProgramState(ProgramState.Idle);
    consoleInputRef.current.blur();
  };

  const onConnected = (event) => setSocketState(SocketState.Success);
  const onReconnectStop = (event) => setSocketState(SocketState.Failed);

  const { onMessage, sendCode, sendInput, sendClose } = useCustomWebSocket(
    process.env.NEXT_PUBLIC_EXECUTION_SERVER_URL,
    onOutput,
    onDone,
    onError,
    onConnected,
    onReconnectStop
  );

  const handleLanguageChange = (language: LanguageName) => {
    setCurrentLanguage(language);
    setCode(getCode(language));
  };

  const handleExecute = async () => {
    setProgramState(ProgramState.Loading);
    execManager.NewExecution();
    await sendCode(code, currentLanguage);
    setProgramState(ProgramState.Executing);
    consoleInputRef.current.focus();
  };

  const handleStop = async () => {
    setProgramState(ProgramState.Loading);
    await sendClose();
    setProgramState(ProgramState.Idle);
    consoleInputRef.current.blur();
  };

  const clearConsole = () => {
    execManager.Clear();
  };

  const sendConsoleInput = (input) => {
    sendInput(input);
  };

  const handleEditorCodeChange = (code) => {
    saveCodeToLocalStorage(currentLanguage, code);
    setCode(code);
  };

  const handleEditorCodeReset = () => {
    clearLocalStorageCode(currentLanguage);
    setCode(getCode(currentLanguage));
  };

  const handleThemeToggle = () => {
    setTheme((prevTheme) =>
      prevTheme === Themes.Dark ? Themes.Light : Themes.Dark
    );
  };

  const containerRef = useRef(null);

  const [editorContainerWidth, setEditorContainerWidth] = useState();
  const [consoleContainerHeight, setConsoleContainerHeight] = useState();

  useEffect(() => {
    setEditorContainerWidth(containerRef.current.offsetWidth);
    setConsoleContainerHeight(containerRef.current.offsetHeight);
  });

  return (
    <ThemeContext.Provider value={currentTheme}>
      <div className={styles.page} data-theme={currentTheme}>
        <div className={styles.nav_container}>
          <Navbar
            currentLangauge={currentLanguage}
            handleLanguageChange={handleLanguageChange}
            setTheme={handleThemeToggle}
            programState={programState}
            setProgramState={setProgramState}
            handleExecute={handleExecute}
            handleStop={handleStop}
            socketStatus={socketState}
          />
        </div>
        <div className={styles.main_area} ref={containerRef}>
          <ResizableContainer
            orientation={Orientation.Horizontal}
            preLoadPercent={75}
            initialPx={editorContainerWidth * 0.75}
            maxPx={editorContainerWidth * 0.9}
            minPx={editorContainerWidth * 0.4}
          >
            <Editor
              currentLanguage={currentLanguage}
              code={code}
              setCode={handleEditorCodeChange}
              handleEditorCodeReset={handleEditorCodeReset}
            />
          </ResizableContainer>
          <div className={styles.main_area_right}>
            <ResizableContainer
              orientation={Orientation.Vertical}
              preLoadPercent={60}
              initialPx={consoleContainerHeight * 0.6}
              maxPx={consoleContainerHeight * 0.8}
              minPx={consoleContainerHeight * 0.2}
            >
              <Console
                ref={consoleInputRef}
                programState={programState}
                executions={execManager.GetExecutions()}
                sendInput={sendConsoleInput}
                clearConsole={clearConsole}
              />
            </ResizableContainer>
            <Notepad />
          </div>
        </div>
        <StatusBar connectionStatus={socketState} />
      </div>
    </ThemeContext.Provider>
  );
}
