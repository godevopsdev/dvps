IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = 'amosal@lallemand.com')
    CREATE USER [amosal@lallemand.com] FROM LOGIN [amosal@lallemand.com];

IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = 'dkalmer@lallemand.com')
    CREATE USER [dkalmer@lallemand.com] FROM LOGIN [dkalmer@lallemand.com];

IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = 'fgagne@lallemand.com')
    CREATE USER [fgagne@lallemand.com] FROM LOGIN [fgagne@lallemand.com];

IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = 'gabidoye@lallemand.com')
    CREATE USER [gabidoye@lallemand.com] FROM LOGIN [gabidoye@lallemand.com];

IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = 'garsenault@lallemand.com')
    CREATE USER [garsenault@lallemand.com] FROM LOGIN [garsenault@lallemand.com];

IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = 'mlajeunesse@lallemand.com')
    CREATE USER [mlajeunesse@lallemand.com] FROM LOGIN [mlajeunesse@lallemand.com];

IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = 'pbeltran@lallemand.com')
    CREATE USER [pbeltran@lallemand.com] FROM LOGIN [pbeltran@lallemand.com];

IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = 'snestruck@lallemand.com')
    CREATE USER [snestruck@lallemand.com] FROM LOGIN [snestruck@lallemand.com];

-- IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = 'svc_ssis_etl')
--     CREATE USER svc_ssis_etl FROM LOGIN svc_ssis_etl;

IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = 'avanade_cvelasco@lallemand.com')
    CREATE USER [avanade_cvelasco@lallemand.com] FROM LOGIN [avanade_cvelasco@lallemand.com];

IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = 'avanade_averhoef@lallemand.com')
    CREATE USER [avanade_averhoef@lallemand.com] FROM LOGIN [avanade_averhoef@lallemand.com];

IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = 'accenture_pgadekar@lallemand.com')
    CREATE USER [accenture_pgadekar@lallemand.com] FROM LOGIN [accenture_pgadekar@lallemand.com];

IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = 'accenture_ssundar@lallemand.com')
    CREATE USER [accenture_ssundar@lallemand.com] FROM LOGIN [accenture_ssundar@lallemand.com];

IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = 'accenture_smallick@lallemand.com')
    CREATE USER [accenture_smallick@lallemand.com] FROM LOGIN [accenture_smallick@lallemand.com];

-- ALTER ROLE db_datareader ADD MEMBER svc_ssis_etl;
-- ALTER ROLE db_datawriter ADD MEMBER svc_ssis_etl;
-- ALTER ROLE db_ddladmin ADD MEMBER svc_ssis_etl;
-- ALTER ROLE db_owner ADD MEMBER svc_ssis_etl;

ALTER ROLE db_owner ADD MEMBER [amosal@lallemand.com];
ALTER ROLE db_owner ADD MEMBER [dkalmer@lallemand.com];
ALTER ROLE db_owner ADD MEMBER [fgagne@lallemand.com];
ALTER ROLE db_owner ADD MEMBER [gabidoye@lallemand.com];
ALTER ROLE db_owner ADD MEMBER [garsenault@lallemand.com];
ALTER ROLE db_owner ADD MEMBER [mlajeunesse@lallemand.com];
ALTER ROLE db_owner ADD MEMBER [pbeltran@lallemand.com];
ALTER ROLE db_owner ADD MEMBER [snestruck@lallemand.com];

ALTER ROLE db_owner ADD MEMBER [avanade_cvelasco@lallemand.com];
ALTER ROLE db_owner ADD MEMBER [avanade_averhoef@lallemand.com];
ALTER ROLE db_owner ADD MEMBER [accenture_pgadekar@lallemand.com];
ALTER ROLE db_owner ADD MEMBER [accenture_smallick@lallemand.com];
ALTER ROLE db_owner ADD MEMBER [accenture_ssundar@lallemand.com];
