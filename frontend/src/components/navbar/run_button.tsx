import styles from "./navbar.module.css";

import RunLogo from "../../../public/Run_Button.svg";
import Image from "next/image";

export default function RunButton() {
  return (
    <button className={styles.run_container}>
      <Image src={RunLogo} height={10} />
      <div className={styles.run_text}>Run</div>
    </button>
  );
}
