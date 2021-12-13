import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const isLoading = (providers: any) => {
  return providers.notion || providers.email || providers.google_drive || providers.filesystem
} 

export const loadingSlice = createSlice({
  name: "loading",
  initialState: {
    providers: {
      notion: false,
      email: false,
      google_drive: false,
      filesystem: false,
    },
    value: false,
  },
  reducers: {
    setNotionLoading: (state, value: PayloadAction<boolean>) => {
      state.providers.notion = value.payload;
      state.value = isLoading(state.providers)
    },
    setEmailLoading: (state, value: PayloadAction<boolean>) => {
      state.providers.email = value.payload;
      state.value = isLoading(state.providers)
    },
    setGoogleDriveLoading: (state, value: PayloadAction<boolean>) => {
      state.providers.google_drive = value.payload;
      state.value = isLoading(state.providers)
    },
    setFilesystemLoading: (state, value: PayloadAction<boolean>) => {
      state.providers.filesystem = value.payload;
      state.value = isLoading(state.providers)
    },
    setLoading: (state, value: PayloadAction<boolean>) => {
      state.providers = {
        notion: value.payload,
        email: value.payload,
        google_drive: value.payload,
        filesystem: value.payload,
      }
      state.value = value.payload
    }
  },
});

export const { setNotionLoading, setEmailLoading, setFilesystemLoading, setGoogleDriveLoading, setLoading } = loadingSlice.actions;

export default loadingSlice.reducer;
