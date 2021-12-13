import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface EmailConfig {
    username: string;
    password: string;
    domain: string;
    port: number;
    tls: boolean;
}

export interface NotionConfig {

}

export interface GoogleDriveConfig {

}

export interface ConfigurationState {
    email: EmailConfig;
    notion: NotionConfig;
    google_drive: GoogleDriveConfig;
}

const initialState = {
    email: {
        username: "",
        password: "",
        domain: "",
        port: 993,
        tls: true,
    },
    notion: {},
    google_drive: {},
} as ConfigurationState

export const configurationSlice = createSlice({
    name:'configuration',
    initialState: initialState,
    reducers: {
        updateEmail: (state, value: PayloadAction<EmailConfig>) => {
            state.email = value.payload
        }
    }
})

export const {updateEmail} = configurationSlice.actions; 

export default configurationSlice.reducer;