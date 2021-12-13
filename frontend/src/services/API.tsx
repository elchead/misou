import React, { createContext, useMemo } from "react";
import axios, { AxiosInstance } from "axios";

import {
  clearResults,
  replaceResult,
  updateResults,
} from "../features/resultsSlice";
import {
  setEmailLoading,
  setFilesystemLoading,
  setGoogleDriveLoading,
  setLoading,
  setNotionLoading,
} from "../features/loadingSlice";
import { ConfigurationState } from "../features/configurationSlice";
import { useAppDispatch } from "../app/hooks";
import { store } from "../app/store";
import type { AppDispatch } from "../app/store";
export interface API {
  ws: WebSocket;
  api: AxiosInstance;
}

export const WSURL = "ws://localhost:5000/api/ws";

export const createAPIClient = (): API => {
  const ws = new WebSocket(WSURL);

  var http_url: string;
  if (process.env.REACT_APP_GERSTLER_PORT) {
    http_url = `${window.location.hostname}:${process.env.REACT_APP_GERSTLER_PORT}`;
  } else {
    if (process.env.NODE_ENV === "development") {
      http_url = "localhost:5000";
    } else {
      if (window.location.port === "80") {
        http_url = `${window.location.hostname}:80`;
      } else {
        http_url = `${window.location.hostname}:5000`;
      }
    }
  }

  const secure = window.location.protocol === "https" ? true : false;
  const api = axios.create({
    baseURL: `${window.location.protocol}://${http_url}/api/`,
  });
  return { api: api, ws: ws };
};

const connect = (dispatch: AppDispatch, ws: WebSocket) => {
  ws.onmessage = (message) => {
    // console.debug("RECEIVED", message.data);
    let obj = JSON.parse(message.data);
    if ("action" in obj) {
      switch (obj.action) {
        case "results": {
          const query = store.getState().query.value;
          if (obj.query === query) {
            if (obj.data) {
              dispatch(updateResults(obj.data));
            }
          }
          break;
        }
        case "loading_status": {
          const query = store.getState().query.value;
          if (obj.query === query) {
            switch (obj.data.provider.toLowerCase()) {
              case "email":
                dispatch(setEmailLoading(obj.data.loading));
                break;
              case "notion":
                dispatch(setNotionLoading(obj.data.loading));
                break;
              case "gdrive":
                dispatch(setGoogleDriveLoading(obj.data.loading));
                break;
              case "filesystem":
                dispatch(setFilesystemLoading(obj.data.loading));
                break;
            }
          }
        }
      }
    } else {
      console.error("Invalid message %s", message);
    }
  };

  ws.onclose = () => {
    setTimeout(() => {
      // connect();
    }, 1000);
  };

  return ws;
};

export function send(api: API, payload: string) {
  if (
    api.ws.readyState === api.ws.CLOSED ||
    api.ws.readyState === api.ws.CLOSING
  ) {
    console.error("Websocket connection not active");
    // TODO: retry
  } else if (api.ws.readyState === api.ws.CONNECTING) {
    console.debug("CONNECTING");
    setTimeout(() => send(api, payload), 1000);
  } else {
    api.ws.send(payload);
  }
}

export function updateConfig(api: API, config: ConfigurationState) {}

export function getConfig(api: API) {
  axios.get("config").then((response) => {});
}

export function sendQuery(api: API, query: String) {
  // @ts-ignore
  window.backend.Search(query).then((response) => {api.dispatch(updateResults(JSON.parse(response)));}).catch(console.log)
  // window.backend.
  send(
    api,
    JSON.stringify({
      action: "query",
      query: query,
    })
  );
}

export const APIContext = createContext<API>({} as API);

export function APIProvider({
  children,
  api,
}: {
  children: React.ReactNode;
  api: API;
}) {
  const dispatch = useAppDispatch();
  connect(dispatch, api.ws);
  // @ts-ignore
  api.dispatch = dispatch;
  return <APIContext.Provider value={api}>{children}</APIContext.Provider>;
}
