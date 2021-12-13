import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Result } from "../app/store";

export const resultsSlice = createSlice({
  name: "results",
  initialState: {
    value: new Array<Result>(),
  },
  reducers: {
    clearResults: (state) => {
      state.value = [];
    },
    updateResult: (state, action: PayloadAction<Result>) => {
      state.value.push(action.payload);
    },
    updateResults: (state, action: PayloadAction<Array<Result>>) => {
      state.value = state.value.concat(action.payload);
    },
    replaceResults: (state, action: PayloadAction<Array<Result>>) => {
      state.value = action.payload;
    },
    replaceResult: (state, action: PayloadAction<Result>) => {
      state.value = [action.payload];
    },
  },
});

export const {
  clearResults,
  updateResult,
  updateResults,
  replaceResults,
  replaceResult,
} = resultsSlice.actions;

export default resultsSlice.reducer;
