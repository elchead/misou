import React from "react";
import {
  configureStore,
  combineReducers,
  ThunkAction,
  Action,
} from "@reduxjs/toolkit";
import queryReducer, { loadResults } from "./querySlice";
import resultsReducer from "./resultsSlice";
import { API } from "../services/wailsApi";
const getStore = () => {
  const mockApi = {
    Search: (query: string) =>
      Promise.resolve(`[{"title":"hi"},{"title":"Sa"}]`),
  };
  return getStoreWithApi(mockApi);
};

const rootReducer = combineReducers({
  query: queryReducer,
  results: resultsReducer,
});

const getStoreWithApi = (api: API) =>
  configureStore({
    reducer: rootReducer,
    middleware: (getDefaultMiddleware) =>
      getDefaultMiddleware({
        thunk: {
          extraArgument: api,
        },
      }),
  });

describe("when load results", () => {
  describe("initially", () => {
    it("does not have loading flag set", () => {
      const store = getStore();
      expect(store.getState().query.loading).toEqual(false);
    });
  });
  describe("while loading", () => {
    const api = {
      Search: (query: string) => new Promise<string>(() => {}),
    };
    const store = getStoreWithApi(api);
    it("has loading flag set", () => {
      store.dispatch(loadResults("hi"));
      expect(store.getState().query.loading).toEqual(true);
    });
  });
  describe("after loading", () => {
    it("gets results and has no loading flag set", async () => {
      const store = getStore();
      await store.dispatch(loadResults("hi"));
      expect(store.getState().query.loading).toEqual(false);
      expect(store.getState().query.value).toEqual("hi");
      expect(store.getState().query.results).toEqual([
        { title: "hi" },
        { title: "Sa" },
      ]);
    });
  });
});
