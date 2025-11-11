@echo off
echo ========================================
echo Building ATM-Go for Windows
echo ========================================

REM Create bin directory if it doesn't exist
if not exist "..\bin" mkdir "..\bin"

REM Build for Windows
echo Building atm.exe...
cd ..
go build -ldflags="-s -w" -o bin\atm.exe cmd\atm\main.go

if %ERRORLEVEL% == 0 (
    echo.
    echo ========================================
    echo Build successful!
    echo ========================================
    echo.
    echo Executable created at: bin\atm.exe
    echo.
    echo To install globally, run:
    echo   copy bin\atm.exe C:\Windows\System32\
    echo.
    echo Or add the bin directory to your PATH.
    echo ========================================
) else (
    echo.
    echo ========================================
    echo Build failed!
    echo ========================================
    exit /b 1
)
