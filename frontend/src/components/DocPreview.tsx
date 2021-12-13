import { Result } from "../app/store";
import { useState, useEffect } from "react";
import { useAppDispatch, useAppSelector } from "../app/hooks";
import { hidePreview } from "../features/previewSlice";
import { getHighlightedText } from "../services/styling";
import { Browser } from "@wailsapp/runtime";

//@ts-ignore
export function DocPreview(props) {
  const res = useAppSelector((state) => state.preview.value);
  const show = useAppSelector((state) => state.preview.show);
  const globalQuery = useAppSelector((state) => state.query.value);
  const dispatch = useAppDispatch();
  useEffect(() => {
    const { shortcut } = props;
    shortcut.registerShortcut(
      () => {
        if (res.link != undefined) Browser.OpenURL(res.link); // TODO make link openable without opening preview
      },
      ["cmd+o"],
      "Open",
      "open link"
    );
    return () => {
      const { shortcut } = props;
      shortcut.unregisterShortcut(["tab", "shift+tab", "cmd+o"]);
    };
  }, [res]);
  if (show) {
    return (
      <div className="doc-preview">
        <div className="doc-preview-buttons">
          <button
            title="Close preview"
            className="button doc-preview-close"
            onClick={() => dispatch(hidePreview())}
          >
            ×
          </button>
          <a
            title="Open on new page"
            href={res.link}
            onClick={(e) => {
              e.preventDefault();
              Browser.OpenURL(res.link);
            }}
            target="_blank"
            className="button doc-preview-open"
          >
            <span className="desktop">open </span>→
          </a>
          <p className="doc-preview-title">
            {getHighlightedText(res.title, globalQuery)}
          </p>
        </div>
        <div className="doc-preview-content">
          <p>{getHighlightedText(res.content, globalQuery)}</p>
        </div>
      </div>
    );
  } else {
    return null;
  }
}
