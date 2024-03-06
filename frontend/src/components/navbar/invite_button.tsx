import styles from "./navbar.module.css";

import InviteIcon from "../../../public/invite_icon.svg";
import Image from "next/image";

export default function InviteButton() {
  return (
    <button className={styles.invite_container}>
      <Image src={InviteIcon} height={15} />
      <div className={styles.invite_text}>Invite</div>
    </button>
  );
}
