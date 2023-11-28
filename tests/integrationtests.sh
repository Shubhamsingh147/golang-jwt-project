curl -X POST -H "Content-Type: application/json"  -d '{
            "first_name":"Shubham",
            "last_name":"Singh",
            "password":"shubham_password",
            "email":"shubham.s@email.com",
            "phone":"+918979524541",
            "user_type":"ADMIN"}'
http://localhost:8000/users/signup

#TODO: more testing scenarios