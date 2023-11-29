# golang-jwt-project

<h3>How to run?</h3>

``docker compose up --build``

That's it!

<h3>Functionalities</h3>

 POST   /users/signup

 POST   /users/login 

 GET    /users       

 GET    /users/:id   

<h3>Test</h3>
Use following curl commands to see demo:

Register an ADMIN user:
``curl -X POST -H "Content-Type: application/json"  -d '{
"first_name":"Shubham",
"last_name":"ADMIN",
"password":"shubham_password",
"email":"shubham.a@email.com", 
"phone":"+918979524541", 
"user_type":"ADMIN"}' http://localhost:8000/users/signup``

Register a USER(Normal) user with same phone number, throws error {"error":"userId with email or phone already exists"}:
``curl -X POST -H "Content-Type: application/json"  -d '{
"first_name":"Shubham",
"last_name":"USER",
"password":"shubham_password",
"email":"shubham.u@email.com",
"phone":"+918979524541",
"user_type":"USER"}' http://localhost:8000/users/signup``

Register a USER(Normal) user with new details:
``curl -X POST -H "Content-Type: application/json"  -d '{
"first_name":"Shubham",
"last_name":"USER",
"password":"shubham_password",
"email":"shubham.u@email.com",
"phone":"+918979524542",
"user_type":"USER"}' http://localhost:8000/users/signup``

Login as Admin User:
``curl -X POST -H "Content-Type: application/json"  -d '{
"password":"shubham_password",
"email":"shubham.u@email.com"}' http://localhost:8000/users/login``

Use token in login to access resource(Get Users):
`` curl -X GET -H "Content-Type: application/json" -H "token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InNodWJoYW0udUBlbWFpbC5jb20iLCJGaXJzdE5hbWUiOiJTaHViaGFtIiwiTGFzdE5hbWUiOiJVU0VSIiwiVWlkIjoiNjU2NmQ2NGY3OTIwNjM3ODg5MTMzODI0IiwiVXNlclR5cGUiOiJVU0VSIiwiZXhwIjoxNzAxMjM4NjUxfQ.GVALG4N_qU_WSsGIF51sZqTK3N68Ai8u8-umfE0WKOM" http://localhost:8000/users ``

(Here token is received from login response, and only admins can see all users)

GetUser/:id:

`` curl -X GET -H "Content-Type: application/json" -H "token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InNodWJoYW0udUBlbWFpbC5jb20iLCJGaXJzdE5hbWUiOiJTaHViaGFtIiwiTGFzdE5hbWUiOiJVU0VSIiwiVWlkIjoiNjU2NmQ2NGY3OTIwNjM3ODg5MTMzODI0IiwiVXNlclR5cGUiOiJVU0VSIiwiZXhwIjoxNzAxMjM4NjUxfQ.GVALG4N_qU_WSsGIF51sZqTK3N68Ai8u8-umfE0WKOM" http://localhost:8000/users/6566d64f7920637889133824 ``

<h3> Future aspects </h3>

Will add more tests soon.
Have a kit of smart-2FA through AI which can be integrated(Human typing based behavioral authentication).
