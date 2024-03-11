import styles from "./executer.module.css";

import RunLogo from "../../../../public/Run_Button.svg";
import Image from "next/image";

export default function RunButton({ onClick }) {
  return (
    <button className={styles.run_container} onClick={onClick}>
      <Image src={RunLogo} height={10} alt="Execute" />
      <div className={styles.runner_text}>Execute</div>
    </button>
  );
}
