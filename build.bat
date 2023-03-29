@echo off

set WKDIR=%~dp0

echo.
cd /d %WKDIR%/frontend
CALL build.bat

echo.
cd /d %WKDIR%/factory
CALL build.bat

cd /d %WKDIR%
cmd /k
