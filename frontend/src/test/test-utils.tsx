import React from "react";
import { render, RenderOptions } from "@testing-library/react";
import { createAPIClient, APIProvider } from "../services/API";
import { store } from "../app/store";
import { Provider } from "react-redux";
import axios, { AxiosInstance } from "axios";

const AllTheProviders = ({ children }: { children: React.ReactNode }) => {
  return (
    <React.StrictMode>
      <Provider store={store}>
        <APIProvider api={createAPIClient()}>{children}</APIProvider>
      </Provider>
    </React.StrictMode>
  );
};

const customRender = (
  ui: React.ReactElement<any, string | React.JSXElementConstructor<any>>,
  options?: Omit<RenderOptions, "queries">
) =>
  render(ui, { wrapper: AllTheProviders as React.ComponentType, ...options });

// re-export everything
export * from "@testing-library/react";

// override render method
export { customRender as render };
