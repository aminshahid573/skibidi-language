@echo off
setlocal

echo Installing SkibidiLang for Windows...

REM Target directory
set TARGET_DIR=%USERPROFILE%\SkibidiLang

REM Create directory
mkdir "%TARGET_DIR%"

REM Move the executable
copy /Y "skibidi.exe" "%TARGET_DIR%\skibidi.exe"

REM Add to PATH via registry
reg query "HKCU\Environment" | findstr /C:"Path" >nul
if %errorlevel%==0 (
    for /f "tokens=2*" %%A in ('reg query "HKCU\Environment" /v Path') do (
        set "PATH_VALUE=%%B"
    )
) else (
    set "PATH_VALUE="
)

echo Adding SkibidiLang to PATH...
setx Path "%PATH_VALUE%;%TARGET_DIR%"

echo Done! You may need to restart your terminal or log out and back in.

pause
