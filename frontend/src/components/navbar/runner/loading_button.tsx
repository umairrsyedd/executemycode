import styles from "./runner.module.css";

export default function LoadingButton() {
  return (
    <button className={styles.loading_container}>
      <div className={styles.runner_text}>Loading</div>
    </button>
  );
}
