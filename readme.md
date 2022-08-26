# Notifications Mini-system

This is a sample project for building a notification system with Go and Vuejs. It is intentionally 
trite for the purpose of not over-engineering from the start with complexity added as needed.

## Running the API

The API requires MySQL for storing notifications and user data. It relies on environment variables
to pass database connection settings. You will need to have the following variables configured.

```
export DB_User=<api_db_user>
export DB_Password=<api_db_password>
export DB_Name=notifications-minisys
```

```
go run api.go
```