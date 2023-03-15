@echo off

set WORK_HOME=%~dp0

for /R %WORK_HOME% %%f in (*.proto) do (
    echo src:%%f
    protoc --go_out=%%~dpf -I %%~dpf %%~nxf
)

pause
