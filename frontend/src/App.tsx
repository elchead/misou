import "./App.css";
import Results from "./components/Results";
import SearchBar from "./components/SearchBar";
import { APIProvider, createAPIClient } from "./services/API";
import { DocPreview } from "./components/DocPreview";
import { useEffect } from "react";
import { withShortcut } from "react-keybind";

const ShortcutSearchBar = withShortcut(SearchBar);
const ShortcutResults = withShortcut(Results);
const ShortcutPreview = withShortcut(DocPreview);

function App() {
  return (
    <APIProvider api={createAPIClient()}>
      <div className="app">
        <div className="sidebar">
          <ShortcutSearchBar />
          <ShortcutResults />
          <ShortcutPreview />
        </div>
      </div>
    </APIProvider>
  );
}

export default App;
