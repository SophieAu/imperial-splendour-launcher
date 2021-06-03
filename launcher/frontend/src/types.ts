export type APIType = {
  Version: () => Promise<string>;
  IsActive: () => Promise<boolean>;
  Play: () => Promise<void>;
  Switch: () => Promise<boolean>;
  GoToWebsite: () => Promise<void>;
  Uninstall: () => Promise<void>;
  Exit: () => Promise<void>;
};
