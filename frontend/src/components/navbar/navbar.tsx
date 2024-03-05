"use client";

import styles from "./navbar.module.css";

export default function Navbar() {
  return (
    <div className={styles.navbar}>
      <div>NastyProgrammer </div>
      <div>Language </div>
      <div>Run</div>
      <div>Share</div>
    </div>
  );
}
