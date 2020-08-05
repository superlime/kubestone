FROM mcr.microsoft.com/windows/servercore:1809

RUN powershell -Command New-Item -Path "C:\tmp" -ItemType Directory

RUN powershell -Command New-Item -Path "C:\run" -ItemType Directory

COPY * C:/run/

ENTRYPOINT ["/run/ntttcp"]