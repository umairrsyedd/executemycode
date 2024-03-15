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
import { useEffect, useState } from "react";
import ResizableContainer, {
  Orientation,
} from "@/components/resizable/resizable_container";
import { ThemeContext, Themes } from "@/context/theme";
import { useLocalStorage } from "@uidotdev/usehooks";
import { useCustomWebSocket } from "@/hooks/useWebSocket";
import { ProgramState } from "@/types/program";
import { SocketState } from "@/types/socket";
import StatusBar from "@/sections/statusbar/statusbar";

export default function Page() {
  const [currentTheme, setTheme] = useLocalStorage("theme", Themes.Dark);
  const [code, setCode] = useState(sampleCodeMap.get(DefaultLanguage));
  const [currentLanguage, setCurrentLanguage] = useState(DefaultLanguage);
  const [programState, setProgramState] = useState(ProgramState.Idle);
  const [consoleOutput, setConsoleOutput] = useState([]);
  const [socketState, setSocketState] = useState(SocketState.Connecting);

  const onOutput = (output) => {
    setConsoleOutput((prevOutput) => [...prevOutput, output]);
  };

  const onDone = (message) => {
    setConsoleOutput((prevOutput) => [...prevOutput, message]);
    setProgramState(ProgramState.Idle);
  };

  const onError = (error) => {
    setConsoleOutput((prevOutput) => [...prevOutput, error]);
  };

  const onConnected = (event) => {
    setSocketState(SocketState.Success);
  };

  const onReconnectStop = (event) => {
    setSocketState(SocketState.Failed);
  };

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
    setCode(sampleCodeMap.get(language));
  };

  const handleExecute = async () => {
    setProgramState(ProgramState.Loading);
    await sendCode(code, currentLanguage);
    setProgramState(ProgramState.Executing);
  };

  const handleStop = async () => {
    setProgramState(ProgramState.Loading);
    await sendClose();
    setProgramState(ProgramState.Idle);
  };

  const clearConsole = () => {
    setConsoleOutput([]);
  };

  const sendConsoleInput = (input) => {
    sendInput(input);
  };

  const handleThemeToggle = () => {
    setTheme((prevTheme) =>
      prevTheme === Themes.Dark ? Themes.Light : Themes.Dark
    );
  };

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
        <div className={styles.main_area}>
          <ResizableContainer
            orientation={Orientation.Horizontal}
            initialPercent={70}
            minSizePercent={30}
            maxSizePercent={75}
          >
            <Editor
              currentLanguage={currentLanguage}
              code={code}
              setCode={setCode}
            />
          </ResizableContainer>
          <div className={styles.main_area_right}>
            <ResizableContainer
              orientation={Orientation.Vertical}
              initialPercent={50}
              minSizePercent={20}
              maxSizePercent={90}
            >
              <Console
                programState={programState}
                output={consoleOutput}
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
