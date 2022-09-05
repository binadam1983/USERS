A "work-in-progress" mini-project to demonstrate user registrations, logins, logouts, and simple DB queries to pull data records which uses the following:

- Golang Gin framework
- Custom middlewares for Logs(in JSON format using 'logrus'), CORS, and Authentication/Authorization
- MySql as persistent database
- Uses Crypto library to hash passwords before saving them to the database
- JWTs for user authentication/authorization
- HTML, CSS and JS for responsive UI

Next up:

- add tests
- Seperation of application/handlers from DB using service layer
