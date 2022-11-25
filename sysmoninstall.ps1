$Url="https://download.sysinternals.com/files/Sysmon.zip"
$DownloadZipFile = $HOME + "\Downloads\" + $(Split-Path -Path $Url -Leaf)
$Sys32 = "C:\Windows\System32\"
Invoke-WebRequest -Uri $Url -OutFile $DownloadZipFile
$ExtractShell = New-Object -ComObject Shell.Application 
$ExtractFiles = $ExtractShell.Namespace($DownloadZipFile).Items()
$ExtractShell.NameSpace($Sys32).CopyHere($ExtractFiles) 
Start-Process $Sys32
$source = 'https://raw.githubusercontent.com/Neo23x0/sysmon-config/master/sysmonconfig-export.xml'
$destination = $Sys32 + $(Split-Path -Path $source -Leaf)
Invoke-RestMethod -Uri $source -OutFile $destination
Sysmon64.exe -i $destination
