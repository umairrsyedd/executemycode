import styles from "./executer.module.css";

import RunLogo from "../../../public/Run_Button.svg";
import StopLogo from "../../../public/Stop_Button.svg";

import Image from "next/image";
import RunButton from "./execute_button";
import { useState } from "react";
import StopButton from "./stop_button";
import LoadingButton from "./loading_button";

const enum ProgramState {
  Executing = "Executing",
  Idle = "Idle",
  Stopping = "Stopping",
}

export default function Executer() {
  const [programState, setProgramState] = useState(ProgramState.Idle);
  const [loading, setLoading] = useState(false);

  const handleRunClick = () => {
    setLoading(true);
    setProgramState(ProgramState.Executing);

    setTimeout(() => {
      setLoading(false);
    }, 1000);
  };

  const handleStopClick = () => {
    setLoading(true);
    setProgramState(ProgramState.Idle);

    setTimeout(() => {
      setLoading(false);
    }, 500);
  };

  return (
    <div className={styles.runner_wrapper}>
      {loading && <LoadingButton />}

      {!loading && programState === ProgramState.Idle && (
        <RunButton onClick={handleRunClick} />
      )}

      {!loading && programState === ProgramState.Executing && (
        <StopButton onClick={handleStopClick} />
      )}
    </div>
  );
}
