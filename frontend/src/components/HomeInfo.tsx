function HomeInfo() {
  return (
    <div className="search-results search-results-empty">
      {/* <h2 className="empty-state-heading">Suggestions</h2> */}
      {/* <div className="search-results-suggestions">
        <button className="search-results-suggestion">linus lee</button>
        <button className="search-results-suggestion">side project idea</button>
        <button className="search-results-suggestion">tools for thought</button>
        <button className="search-results-suggestion">
          incremental note-taking
        </button>
        <button className="search-results-suggestion">taylor swift</button>
        <button className="search-results-suggestion">uc berkeley</button>
        <button className="search-results-suggestion">new york city</button>
        <button className="search-results-suggestion">dorm room fund</button>
      </div> */}
      <h2 className="empty-state-heading">Keybindings</h2>
      <div className="keyboard-map">
        <ul className="keyboard-map-list">
          <li className="keyboard-map-item">
            <div className="keybinding-keys">
              <kbd className="">Tab</kbd>
            </div>
            <div className="keybinding-detail">Next search result</div>
          </li>
          <li className="keyboard-map-item">
            <div className="keybinding-keys">
              <kbd className="">Shift</kbd>
              <kbd className="">Tab</kbd>
            </div>
            <div className="keybinding-detail">Previous search result</div>
          </li>
          <li className="keyboard-map-item">
            <div className="keybinding-keys">
              <kbd className="">Enter</kbd>
            </div>
            <div className="keybinding-detail">Show preview pane</div>
          </li>
          <li className="keyboard-map-item">
            <div className="keybinding-keys">
              <kbd className="">Cmd+o</kbd>
            </div>
            <div className="keybinding-detail">Open link (inside preview)</div>
          </li>
          <li className="keyboard-map-item">
            <div className="keybinding-keys">
              <kbd className="">Ctrl+q</kbd>
            </div>
            <div className="keybinding-detail">Hide preview pane</div>
          </li>
          <li className="keyboard-map-item">
            <div className="keybinding-keys">
              <kbd className="">Ctrl+w</kbd>
            </div>
            <div className="keybinding-detail">Focus search box</div>
          </li>
          <li className="keyboard-map-item">
            <div className="keybinding-keys">
              <kbd className="">Ctrl+w && Cmd+del</kbd>
            </div>
            <div className="keybinding-detail">Delete search query</div>
          </li>
          {/*
          <li className="keyboard-map-item">
            <div className="keybinding-keys">
              <kbd className="">`</kbd>
            </div>
            <div className="keybinding-detail">
              Switch light/dark color theme
            </div>
          </li> */}
        </ul>
      </div>
      <h2 className="empty-state-heading">About Misou</h2>
      <div className="about">
        <p className="">
          Misou is a personal search engine inspired by{" "}
          <a
            href="https://github.com/thesephist/monocle"
            target="_blank"
            className=""
          >
            Monocle
          </a>
          .
          {/* <a href="https://thesephist.com" target="_blank" className="">
            Linus
          </a>
          . It's built with{" "}
          <a href="https://dotink.co" target="_blank" className="">
            Ink
          </a>{" "}
          and{" "}
          <a
            href="https://github.com/thesephist/torus"
            target="_blank"
            className=""
          >
            Torus
          </a>
          , and open source on{" "}
          <a
            href="https://github.com/thesephist/monocle"
            target="_blank"
            className=""
          >
            GitHub
          </a>
          . */}
        </p>
        {/* <p className="">
          Monocle is powered by a full-text indexing and search engine written
          in pure Ink, and searches across Linus's blogs and private note
          archives, contacts, tweets, and over a decade of journals.
        </p> */}
      </div>
    </div>
  );
}

export default HomeInfo;
