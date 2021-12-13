export interface API {
  Search: (query: string) => Promise<string>;
}

export const wailsApi = {
  //@ts-ignore
  Search: (query: string) => window.backend.Search(query),
};
