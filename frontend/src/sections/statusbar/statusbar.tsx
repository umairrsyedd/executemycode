import { getClassesForConnStatus, getIconForConnStatus } from "./utils";
import styles from "./statusbar.module.css";
import { SocketState } from "@/types/socket";
import { LuServer, LuServerCrash, LuServerOff } from "react-icons/lu";
import { FaHeart } from "react-icons/fa";
import RCLogo from "../../../public/rc_logo.svg";
import Image from "next/image";
import Link from "next/link";

import { FaGithub } from "react-icons/fa";
import { FaLinkedin } from "react-icons/fa";

import { GITHUB_LINK, LINKEDIN_LINK, RC_SCOUT_LINK } from "./links";

export default function StatusBar({ connectionStatus }) {
  let additionalClass = getClassesForConnStatus(connectionStatus);

  return (
    <div className={styles.container}>
      <span className={styles.socials}>
        <Link href={GITHUB_LINK} target="_blank" rel="noopener">
          <FaGithub size={20} />
        </Link>

        <Link href={LINKEDIN_LINK} target="_blank" rel="noopener">
          <FaLinkedin size={20} />
        </Link>
      </span>

      <div className={styles.left_content}>
        Made with
        <span className={styles.heart_logo}>
          {<FaHeart color="red" size={10} />}
        </span>
        by Umair at
        <span className={styles.rc_logo}>
          <Link href={RC_SCOUT_LINK} target="_blank" rel="noopener">
            <Image src={RCLogo} height={20} alt="Recurse Center Logo" />
          </Link>
        </span>
      </div>

      <div className={`${styles.status_bar} ${additionalClass}`}>
        {getIconForConnStatus(connectionStatus)}
        <span>{connectionStatus}</span>
      </div>
    </div>
  );
}
