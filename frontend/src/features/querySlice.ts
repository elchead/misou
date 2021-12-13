import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { AppDispatch } from "../app/store";
import { updateResult, updateResults } from "./resultsSlice";
import { Result } from "../app/store";
interface QueryState {
  value: string;
  loading: boolean;
  results: Result[];
  elapsedTime: number;
}

const initalState = {
  value: "",
  loading: false,
  results: [],
  elapsedTime: 0,
} as QueryState;

export const querySlice = createSlice({
  name: "query",
  initialState: initalState,
  reducers: {
    clearQuery: (state) => {
      state.value = "";
      state.results = [];
      state.loading = false;
    },
    updateQuery: (state, action: PayloadAction<string>) => {
      state.value = action.payload;
      state.loading = true;
    },
    addResult: (state, action: PayloadAction<Result>) => {
      state.results.push(action.payload);
      state.loading = false;
    },
    setResults: (state, action: PayloadAction<Result[]>) => {
      state.results = action.payload === null ? [] : action.payload;
      state.loading = false;
    },
    updateTime: (state, action: PayloadAction<number>) => {
      state.elapsedTime = action.payload;
    },
  },
});

interface api {
  Search: (query: string) => Promise<string>;
}

export const loadResults =
  (query: string) => async (dispatch: AppDispatch, _: any, api: api) => {
    dispatch(updateQuery(query));
    const startTime = new Date();
    const res = await api.Search(query);
    const endTime = new Date();
    const elapsedTime = endTime.getTime() - startTime.getTime();
    dispatch(querySlice.actions.updateTime(elapsedTime));
    dispatch(querySlice.actions.setResults(JSON.parse(res)));
    dispatch(updateResults(JSON.parse(res))); // TODO
  };

export const { clearQuery, updateQuery } = querySlice.actions;

export default querySlice.reducer;
