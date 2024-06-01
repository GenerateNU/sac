# Configuration

SAC uses environment variables for configuration. SAC will, by default, pull the envvars from the process in which it was launched, but you can also pass in a .env file (using the --config flag) and use the values defined in there. See .env.template for an example .env file.

| Name                                 | Group      | Description                           |
|--------------------------------------|------------|---------------------------------------|
| SAC_APPLICATION_PORT                 | app        | port to run the server on.            |
| SAC_APPLICATION_HOST                 | app        | host to run the server on.            |
| SAC_APPLICATION_BASE_URL             | app        | base url to run the server on.        |
| SAC_DB_USERNAME                      | db         | username for database.                |
| SAC_DB_PASSWORD                      | db         | password for database.                |
| SAC_DB_PORT                          | db         | port for database.                    |
| SAC_DB_HOST                          | db         | host for database.                    |
| SAC_DB_NAME                          | db         | name of database.                     |
| SAC_DB_REQUIRE_SSL                   | db         | if the db connection requires ssl.    |
| SAC_REDIS_ACTIVE_TOKENS_USERNAME     | redis      | username for active tokens redis.     |
| SAC_REDIS_ACTIVE_TOKENS_PASSWORD     | redis      | password for active tokens redis.     |
| SAC_REDIS_ACTIVE_TOKENS_HOST         | redis      | host for active tokens redis.         |
| SAC_REDIS_ACTIVE_TOKENS_PORT         | redis      | port for active tokens redis.         |
| SAC_REDIS_ACTIVE_TOKENS_DB           | redis      | db for active tokens redis.           |
| SAC_REDIS_BLACKLIST_USERNAME         | redis      | username for blacklist redis.         |
| SAC_REDIS_BLACKLIST_PASSWORD         | redis      | password for blacklist redis.         |
| SAC_REDIS_BLACKLIST_HOST             | redis      | host for blacklist redis.             |
| SAC_REDIS_BLACKLIST_PORT             | redis      | port for blacklist redis.             |
| SAC_REDIS_BLACKLIST_DB               | redis      | db for blacklist redis.               |
| SAC_REDIS_LIMITER_USERNAME           | redis      | username for limiter redis.           |
| SAC_REDIS_LIMITER_PASSWORD           | redis      | password for limiter redis.           |
| SAC_REDIS_LIMITER_HOST               | redis      | host for limiter redis.               |
| SAC_REDIS_LIMITER_PORT               | redis      | port for limiter redis.               |
| SAC_REDIS_LIMITER_DB                 | redis      | db for limiter redis.                 |
| SAC_SUDO_PASSWORD                    | superuser  | password for the superuser.           |
| SAC_AUTH_ACCESS_KEY                  | auth       | access key for auth.                  |
| SAC_AUTH_REFRESH_KEY                 | auth       | refresh key for auth.                 |
| SAC_AWS_BUCKET_NAME                  | aws        | bucket name for aws s3.               |
| SAC_AWS_ID                           | aws        | id for aws s3.                        |
| SAC_AWS_SECRET                       | aws        | secret for aws s3.                    |
| SAC_AWS_REGION                       | aws        | region for aws s3.                    |
| SAC_RESEND_API_KEY                   | resend     | api key for resend.                   |
| SAC_CALENDAR_MAX_TERMINATION_DATE    | calendar   | max termination date for calendar integrations. |

# Deployment design
SAC