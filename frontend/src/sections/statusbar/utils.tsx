import { SocketState } from "@/types/socket";
import { LuServer, LuServerCrash, LuServerOff } from "react-icons/lu";

import styles from "./statusbar.module.css";

export const getClassesForConnStatus = (connectionStatus: any) => {
  let additionalClass = "";
  switch (connectionStatus) {
    case SocketState.Connecting:
      additionalClass = `${styles.shimmer} ${styles.status_bar__connecting}`;
      break;
    case SocketState.Failed:
      additionalClass = `${styles.status_bar__failed}`;
      break;
    case SocketState.Success:
      additionalClass = `${styles.status_bar__success}`;
      break;
    default:
      break;
  }
  return additionalClass;
};

export const getIconForConnStatus = (connectionStatus: any) => {
  let size = 15;
  switch (connectionStatus) {
    case SocketState.Connecting:
      return <LuServerCrash size={size} />;
    case SocketState.Success:
      return <LuServer size={size} />;
    case SocketState.Failed:
      return <LuServerOff size={size} />;
  }
};
