A Full-Stack, "Progressive Web App" to demonstrate user registration, login, logout, and performing simple DB queries; all that using:

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
- A graceful shutdown when CTRL-C pressed by the user

Heads Up: simple grid animation is applied at login/registration page that chrome doesn't support, but firefox does



## Installation


You would need to:

1. Clone the repo
2. Set up 'secret key' for JWTs as GO env variable
3. a) Have MySql DB running on your machine. b) Create DB 'users' and create a table inside it with the following schema:

![](images/usersTable.jpg)

Mark the 'email' column 'primary key' and you're good to go.




## Usage

1. Server runs on port '3333', localhost.
2. The app features 5 handlers: home(user), Register, login, logout, and users(protected with Auth Middleware)
3. Index page is at 'localhost:3333/user'
    - Index page provides Registration and Login forms.
    - Successful Registration requires: a) a unique email address, meaning that the email address is not already registered in the DB b) Email provided has the valid format
    - it also features a simple grid animation

    ![](images/reg.gif)

    - A successful registration (& login) will log you in automatically and will generate a cookie with the life span of 5 minutes.    

    ![](images/reg_successful.gif)

4. Once in, it provides you with 'Get Users' and 'logout' options. 'Get Users' is protected with Auth Middleware and will run certain checks on you. if successful, it will pull all users from the DB and render them on the screen. You get a hamburger menu for small screen with options sliding in from the left when clicked.

![](images/getusers.gif)

5. The rendering view updates without requiring page reloads

![](images/ajax.gif)

6. Clicking Logout logs the user out and removes the cookie and take him/her to the logout page.

![](images/logout.gif)

## Testing

From Power Shell, cd into 'controllers' folder and Run:
```

PS C:\Users\smc\go\src\github.com\users\controllers> go.exe test -timeout 30s -v


=== RUN   TestRegisterValidUsernamePassword
time="2022-09-25T23:02:34+05:00" level=info msg="200 map[Token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RpbmcxMjNAdGVzdGluZy5jb20iLCJleHAiOjE2NjQxMjkxMzR9.5EyopnEkgzYXedYLtQj9zrp8UgsA_X7oTQUKaLduEhY]"
[GIN] 2022/09/25 - 23:02:34 |?[97;42m 200 ?[0m|    363.7776ms |       192.0.2.1 |?[97;46m POST    ?[0m "/user/register"
--- PASS: TestRegisterValidUsernamePassword (0.37s)
=== RUN   TestRegisterUsernameAlreadyTaken
time="2022-09-25T23:02:34+05:00" level=info msg="email=testing123%40testing.com&password=testing123"
[GIN] 2022/09/25 - 23:02:35 |?[90;43m 400 ?[0m|    223.1475ms |       192.0.2.1 |?[97;46m POST    ?[0m "/user/register"
--- PASS: TestRegisterUsernameAlreadyTaken (0.23s)
=== RUN   TestRegisterUsernameInvalid
[GIN] 2022/09/25 - 23:02:35 |?[90;43m 400 ?[0m|            0s |       192.0.2.1 |?[97;46m POST    ?[0m "/user/register"
--- PASS: TestRegisterUsernameInvalid (0.00s)
=== RUN   TestLoginValidCredentials
time="2022-09-25T23:02:35+05:00" level=info msg="200 map[Token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RpbmcxMjNAdGVzdGluZy5jb20iLCJleHAiOjE2NjQxMjkxMzV9.y_C5FdiLhnEuS2GRPRhsqWtkJch_5bYxVPYZn9alIDg]"
time="2022-09-25T23:02:35+05:00" level=info msg=testing123@testing.com
[GIN] 2022/09/25 - 23:02:35 |?[97;42m 200 ?[0m|    423.2557ms |       192.0.2.1 |?[97;46m POST    ?[0m "/user/login"
--- PASS: TestLoginValidCredentials (0.43s)
=== RUN   TestLoginInvalidCredentials
time="2022-09-25T23:02:35+05:00" level=info msg="Incorrect credentials OR Account does not exist. if later, Please Register first"

[GIN] 2022/09/25 - 23:02:35 |?[90;43m 404 ?[0m|    204.4881ms |       192.0.2.1 |?[97;46m POST    ?[0m "/user/login"
--- PASS: TestLoginInvalidCredentials (0.21s)
=== RUN   TestLogout
time="2022-09-25T23:02:36+05:00" level=info msg="200 map[Token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RpbmcxMjNAdGVzdGluZy5jb20iLCJleHAiOjE2NjQxMjkxMzZ9.fdCUXz_RL7wLvNCknOYHrda6d85StqxNLjinsz1wabw]"
time="2022-09-25T23:02:36+05:00" level=info msg=testing123@testing.com
[GIN] 2022/09/25 - 23:02:36 |?[97;42m 200 ?[0m|    400.2575ms |       192.0.2.1 |?[97;46m POST    ?[0m "/user/login"
[GIN] 2022/09/25 - 23:02:36 |?[90;43m 404 ?[0m|            0s |       192.0.2.1 |?[97;44m GET     ?[0m "/user/logout"
--- PASS: TestLogout (0.41s)
=== RUN   TestGetUsers
time="2022-09-25T23:02:36+05:00" level=info msg="200 map[Token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RpbmcxMjNAdGVzdGluZy5jb20iLCJleHAiOjE2NjQxMjkxMzZ9.fdCUXz_RL7wLvNCknOYHrda6d85StqxNLjinsz1wabw]"
time="2022-09-25T23:02:36+05:00" level=info msg=testing123@testing.com
[GIN] 2022/09/25 - 23:02:36 |?[97;42m 200 ?[0m|     401.432ms |       192.0.2.1 |?[97;46m POST    ?[0m "/user/login"
[GIN] 2022/09/25 - 23:02:36 |?[90;47m 302 ?[0m|      5.0039ms |       192.0.2.1 |?[97;44m GET     ?[0m "/user/users"
--- PASS: TestGetUsers (0.41s)
=== RUN   TestHomepage
[GIN] 2022/09/25 - 23:02:36 |?[97;42m 200 ?[0m|            0s |       192.0.2.1 |?[97;44m GET     ?[0m "/user"
--- PASS: TestHomepage (0.00s)
PASS
ok      github.com/users/controllers    2.118s
PS C:\Users\smc\go\src\github.com\users\controllers>

```

## Progressive Web App Stats

![](images/pwa_ready.JPG)


## Your Thoughts:

Although I would have loved to, the app doesn't strictly follow any architecture. Check out my other repo,'The_Clean_Architecture_CRUD', that follow the SOLID principle. Other than that, you're most welcome to open an issue for any modification.

Fork is welcomed!
