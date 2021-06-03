# Imperial Splendour Launcher

## How to Create a New Imperial Splendour Release

### Requirements

Note that the slim setup is currently theoretical only as hosting large files is too expensive to be sustainable for us. For potential future use, `setupSlim.iss` should still be updated though.

* a Windows machine
* at least 15GB of hard drive space (for the mod files)
* a current install of [Inno Setup](https://jrsoftware.org/isinfo.php) 


### 1. Update Imperial Splendour Version

The version needs to be updated in the following places:
* Website (the `<URL>/version`) (**❗️Hold off on pushing the change though**)
* `IS_Info.json` file
* Setup files (both `setupBundled.iss` and `setupSlim.iss`)

### 2. Collect Mod Files

Collect all mod files in one otherwise empty folder. The folder structure will need to be:
* `ImperialSplendour.exe` ('_Launcher File Name_' in the [Setup-specific Constants Table](#setup-specific-constants))
* `IS_Files/` ('_Mod Files Folder_' in the [Imperial Splendour-specific Constants Table](#imperial-splendour-specific-constants))
  * `IS_FileList.txt` ('_File List Location_' in the [Imperial Splendour-specific Constants Table](#imperial-splendour-specific-constants))
  * `IS_Info.json` ('_Info File Location_' in the [Imperial Splendour-specific Constants Table](#imperial-splendour-specific-constants))
  * `user.empire_script.txt` ('_User Script_' in the [ETW Constants Table](#etw-constants))
  * all the other mod files (`.pack`, `.tga`, `.esf` files only, *need* to be listed in the `IS_FileList.txt`)
* `IS_Uninstaller/` ('_Uninstall Folder_' in the [Imperial Splendour-specific Constants Table](#imperial-splendour-specific-constants))
  * `deactivator.exe` ('_Deactivator Name_' in the [Setup-specific Constants Table](#setup-specific-constants))
* `setupBundled.iss`

### 3. Build the Setup

Add the `setupBundled.iss` file to the mod file folder and then either use the Inno Setup Compiler IDE or the cli to build the setup.


### 4. Finish Up
1. Upload the bundled setup to the following hosts:
    * Google Drive
    * ModDB
    * MediaFire
2. Optional: Upload zipped 'raw-mod-files-only' versions of the mod update
3. Update the download links on the website
4. Optional: Add a release blog post to the website

```
!!!!!!!!!!!!!!!!!!! MAYBE ADD A NEW SECTION TO THE DOWNLOAD PAGE
```

And finally:
* Push the updated website
* post the news on social media and Discord



## Global constants

### Locations

* Launcher: `launcher/backend/a_constants.go`
* Deactivator: `setup/deactivator/app/a_constants.go`
* Bundled Setup: `setup/setupBundled.iss`
* Slim Setup: `setup/setupSlim.iss`
* Website: `<website repo>/src/constants.ts`

### ETW Constants
| Constant                      | Launcher     | Deactivator  | Bundled Setup  | Slim Setup     | Website |
|-------------------------------|--------------|--------------|----------------|----------------|---------|
| Default ETW Installation Path | -/-          | -/-          | ETWDefaultPath | ETWDefaultPath | -/-     |
| APPDATA Folder                | appDataPath  | appDataPath  | -/-            | -/-            | -/-     |
| Data Folder                   | dataPath     | dataPath     | -/-            | -/-            | -/-     |
| Campaign Folder               | campaignPath | campaignPath | -/-            | -/-            | -/-     |
| User Script???                | userScript   | userScript   | -/-            | -/-            | -/-     |
| Steam URI                     | etwSteamURI  | -/-          | -/-            | -/-            | -/-     |



### Imperial Splendour-specific Constants

| Constant                  | Launcher      | Deactivator   | Bundled Setup      | Slim Setup         | Website |
|---------------------------|---------------|---------------|--------------------|--------------------|---------|
| Mod Files Folder          | modPath       | modPath       | -/-                | -/-                | -/-     |
| Uninstall Folder          | uninstallPath | uninstallPath | UninstallDir       | UninstallDir       | -/-     |
| File List Location        | fileListFile  | fileListFile  | -/-                | -/-                | -/-     |
| Info File Location        | infoFile      | infoFile      | -/-                | -/-                | -/-     |
| App Name                  | AppName       | -/-           | MyAppName          | MyAppName          | -/-     |
| Website                   | websiteURL    | -/-           | MyAppURL           | MyAppURL           | -/-     |

### Setup-specific Constants

| Constant             | Bundled Setup  | Slim Setup     | Website |
|----------------------|----------------|----------------|---------|
| App Version                | MyAppVersion       | MyAppVersion       | -/-     |
| App Publisher        | MyAppPublisher | MyAppPublisher | -/-     |
| Launcher File Name   | MyAppExeName   | MyAppExeName   | -/-     |
| Setup Display Name   | SetupName      | SetupName      | -/-     |
| Mod Files DLoad Link | -/-            | DownloadLink   | -/-     |
| Deactivator Name | UninstallHelperExe | UninstallHelperExe | -/-     |





## Credits
* Font: [IM Fell English SC](https://fonts.google.com/specimen/IM+Fell+English+SC) and [IM Fell English](https://fonts.google.com/specimen/IM+Fell+English) by Igino Marini
* ETW Logo: https://www.pinclipart.com/maxpin/xRioh/