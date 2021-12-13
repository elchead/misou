import { useEffect, useState } from "react";

function Loading() {
  const [i, setI] = useState(0);

  useEffect(() => {
    const timeoutId = setTimeout(() => {
        if (i >= 3) {
          setI(0);
        } else {
          setI(i + 1);
        }
      }, 400);
      return function cleanup() {
          clearTimeout(timeoutId)
      }
  }, [i])

  return (
    <div className="transition-all duration-200 mb-6 dark:text-gray-500 text-4xl font-semibold flex justify-center">
      <div className="w-44">
        <span>Loading</span>
        {i >= 1 ? <span>.</span> : null}
        {i >= 2 ? <span>.</span> : null}
        {i >= 3 ? <span>.</span> : null}
      </div>
    </div>
  );
}

export default Loading;
