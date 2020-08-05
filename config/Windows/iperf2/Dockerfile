FROM mcr.microsoft.com/windows/servercore:1809

RUN powershell -Command New-Item -Path "C:\tmp" -ItemType Directory

COPY * C:/

ENTRYPOINT ["iperf.bat"]