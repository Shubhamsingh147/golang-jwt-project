# golang-jwt-project

<h3>How to run?</h3>

``docker compose up --build``

That's it!

<h3>Functionalities</h3>

 POST   /users/signup             --> golang-jwt-artica/routes.AuthRoutes.SignUp.func1 (2 handlers)
 POST   /users/login              --> golang-jwt-artica/routes.AuthRoutes.Login.func2 (2 handlers)
 GET    /users                    --> golang-jwt-artica/routes.UserRoutes.GetUsers.func2 (3 handlers)
 GET    /users/:id                --> golang-jwt-artica/routes.UserRoutes.GetUser.func3 (3 handlers)
 GET    /api-1                    --> main.main.func1 (3 handlers)
 GET    /api-2                    --> main.main.func2 (3 handlers)

api-1 is First factor auth
api-2 is Second factor auth

<h3>Test</h3>
Use following curl commands to see demo:

Register a user:
``curl -X POST -H "Content-Type: application/json"  -d '{
            "first_name":"Shubham",
            "last_name":"Singh",
            "password":"shubham_password",
            "email":"shubham.s@email.com", 
            "phone":"+918979524541", 
            "user_type":"ADMIN"}' 
http://localhost:8000/users/signup``

<h3> Future aspects </h3>

Will add all tests soon.