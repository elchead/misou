import { useEffect } from "react";
import { Result } from "../app/store";
import { useAppDispatch, useAppSelector } from "../app/hooks";
import { updateResultForPreview, updateKey } from "../features/previewSlice";
import { getHighlightedText } from "../services/styling";
import { Browser } from "@wailsapp/runtime";
interface ResultCardProps {
  result: Result;
}

function ResultItem({ id, result }: { result: Result; id: number }) {
  const dispatch = useAppDispatch();
  const selectedItem = useAppSelector((state) => state.preview.key);
  const previewItem = useAppSelector((state) => state.preview.show);
  var style = "search-result";
  if (selectedItem === id) {
    style += " selected";
    if (previewItem) handleResult(id, result);
  }
  const globalQuery = useAppSelector((state) => state.query.value);
  function handleResult(key: number, result: Result) {
    if (result.provider === "history" || result.provider === "bookmark") {
      // TODO better use result.Content === "" ?
      Browser.OpenURL(result.link);
    } else {
      dispatch(updateResultForPreview(result));
      dispatch(updateKey(key));
    }
  }
  return (
    // <ul className="search-result"></ul>
    <li onClick={() => handleResult(id, result)} className={style}>
      <span className="search-result-module">{result.provider}</span>
      <span className="search-result-title">{result.title}</span>
      <span data-doc-id="www151" className="search-result-content">
        {getHighlightedText(result.content, globalQuery)}
      </span>
    </li>
  );
}

export default ResultItem;
