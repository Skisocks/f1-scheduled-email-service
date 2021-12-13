# formula-1-info/email-service-go
This is an automated email service that provides start time information to subscribers on the day of qualifying 
and the race. 

The service uses api-sport.io to retrieve data regarding current F1 events and is currently configured
to use a Postgres database for accessing the current subscribers' information. 
Docker is used in deployment running a cron job to start the service at 9:00 every day.

## Config
The project uses a "config.yaml" file to be located in the root of the project to set global variables. An example of this 
file can be found in the repository.

| Config Variable | Description | Type |
| ------ | ------ | ------ |
| BASE_URL | The common prefix of the API  | String |
| DRIVER_ENDPOINT | The endpoint of the Ergast API that provides driver related information | String |
| DRIVERS_STANDINGS_ENDPOINT | The endpoint of the Ergast API that provides the drivers championship positions | String |
| SEASON | The query parameter that defines which season the request is made for | String |
| CONSTRUCTORS_STANDINGS_ENDPOINT | The endpoint of the Ergast API that provides the constructors championship positions | String | 
| TIMEOUT | Defines the interval in seconds before the connection times out  | Int |
|  |  |  |
|  |  |  |
|  |  |  |
|  |  |  |
|  |  |  |
|  |  |  |
| SENDER_EMAIL_PASSWORD | Password of the sending SMTP server | |
| SENDER_EMAIL_ADDRESS | Email address of the sending SMTP server | |
| SENDER_EMAIL_PASSWORD | Password of the sending SMTP server | |

## License
[MIT](https://choosealicense.com/licenses/mit/)
