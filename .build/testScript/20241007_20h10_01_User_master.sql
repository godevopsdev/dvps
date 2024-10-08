USE MASTER;
IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'amosal@lallemand.com')
    CREATE LOGIN [amosal@lallemand.com] FROM EXTERNAL PROVIDER;

IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'dkalmer@lallemand.com')
    CREATE LOGIN [dkalmer@lallemand.com] FROM EXTERNAL PROVIDER;

IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'fgagne@lallemand.com')
    CREATE LOGIN [fgagne@lallemand.com] FROM EXTERNAL PROVIDER;

IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'gabidoye@lallemand.com')
    CREATE LOGIN [gabidoye@lallemand.com] FROM EXTERNAL PROVIDER;

IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'garsenault@lallemand.com')
    CREATE LOGIN [garsenault@lallemand.com] FROM EXTERNAL PROVIDER;

IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'mlajeunesse@lallemand.com')
    CREATE LOGIN [mlajeunesse@lallemand.com] FROM EXTERNAL PROVIDER;

IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'pbeltran@lallemand.com')
    CREATE LOGIN [pbeltran@lallemand.com] FROM EXTERNAL PROVIDER;

IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'snestruck@lallemand.com')
    CREATE LOGIN [snestruck@lallemand.com] FROM EXTERNAL PROVIDER;

IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'avanade_cvelasco@lallemand.com')
    CREATE LOGIN [avanade_cvelasco@lallemand.com] FROM EXTERNAL PROVIDER;

IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'avanade_averhoef@lallemand.com')
    CREATE LOGIN [avanade_averhoef@lallemand.com] FROM EXTERNAL PROVIDER;

IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'accenture_pgadekar@lallemand.com')
    CREATE LOGIN [accenture_pgadekar@lallemand.com] FROM EXTERNAL PROVIDER;

IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'accenture_smallick@lallemand.com')
    CREATE LOGIN [accenture_smallick@lallemand.com] FROM EXTERNAL PROVIDER;

IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = N'accenture_ssundar@lallemand.com')
    CREATE LOGIN [accenture_ssundar@lallemand.com] FROM EXTERNAL PROVIDER;

