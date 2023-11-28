# golang-jwt-project

<h3>How to run?</h3>

``docker compose up --build``

That's it!

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

