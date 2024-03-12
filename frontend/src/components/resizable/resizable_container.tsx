import { useRef, useState } from "react";

import styles from "./resizable.module.css";
import _ from "lodash";

export const enum Orientation {
  Horizontal = "X",
  Vertical = "Y",
}

// TODO: 1. Maybe Use Pixels instead of Percent.
// TODO: 2. Try to Rely on Vanilla Dom Manupulation.

export default function ResizableContainer({
  orientation,
  initialPercent,
  minSizePercent,
  maxSizePercent,
  throttleResize = 16,
  children,
}: {
  orientation: Orientation;
  initialPercent: number;
  minSizePercent: number;
  maxSizePercent: number;
  throttleResize: number;
  children: React.ReactNode;
}) {
  const containerRef = useRef(null);
  const [currentPercent, setCurrentPercent] = useState(initialPercent);

  let initialResizerPosition = null;
  const handleResizeStart = (event) => {
    initialResizerPosition =
      orientation === Orientation.Horizontal ? event.clientX : event.clientY;
    document.addEventListener("mousemove", handleResize);
    document.addEventListener("mouseup", handleResizeEnd);
  };

  const handleResize = (event) => {
    handleResizeThrottled(event);
  };

  const handleResizeThrottled = _.throttle((event) => {
    const currentResizer =
      orientation === Orientation.Horizontal ? event.clientX : event.clientY;

    const delta = currentResizer - initialResizerPosition;

    let newPercent = 0;
    if (orientation === Orientation.Horizontal) {
      const containerWidth = containerRef.current.offsetWidth;
      newPercent = ((containerWidth + delta) / containerWidth) * 100;
    } else {
      const containerHeight = containerRef.current.offsetHeight;
      newPercent = ((containerHeight + delta) / containerHeight) * 100;
    }

    newPercent = Math.max(minSizePercent, Math.min(maxSizePercent, newPercent));

    requestAnimationFrame(() => {
      setCurrentPercent(newPercent);
    });
  }, throttleResize);

  const handleResizeEnd = () => {
    document.removeEventListener("mousemove", handleResize);
    document.removeEventListener("mouseup", handleResizeEnd);
  };

  const containerStyles = {
    display: `flex`,
    width:
      orientation === Orientation.Horizontal ? `${currentPercent}%` : undefined,
    height:
      orientation === Orientation.Vertical ? `${currentPercent}%` : undefined,
    flexDirection: orientation === Orientation.Horizontal ? `row` : `column`,
  };

  return (
    <div ref={containerRef} style={containerStyles}>
      {children}
      <div
        className={
          orientation === Orientation.Horizontal
            ? styles.resizer_horizontal
            : styles.resizer_vertical
        }
        onMouseDown={handleResizeStart}
      ></div>
    </div>
  );
}
