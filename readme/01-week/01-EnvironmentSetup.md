# Environment Setup for Windows (WSL2 + Docker Desktop)
1. Install WSL2 (Windows Subsystem for Linux)

WSL2 is required for Docker to run on Windows.

    Open PowerShell as Administrator and run:
```
wsl --install
```
This installs Ubuntu by default.

Restart your PC after installation.

To check if WSL2 is installed properly, run:
```
 wsl --list --verbose
```
You should see:
```
NAME                   STATE           VERSION
Ubuntu-22.04           Running         2
```
If the version is 1, upgrade it:

    wsl --set-version Ubuntu-22.04 2

2. Install Docker Desktop

    Download Docker Desktop from:
    👉[ Docker Desktop for Windows](https://docs.docker.com/desktop/setup/install/windows-install/)

    Install it and enable WSL2 integration during setup.

    Restart your PC after installation.

    Open PowerShell and run:
```
docker --version
```
If Docker is installed correctly, it will show the version.
