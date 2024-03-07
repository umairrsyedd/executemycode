import styles from "./runner.module.css";

import StopLogo from "../../../../public/Stop_Button.svg";
import Image from "next/image";

export default function StopButton({ onClick }) {
  return (
    <button className={styles.stop_container} onClick={onClick}>
      <Image src={StopLogo} height={10} alt="Stop" />
      <div className={styles.runner_text}>Stop</div>
    </button>
  );
}
