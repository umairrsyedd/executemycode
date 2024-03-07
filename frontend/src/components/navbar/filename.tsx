import styles from "./navbar.module.css";
export default function FileName() {
  return (
    <div contentEditable={true} className={styles.filename}>
      CrazyProgrammer
    </div>
  );
}
