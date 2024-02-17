# Get the absolute path to the current script
$ScriptPath = (Get-Item -Path $MyInvocation.MyCommand.Path).DirectoryName

# Build the Go CLI tool
go build -o "sac-cli" "cli/main.go"

# Check if sac-cli is already installed
if (Test-Path -Path "$Env:USERPROFILE\AppData\Local\Programs\sac-cli\sac-cli.exe") {
    exit 1
}

# Copy the sac-cli executable to a directory in the user's PATH
$InstallPath = "$Env:USERPROFILE\AppData\Local\Programs\sac-cli"
if (-not (Test-Path -Path $InstallPath)) {
    New-Item -ItemType Directory -Path $InstallPath | Out-Null
}
Copy-Item -Path "sac-cli" -Destination "$InstallPath\sac-cli.exe" -Force

# Add the installation path to the user's PATH
$PathEnvVar = [System.Environment]::GetEnvironmentVariable("PATH", [System.EnvironmentVariableTarget]::User)
if (-not ($PathEnvVar -like "*$InstallPath*")) {
    [System.Environment]::SetEnvironmentVariable("PATH", "$InstallPath;$PathEnvVar", [System.EnvironmentVariableTarget]::User)
}

# Inform the user
Write-Host "Installation complete. You can now run 'sac-cli' from anywhere."
