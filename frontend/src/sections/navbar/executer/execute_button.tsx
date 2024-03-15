import styles from "./executer.module.css";

import RunLogo from "../../../../public/Run_Button.svg";
import Image from "next/image";
import { SocketState } from "@/types/socket";

export default function ExecuteButton({ onClick, socketStatus }) {
  return (
    <button
      className={styles.run_container}
      onClick={onClick}
      disabled={
        socketStatus === SocketState.Connecting ||
        socketStatus === SocketState.Failed
      }
    >
      <Image src={RunLogo} height={10} alt="Execute" />
      <div className={styles.runner_text}>Execute</div>
    </button>
  );
}
