#define Name "Clipboard Monitor"
#define Version "1.1.0"
#define Publisher "Cupinxas ltda."
#define URL "https://github.com/andersoncorso/clipboard-monitor"
#define ExeName "ClipboardMonitor.exe"
#define ExeDir "C:\Users\anderson_corso\Documents\dev\clipboard-monitor"
#define BuildDir "C:\Users\anderson_corso\Documents\dev\clipboard-monitor\build"

[Setup]
AppId={{407F42BD-97FF-47A2-9F83-521118AB32C7}}
AppName={#Name}
AppVersion={#Version}
AppPublisher={#Publisher}
AppPublisherURL={#URL}
AppSupportURL={#URL}
AppUpdatesURL={#URL}
DefaultDirName={autopf}\{#Name}
UninstallDisplayIcon={app}\{#ExeName}
ArchitecturesAllowed=x64compatible
ArchitecturesInstallIn64BitMode=x64compatible
DisableProgramGroupPage=yes
PrivilegesRequired=lowest
OutputBaseFilename={#Name}_v{#Version}_Setup
OutputDir=C:\Users\anderson_corso\Documents\dev\clipboard-monitor\installer
SetupIconFile={#BuildDir}\assets\icon.ico
SolidCompression=yes
WizardStyle=modern

[Languages]
Name: "brazilianportuguese"; MessagesFile: "compiler:Languages\BrazilianPortuguese.isl"

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked

[Files]
Source: "{#ExeDir}\{#ExeName}"; DestDir: "{app}"; Flags: ignoreversion
Source: "{#BuildDir}\data\data.txt"; DestDir: "{app}\build\data"; Flags: ignoreversion
Source: "{#BuildDir}\assets\icon.ico"; DestDir: "{app}\build\assets"; Flags: ignoreversion

[Icons]
Name: "{autoprograms}\{#Name}"; Filename: "{app}\{#ExeName}"; IconFilename: "{app}\build\assets\icon.ico"
Name: "{autodesktop}\{#Name}"; Filename: "{app}\{#ExeName}"; Tasks: desktopicon; IconFilename: "{app}\build\assets\icon.ico"

[Run]
Filename: "{app}\{#ExeName}"; Description: "{cm:LaunchProgram,{#StringChange(Name, '&', '&&')}}"; Flags: nowait postinstall skipifsilent

[Registry]
Root: HKCU; Subkey: "Software\Microsoft\Windows\CurrentVersion\Run"; ValueType: string; ValueName: "{#Name}"; ValueData: """{app}\{#ExeName}"""; Flags: uninsdeletevalue
