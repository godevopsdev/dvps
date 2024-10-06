# dvps

![Go](https://img.shields.io/badge/Go-1.23-blue)
![Azure DevOps](https://img.shields.io/badge/Azure%20DevOps-Extension-blue) (later) 

## Overview

**dvps** is a powerful and lightweight Go-based application designed to execute SQL scripts in a variety of environments. This tool is crafted to simplify database migrations, updates, or any other script execution tasks.

The tool **WILL** integrates seamlessly with **Azure DevOps**, allowing users to run SQL scripts directly from their pipelines. The extension will be available on the **Azure DevOps Marketplace**, making it easy to add this functionality to your CI/CD workflows.

## Features

- **SQL Script Execution**: Run SQL scripts against any supported database (PostgreSQL, SQL Server).
- **Database Connectivity**: Connect to databases using standard connection strings.
- **Environment Support**: Use environment variables to securely pass database credentials.
- **Pipeline Integration**: Integrate with Azure DevOps pipelines to automate SQL script execution.
- **Error Handling**: Comprehensive error handling and logging for script execution.
- **Extensible**: Easily customizable and extendable for specific use cases.

## Installation

### Go Installation

You can install the tool locally by running the following Go command:

```bash
go install github.com/godevopsdev/dvps@latest
