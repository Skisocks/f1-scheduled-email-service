# formula-1-info/email-service-go
This is an automated email service that provides start time information to subscribers on the day of qualifying 
and the race. 

The service uses api-sport.io to retrieve data regarding current F1 events and is currently configured
to use a Postgres database for accessing the current subscribers' information. 
Docker is used in deployment running a cron job to start the service at 9:00 every day.

## Config
The project uses a config.yaml file located in the root of the project to set global variables. An example of this 
file can be found in the repository.

## License
[MIT](https://choosealicense.com/licenses/mit/)
