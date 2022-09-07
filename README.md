A "work-in-progress" mini-project to demonstrate user registrations, logins, logouts, and performing simple DB queries - all that using:

- Golang Gin framework
- Custom middlewares for Logs(in JSON format using 'logrus'), CORS, and Authentication/Authorization
- Golang 'Crypto' library to hash passwords before saving them to the DB
- MySql as persistent database
- Uses Crypto library to hash passwords before saving them to the database
- JWTs for user authentication/authorization
- HTML, CSS and JS for responsive UI

It also features ajax-style view update that doesn't require page reloads for every record that is added/updated/deleted from the DB.

Next up:

- Add tests to enable CI
- Seperate application/handlers from DB using service layer

Heads Up: simple grid animation is applied at login/registration page that chrome doesn't support, but firefox does
