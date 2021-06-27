export type APIType = {
  Prepare: (
    sourcePath: string,
    versionNumber: string,
    packageRawFiles: boolean,
    fileListPath: string
  ) => Promise<string>;
  Bundle: () => Promise<boolean>;
  SelectSourceDir: () => Promise<string>;
  SelectFileListLocation: () => Promise<string>;
  EnsureInnoSetup: () => Promise<void>;
  Exit: () => Promise<void>;
};

export type StoreType = {
  subscribe: (callback: (newState: string[]) => void) => void;
};
