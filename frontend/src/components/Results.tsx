import { useState, useEffect } from "react";
import { useAppDispatch, useAppSelector } from "../app/hooks";
import Result from "./Result";
import {
  updateResultForPreview,
  incrementKey,
  decrementKey,
  hidePreview,
} from "../features/previewSlice";
import HomeInfo from "./HomeInfo";

// @ts-ignore
function Results(props) {
  const dispatch = useAppDispatch();
  const results = useAppSelector((state) => state.query.results);
  const query = useAppSelector((state) => state.query.value);
  const loading = useAppSelector((state) => state.query.loading);
  const elapsedTime = useAppSelector((state) => state.query.elapsedTime);
  const [showAll, setShowAll] = useState(false);
  const showResults = () => (showAll ? results : results.slice(0, 10));
  const selectedItem = useAppSelector((state) => state.preview.key);
  const selectNextResult = () => dispatch(incrementKey());
  const selectPreviousResult = () => dispatch(decrementKey());
  const openPreview = () => {
    dispatch(updateResultForPreview(results[selectedItem]));
  };
  const closePreview = () => {
    dispatch(hidePreview());
  };
  useEffect(() => {
    const { shortcut } = props;
    shortcut.registerShortcut(
      selectNextResult,
      ["tab"],
      "NextResult",
      "jump to next result"
    );
    shortcut.registerShortcut(
      selectPreviousResult,
      ["shift+tab"],
      "PreviousResult",
      "jump to previous result"
    );
    shortcut.registerShortcut(
      openPreview,
      ["enter"],
      "OpenPreview",
      "open preview of result"
    );
    shortcut.registerShortcut(
      closePreview,
      ["ctrl+q"],
      "ClosePreview",
      "close preview of result"
    );
    return () => {
      const { shortcut } = props;
      shortcut.unregisterShortcut(["tab", "shift+tab"]);
    };
  }, []);
  return (
    <>
      <div className="sidebar-stats">
        {loading ? "loading index...\n" : results.length === 0}
        {!loading && query && elapsedTime !== 0 && results.length === 0
          ? `No results for query: ${query} (took ${elapsedTime} ms)`
          : query &&
            !loading &&
            `${results.length} results (took ${elapsedTime} ms)`}
      </div>
      <div className="search-results search-results-empty">
        <ol className="search-results-list">
          {showResults().map((result, i) => {
            return result != null && <Result key={i} id={i} result={result} />;
          })}
          {results.length > 10 && !showAll && (
            <button
              className="search-results-show-all"
              onClick={() => setShowAll(true)}
            >
              Show more ...
            </button>
          )}
        </ol>
        {results.length === 0 ? <HomeInfo /> : null}
      </div>
    </>
  );
}

export default Results;
