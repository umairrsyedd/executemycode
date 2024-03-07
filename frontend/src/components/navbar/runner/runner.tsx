import styles from "./runner.module.css";

import RunLogo from "../../../public/Run_Button.svg";
import StopLogo from "../../../public/Stop_Button.svg";

import Image from "next/image";
import RunButton from "./run_button";
import { useState } from "react";
import StopButton from "./stop_button";
import LoadingButton from "./loading_button";

const enum ProgramState {
  Running = "Running",
  Idle = "Idle",
  Stopping = "Stopping",
}

export default function Runner() {
  const [programState, setProgramState] = useState(ProgramState.Idle);
  const [loading, setLoading] = useState(false);

  const handleRunClick = () => {
    setLoading(true);
    setProgramState(ProgramState.Running);

    setTimeout(() => {
      setLoading(false);
    }, 500);
  };

  const handleStopClick = () => {
    setLoading(true);
    setProgramState(ProgramState.Idle);

    setTimeout(() => {
      setLoading(false);
    }, 500);
  };

  return (
    <>
      {loading && <LoadingButton />}

      {!loading && programState === ProgramState.Idle && (
        <RunButton onClick={handleRunClick} />
      )}

      {!loading && programState === ProgramState.Running && (
        <StopButton onClick={handleStopClick} />
      )}
    </>
  );
}
