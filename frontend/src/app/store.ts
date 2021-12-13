import { configureStore, ThunkAction, Action } from "@reduxjs/toolkit";
import resultsReducer from "../features/resultsSlice";
import queryReducer from "../features/querySlice";
import loadingReducer from "../features/loadingSlice";
import configurationReducer from "../features/configurationSlice";
import previewReducer from "../features/previewSlice";
import { wailsApi } from "../services/wailsApi";

export const store = configureStore({
  reducer: {
    results: resultsReducer,
    query: queryReducer,
    loading: loadingReducer,
    configuration: configurationReducer,
    preview: previewReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      thunk: {
        extraArgument: wailsApi,
      },
    }),
});

export interface Result {
  title: string;
  link: string;
  content: string;
  provider: string;
  matches: number;
  contentType: string;
}

// export interface RichResult {
//   title: String;
//   link: string;
//   parts: Array<Part>;
//   provider: String;
// }

// export function isRichResult(result: any): result is RichResult {
//   return "parts" in result
// }

export function isHtmlResult(result: Result): result is Result {
  return result.contentType === "html";
}

export interface Part {
  content: String;
  highlight: boolean;
  index: number;
}

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
