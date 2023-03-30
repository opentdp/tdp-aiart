@echo off

set WKDIR=%~dp0

start %WKDIR%/frontend/serve.bat

start %WKDIR%/factory/serve.bat
