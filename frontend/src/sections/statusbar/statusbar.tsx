import { getClassesForConnStatus, getIconForConnStatus } from "./utils";
import styles from "./statusbar.module.css";
import { SocketState } from "@/types/socket";
import { LuServer, LuServerCrash, LuServerOff } from "react-icons/lu";

export default function StatusBar({ connectionStatus }) {
  let additionalClass = getClassesForConnStatus(connectionStatus);

  return (
    <div className={styles.container}>
      <div className={`${styles.status_bar} ${additionalClass}`}>
        {getIconForConnStatus(connectionStatus)}
        <span>{connectionStatus}</span>
      </div>
    </div>
  );
}
