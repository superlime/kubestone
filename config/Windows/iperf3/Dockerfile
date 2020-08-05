FROM mcr.microsoft.com/windows/servercore:1809

RUN powershell -Command New-Item -Path "C:\usr\bin" -ItemType Directory

WORKDIR C:/usr/bin/

COPY * C:/usr/bin/

ENTRYPOINT ["./iperf3"]