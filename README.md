A "work-in-progress" mini-project to demonstrate user registrations, logins, logouts, and performing simple DB queries - all that using:

- Golang Gin framework
- Custom middlewares for Logs(in JSON format using 'logrus'), CORS, and Authentication/Authorization
- MySql as persistent database
- Uses Crypto library to hash passwords before saving them to the database
- JWTs for user authentication/authorization
- HTML, CSS and JS for responsive UI

Next up:

- Add tests to enable CI
- Seperate application/handlers from DB using service layer

Heads Up: simple grid animation is applied at login/registration page that chrome doesn't support, but firefox does
