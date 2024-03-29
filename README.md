# Imperial Splendour Launcher Toolkit
> Imperial Splendour attempts to create the best Empire: Total War experience possible without destroying the essence of the game, with an eye towards expanding and enriching the experience, while adding to the game's realism whenever possible.

It is the year 1783; we begin the adventure just after the American Revolution which shows to the world that the order of a monarch can be disputed by his own subjects and that a colony can claim its independence in name of Liberty. However, in Europe, most of regimes are absolute monarchies; the American events are really bad news for the stability of their reign. Such events inspired by the Enlightenment could spread and tilt monarchies in place for nearly a thousand years.

![](header.png)

[About this Repo](#about-this-repo) | [Releasing a New Version](#releasing-a-new-imperial-splendour-version) | [Global constants](#global-constants) | [Bundling Imperial Splendour Manually](#bundling-imperial-splendour-manually) | [Credits and License Information](#credits-and-license-information)

## About this Repo

This repo contains the various apps to help create the best user experience for playing Imperial Splendour:
* The Launcher: the most integral part which will be used by the user regularly
* The Bundler: Used by the Imperial Splendour devs to create setup file for the users
* The Deactivator: A script making sure that everything goes well when uninstalling Imperial Splendour
* The InnoSetup scripts: Used by the bundler to create the setup.exe for the users. Currently, only the `setupBundled.iss` file is being used.


## Releasing a new Imperial Splendour Version
1. Upload the bundled setup to the following hosts:
    * Google Drive
    * ModDB
    * MediaFire
2. Optional: Upload zipped 'raw-mod-files-only' (so, only the `IS_Files/` folder without the `IS_Info.json` file) versions of the mod update.
3. Update the download links on the website
4. Update `version.txt` with the new version number
5. Optional: Add a release blog post to the website

**And finally:**
* Push the updated website
* post the news on social media and Discord



## Global constants

### Locations

* Launcher: `launcher/backend/a_constants.go`
* Deactivator: `setup/deactivator/app/a_constants.go`
* Bundled Setup: `setup/setupBundled.iss`
* Slim Setup: `setup/setupSlim.iss`
* Bundler: `bundler/backend/a_constants.go`

### ETW Constants
| Constant                      | Launcher     | Deactivator  | Bundled Setup  | Slim Setup     | Bundler    |
|-------------------------------|--------------|--------------|----------------|----------------|------------|
| Default ETW Installation Path | -/-          | -/-          | ETWDefaultPath | ETWDefaultPath | -/-        |
| APPDATA Folder                | appDataPath  | appDataPath  | -/-            | -/-            | -/-        |
| Data Folder                   | dataPath     | dataPath     | -/-            | -/-            | -/-        |
| Campaign Folder               | campaignPath | campaignPath | CampaignPath   | CampaignPath   | -/-        |
| User Script                   | userScript   | userScript   | -/-            | -/-            | userScript |
| Steam URI                     | etwSteamURI  | -/-          | -/-            | -/-            | -/-        |



### Imperial Splendour-specific Constants

| Constant           | Launcher      | Deactivator   | Bundled Setup | Slim Setup   | Bundler       |
|--------------------|---------------|---------------|---------------|--------------|---------------|
| Mod Files Folder   | modPath       | modPath       | -/-           | -/-          | modPath       |
| Uninstall Folder   | uninstallPath | uninstallPath | UninstallDir  | UninstallDir | uninstallPath |
| File List Location | fileListFile  | fileListFile  | -/-           | -/-          | fileListFile  |
| Info File Location | infoFile      | infoFile      | -/-           | -/-          | infoFile      |
| App Name           | AppName       | -/-           | MyAppName     | MyAppName    |               |
| Website            | websiteURL    | -/-           | MyAppURL      | MyAppURL     | -/-           |

### Setup-specific Constants

| Constant             | Bundled Setup      | Slim Setup         | Bundler         |
|----------------------|--------------------|--------------------|-----------------|
| App Version          | MyAppVersion       | MyAppVersion       |                 |
| App Publisher        | MyAppPublisher     | MyAppPublisher     |                 |
| Launcher File Name   | MyAppExeName       | MyAppExeName       | launcherFile    |
| Setup Display Name   | SetupName          | SetupName          |                 |
| Mod Files DLoad Link | -/-                | DownloadLink       |                 |
| Deactivator Name     | UninstallHelperExe | UninstallHelperExe | deactivatorFile |
| Temp Folder          | TmpFolder          | TmpFolder          | tempPath        |




## Bundling Imperial Splendour Manually

### Requirements

Note that the slim setup is currently theoretical only as hosting large files is too expensive to be sustainable for us. For potential future use, `setupSlim.iss` should still be updated though.

* a Windows machine
* at least 15GB of hard drive space (for the mod files)
* a current install of [Inno Setup](https://jrsoftware.org/isinfo.php)
* the `deactivator.exe` and `ImperialSplendour.exe` files which you can find in the `artifacts/` folder of this repo
* the `setupBundled.iss` and `appicon.ico` files which you can find in the `setup/` folder of this repo

### 1. Update Imperial Splendour Version

The version needs to be updated in the following places:
* Website (the `<URL>/version`) (**❗️Hold off on pushing the change though**)
* `IS_Info.json` file
* Setup files (both `setupBundled.iss` and `setupSlim.iss`)

### 2. Collect Mod Files

Collect all mod files in one otherwise empty folder. The folder structure will need to be:
* `ImperialSplendour/`
  * `ImperialSplendour.exe` ('_Launcher File Name_' in the [Setup-specific Constants Table](#setup-specific-constants))
  * `IS_Files/` ('_Mod Files Folder_' in the [Imperial Splendour-specific Constants Table](#imperial-splendour-specific-constants))
    * `IS_FileList.txt` ('_File List Location_' in the [Imperial Splendour-specific Constants Table](#imperial-splendour-specific-constants))
    * `IS_Info.json` ('_Info File Location_' in the [Imperial Splendour-specific Constants Table](#imperial-splendour-specific-constants))
    * `user.empire_script.txt` ('_User Script_' in the [ETW Constants Table](#etw-constants))
    * all the other mod files (`.pack`, `.tga`, `.esf` files only, *need* to be listed in the `IS_FileList.txt`)
  * `IS_Uninstall/` ('_Uninstall Folder_' in the [Imperial Splendour-specific Constants Table](#imperial-splendour-specific-constants))
    * `deactivator.exe` ('_Deactivator Name_' in the [Setup-specific Constants Table](#setup-specific-constants))
* `setupBundled.iss`
* `appicon.ico`

### 3. Build the Setup

Add the `setupBundled.iss` file to the mod file folder and then either use the Inno Setup Compiler IDE or the cli to build the setup.


## Credits and License Information
* Font: [IM Fell English SC](https://fonts.google.com/specimen/IM+Fell+English+SC) and [IM Fell English](https://fonts.google.com/specimen/IM+Fell+English) by Igino Marini
* ETW Logo: https://www.pinclipart.com/maxpin/xRioh/

Distributed under the MIT license. See ``LICENSE.md`` for more information.
