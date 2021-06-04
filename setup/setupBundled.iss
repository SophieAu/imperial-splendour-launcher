; Script generated by the Inno Script Studio Wizard.
; SEE THE DOCUMENTATION FOR DETAILS ON CREATING INNO SETUP SCRIPT FILES!

#define MyAppName "Imperial Splendour"
#define MyAppVersion "2.0"
#define MyAppPublisher "Imperial Splendour"
#define MyAppURL "https://imperialsplendour.com/"
#define MyAppExeName "ImperialSplendour.exe"
#define TmpFolder "ImperialSplendour"
#define UninstallDir "IS_Uninstall"
#define UninstallHelperExe "deactivator.exe"
#define SetupName "ImperialSplendourSetup"
#define ETWDefaultPath "steam\steamapp\common\Empire Total War"
#define CampaignPath "data/campaigns/imperial_splendour"

[Setup]
; NOTE: The value of AppId uniquely identifies this application.
; Do not use the same AppId value in installers for other applications.
; (To generate a new GUID, click Tools | Generate GUID inside the IDE.)
AppId={{DC820B8B-AE7B-4555-AB8A-C1399266A06F}
AppName={#MyAppName}
AppVersion={#MyAppVersion}
;AppVerName={#MyAppName} {#MyAppVersion}
AppPublisher={#MyAppPublisher}
AppPublisherURL={#MyAppURL}
AppSupportURL={#MyAppURL}
AppUpdatesURL={#MyAppURL}
DisableDirPage=yes
DefaultDirName={commonpf}
AlwaysShowDirOnReadyPage=yes
DefaultGroupName={#MyAppName}
OutputBaseFilename={#SetupName}
Compression=lzma
SolidCompression=yes
; 10GB = 10,737,418,240
ExtraDiskSpaceRequired=10737418240 
DirExistsWarning=yes
SetupIconFile=appicon.ico
UninstallDisplayIcon={app}\{#MyAppExeName}
UninstallFilesDir={app}\{#UninstallDir}

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"                              

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked

[Files]
Source: "{#TmpFolder}\*.*"; DestDir: "{app}"; Flags: recursesubdirs 

[Dirs]
Name: "{app}\{#CampaignPath}"

[Icons]
Name: "{group}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"
Name: "{group}\{cm:UninstallProgram,{#MyAppName}}"; Filename: "{uninstallexe}"
Name: "{commondesktop}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"; Tasks: desktopicon

[Run]
Filename: "{app}\{#MyAppExeName}"; Description: "{cm:LaunchProgram,{#StringChange(MyAppName, '&', '&&')}}"; Flags: nowait postinstall skipifsilent

[UninstallRun]
Filename: "{app}\{#UninstallDir}\{#UninstallHelperExe}"; WorkingDir: "{app}"

[CustomMessages]
NewerVersionExists=A newer version of {#AppName} is already installed.%n%nInstaller version: {#AppVersion}%nCurrent version: 
ETWNotFound=Couldn''t find your Empire Total War installation. Please make sure you have it installed correctly and manually select the install folder.
DeactivationError=There was an error preparing the upgrade. Please delete any remaining Imperial Splendour files manually and try again.


[Code]

var ExpectedPath: String;
var HasInstallation: Boolean;
var
  StartupInfoPage: TInputQueryWizardPage;
  InputDirPage:    TInputDirWizardPage;


function InitializeSetup: Boolean;
var Version: String;
var Location: String;
begin
  HasInstallation := RegValueExists(HKEY_LOCAL_MACHINE,'Software\Microsoft\Windows\CurrentVersion\Uninstall\{#AppId}_is1', 'DisplayVersion')
  if not HasInstallation then
    Result := True;
    Exit;
  end;

  
  RegQueryStringValue(HKEY_LOCAL_MACHINE,'Software\Microsoft\Windows\CurrentVersion\Uninstall\{#AppId}_is1', 'DisplayVersion', Version);
  if Version > '{#AppVersion}' then
    MsgBox(ExpandConstant('{cm:NewerVersionExists} '+Version), mbInformation, MB_OK);
    Result := False;
    Exit;

  RegQueryStringValue(HKEY_LOCAL_MACHINE,'Software\Microsoft\Windows\CurrentVersion\Uninstall\{#AppId}_is1', 'InstallLocation', Location);
  WizardForm.DirEdit.Text := Location
  Result := True;
end;


procedure InitializeWizard;
begin
  ExpectedPath := ExpandConstant('{commonpf}\{#ETWDefaultPath}');

  StartupInfoPage:= CreateInputQueryPage(wpWelcome,
    'Welcome',
    'Please read the following important information before continuing',
    'Requirements:'#13#10 +
    '- an Empire: Total War installation through Steam'#13#10 +
    '- at least 10GB of hard drive space'#13#10 +
    ''#13#10 +
    'NOTE: Having a different mod installed alongside Imperial Splendour can lead to issues when using the launcher. In this case, you can still use the installer but we recommend to switch between mods manually.');
    

  InputDirPage := CreateInputDirPage(StartupInfoPage.ID,
    'Select existing Empire: Total War Installation Location',
    '',
    'Imperial Splendour will be installed inside the following Empire: Total War folder.'#13#10 +
    ''#13#10 +
    'To continue, click Next. If you would like to select a different folder, click Browse.',
    False, '');
  InputDirPage.Add('');
  if DirExists(ExpectedPath) then
    InputDirPage.Values[0] := ExpectedPath
  else
    InputDirPage.Values[0] := '';
  
end;


function NextButtonClick(CurPageID: Integer): Boolean;
begin
  if CurPageID = StartupInfoPage.ID then
    if InputDirPage.Values[0] = '' then
      MsgBox(ExpandConstant('{cm:ETWNotFound}'), mbError, MB_OK);
    
  if CurPageID = InputDirPage.ID then
    WizardForm.DirEdit.Text := InputDirPage.Values[0];
  
  Result := true;
end;


function ShouldSkipPage(PageID: Integer): Boolean
begin
  if PageID = InputDirPage.ID and HasInstallation then
    Result := True
  else
    Result := False;
end;


function PrepareToInstall(var NeedsRestart: Boolean): String;
var ResultCode: Integer;
begin
  if not Exec(ExpandConstant('{#UninstallDir}\{#UninstallHelperExe}'), '-strict', '', SW_SHOW, ewWaitUntilTerminated, ResultCode) then
    Result := ExpandConstant('{cm:DeactivationError}')
  else
    Result := ""
end;
