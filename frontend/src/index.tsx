import React, { createContext } from "react";
import ReactDOM from "react-dom";
import "./index.css";
import App from "./App";
import { store } from "./app/store";
import { Provider } from "react-redux";
import { ShortcutProvider } from "react-keybind";

import * as serviceWorker from "./serviceWorker";
import * as Wails from "@wailsapp/runtime";

Wails.Init(() => {
  ReactDOM.render(
    <React.StrictMode>
      <Provider store={store}>
        <ShortcutProvider>
          <App />
        </ShortcutProvider>
      </Provider>
    </React.StrictMode>,
    document.getElementById("app")
  );
});

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
