#!/usr/bin/env bash


for i in {1..50};
do
    /opt/mssql-tools/bin/sqlcmd -S localhost -U 'sa' -P "$MSSQL_SA_PASSWORD" -Q 'RESTORE DATABASE ddExample FROM DISK = "/var/opt/mssql/backup/mssql_dd.bak" WITH MOVE "ddExample" TO "/var/opt/mssql/data/mssql_dd.mdf", MOVE "ddExample_log" TO "/var/opt/mssql/data/ddExample_log.ldf"'
    if [ $? -eq 0 ]
    then
        echo "[DATABASE RESTORED]"
        break
    else
        echo "not ready yet..."
        sleep 1
    fi
done