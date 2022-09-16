A "work-in-progress" mini-project to demonstrate user registration, login, logout, and performing simple DB queries - all that using:

- Golang Gin framework
- Custom middlewares for Logs(in JSON format using 'logrus'), CORS, and Authentication/Authorization
- Golang 'Crypto' library to hash passwords before saving them to the DB
- MySql as persistent database
- JWTs for user authentication/authorization
- Golang Testify library for unit tests
- HTML, CSS and JS for responsive UI

It also features:

- AJAX-style view update that doesn't require page reloads for every record that is added/updated/deleted from the DB.
- Microservices architecture to support CI/CD and is added to gitlab CI/CD pipeline

Heads Up: simple grid animation is applied at login/registration page that chrome doesn't support, but firefox does
