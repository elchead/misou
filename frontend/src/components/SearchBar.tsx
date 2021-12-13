import { connect } from "react-redux";
import {
  ChangeEvent,
  useEffect,
  useState,
  MouseEvent,
  useContext,
  useRef,
} from "react";
import { MdClear } from "react-icons/md";
import { useDebounce } from "use-debounce/lib";
import { useAppDispatch, useAppSelector } from "../app/hooks";
import { setLoading } from "../features/loadingSlice";
import { loadResults, clearQuery, updateQuery } from "../features/querySlice";

import { IShortcutProviderRenderProps } from "react-keybind";
import { clearResults } from "../features/resultsSlice";

function SearchBar({
  loadResults,
  clearResults = () => {},
}: {
  loadResults: (query: string) => void;
  clearResults: () => void;
}) {
  const clearQueryAndResults = () => {
    setQuery("");
    dispatch(clearQuery());
    clearResults();
    dispatch(setLoading(false));
  };
  const globalQuery = useAppSelector((state) => state.query.value);
  const [query, setQuery] = useState(globalQuery);
  const inputElement = useRef(null);
  const dispatch = useAppDispatch();

  const [debouncedQuery] = useDebounce(query, 100);

  useEffect(() => {
    if (debouncedQuery.length > 0) {
      loadResults(debouncedQuery);
    }
  }, [debouncedQuery]);

  const onChange = (event: ChangeEvent<HTMLInputElement>) => {
    if (event.target.value) {
      setQuery(event.target.value);
      dispatch(updateQuery(event.target.value));
    } else {
      clearQueryAndResults();
    }
  };

  const onClick = (event: MouseEvent<HTMLElement>) => {
    event.preventDefault();
    clearQueryAndResults();
  };

  return (
    <div className="search-box">
      <input
        placeholder="Type to search..."
        className="search-box-input"
        autoFocus
        value={query}
        onChange={onChange}
        ref={inputElement}
      />
      {query !== "" ? (
        <button
          title="Clear search"
          className="search-box-clear"
          onClick={onClick}
        >
          ×
        </button>
      ) : null}
    </div>
  );
}

// TODO DRY! but need ref to input element...
function SearchBarWithShortcuts({
  loadResults,
  clearResults,
  shortcut,
}: {
  loadResults: (query: string) => void;
  clearResults: () => void;
  shortcut?: any;
}) {
  const clearQueryAndResults = () => {
    setQuery("");
    dispatch(clearQuery());
    clearResults();
    dispatch(setLoading(false));
  };
  const globalQuery = useAppSelector((state) => state.query.value);
  const [query, setQuery] = useState(globalQuery);
  const inputElement = useRef(null);
  const dispatch = useAppDispatch();

  const [debouncedQuery] = useDebounce(query, 300);

  useEffect(() => {
    // TODO:
    shortcut.registerShortcut(
      //@ts-ignore
      () => inputElement.current.focus(),
      ["ctrl+w"],
      "Focus",
      "focus search"
    );
    if (debouncedQuery.length > 0) {
      loadResults(debouncedQuery);
    }
  }, [debouncedQuery]);

  const onChange = (event: ChangeEvent<HTMLInputElement>) => {
    if (event.target.value) {
      setQuery(event.target.value);
      dispatch(updateQuery(event.target.value));
    } else {
      clearQueryAndResults();
    }
  };

  const onClick = (event: MouseEvent<HTMLElement>) => {
    event.preventDefault();
    clearQueryAndResults();
  };

  return (
    <div className="search-box">
      <input
        placeholder="Type to search..."
        className="search-box-input"
        autoFocus
        value={query}
        onChange={onChange}
        ref={inputElement}
      />
      {query !== "" ? (
        <button
          title="Clear search"
          className="search-box-clear"
          onClick={onClick}
        >
          ×
        </button>
      ) : null}
    </div>
  );
}

const mapStateToProps = null;
const mapDispatchToProps = { loadResults, clearResults };
export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SearchBarWithShortcuts);
export { SearchBar as SearchBarPure };
