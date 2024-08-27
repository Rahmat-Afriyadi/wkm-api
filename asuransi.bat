@echo off


:loop
echo Running Go program...
go run server.go
if %ERRORLEVEL% NEQ 0 (
    echo Program crashed with exit code %ERRORLEVEL%. Restarting...
    timeout /t 2 >nul
    goto loop
) else (
    echo Program exited normally.
)


pause