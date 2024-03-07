import styles from "./runner.module.css";

import RunLogo from "../../../../public/Run_Button.svg";
import Image from "next/image";

export default function RunButton({ onClick }) {
  return (
    <button className={styles.run_container} onClick={onClick}>
      <Image src={RunLogo} height={10} alt="Run" />
      <div className={styles.runner_text}>Run</div>
    </button>
  );
}
