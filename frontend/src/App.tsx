import "./App.css";
import Results from "./components/Results";
import SearchBar from "./components/SearchBar";
import { DocPreview } from "./components/DocPreview";
import { useEffect } from "react";
import { withShortcut } from "react-keybind";

const ShortcutSearchBar = withShortcut(SearchBar);
const ShortcutResults = withShortcut(Results);
const ShortcutPreview = withShortcut(DocPreview);

function App() {
  return (
    <div className="app">
      <div className="sidebar">
        <ShortcutSearchBar />
        <ShortcutResults />
        <ShortcutPreview />
      </div>
    </div>
  );
}

export default App;
