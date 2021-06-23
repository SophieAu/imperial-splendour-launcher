export type APIType = {
  Prepare: (
    sourcePath: string,
    versionNumber: string,
    packageRawFiles: boolean,
    fileListPath: string
  ) => Promise<string>;
  Bundle: () => Promise<boolean>;
  Exit: () => Promise<void>;
  SelectSourceDir: () => Promise<string>;
  SelectFileListLocation: () => Promise<string>;
};

export type StoreType = {
  subscribe: (callback: (newState: string[]) => void) => void;
};
