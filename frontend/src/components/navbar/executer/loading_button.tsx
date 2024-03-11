import styles from "./executer.module.css";

export default function LoadingButton() {
  return (
    <button className={styles.loading_container}>
      <div className={styles.lds_ring}>
        <div></div>
        <div></div>
        <div></div>
      </div>
      <div className={styles.runner_text}>Working</div>
    </button>
  );
}
