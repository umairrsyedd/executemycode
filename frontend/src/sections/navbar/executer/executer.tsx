import styles from "./executer.module.css";

import RunLogo from "../../../public/Run_Button.svg";
import StopLogo from "../../../public/Stop_Button.svg";

import Image from "next/image";
import ExecuteButton from "./execute_button";
import { useState } from "react";
import StopButton from "./stop_button";
import LoadingButton from "./loading_button";
import { ProgramState } from "@/types/program";

export default function Executer({ programState, handleExecute, handleStop }) {
  const loading = programState === ProgramState.Loading;

  return (
    <div className={styles.runner_wrapper}>
      {loading && <LoadingButton />}

      {!loading && programState === ProgramState.Idle && (
        <ExecuteButton onClick={handleExecute} />
      )}

      {!loading && programState === ProgramState.Executing && (
        <StopButton onClick={handleStop} />
      )}
    </div>
  );
}
