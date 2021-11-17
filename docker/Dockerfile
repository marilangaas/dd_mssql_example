FROM mcr.microsoft.com/mssql/server:2019-latest
ENV MSSQL_SA_PASSWORD="myPassw0rd"
ENV ACCEPT_EULA="Y"
COPY ddExample.bak /var/opt/mssql/backup/mssql_dd.bak
COPY restore.sh /
USER root
RUN chmod +x /restore.sh
RUN /opt/mssql/bin/sqlservr & /restore.sh



