IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'AZ_IT_SQLLEO_CICD')
    CREATE LOGIN AZ_IT_SQLLEO_CICD FROM EXTERNAL PROVIDER;

IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'Az_SIT_IT_LEOINT2')
    CREATE LOGIN Az_SIT_IT_LEOINT2 FROM EXTERNAL PROVIDER;