@echo off

cd /d %~dp0

set WORK_HOME=%~dp0
setx /M PATH "%WORK_HOME%;%Path%"
setx PATH "%WORK_HOME%;%Path%"

pause
