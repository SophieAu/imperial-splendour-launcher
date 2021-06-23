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
  Exit: () => Promise<void>;
};

export type StoreType = {
  subscribe: (callback: (newState: string[]) => void) => void;
};
