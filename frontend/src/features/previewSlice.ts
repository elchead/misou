import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Result } from "../app/store";

export const previewSlice = createSlice({
  name: "preview",
  initialState: {
    value: {} as Result,
    show: false,
    key: 0,
  },
  reducers: {
    updateResult: (state, action: PayloadAction<Result>) => {
      state.value = action.payload;
      state.show = true;
    },
    updateKey: (state, action: PayloadAction<number>) => {
      state.key = action.payload;
    },
    incrementKey: (state) => {
      state.key++;
    },
    decrementKey: (state) => {
      state.key--;
    },
    hidePreview: (state) => {
      state.show = false;
    },
  },
});

export const {
  updateResult: updateResultForPreview,
  updateKey,
  incrementKey,
  decrementKey,
  hidePreview,
} = previewSlice.actions;

export default previewSlice.reducer;
