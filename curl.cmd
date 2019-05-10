curl \
  -X POST \
  http://localhost:8282/posts \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTc0OTAxMjIsImlkIjoiYmo4bjRjb3Q4NzQyb21uZWpqZDAifQ.hlYfQ-u-5ATeYLe8xy9lDQaPbEuVb-0iI5qPRJj1ALU" \
  -H "Content-Type: application/json" \
  -d '{"title":"58465b4ea6fe886d3215c6df","content":"hello"}'

curl \
  -X GET \
  http://localhost:8282/posts \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTc0OTAxMjIsImlkIjoiYmo4bjRjb3Q4NzQyb21uZWpqZDAifQ.hlYfQ-u-5ATeYLe8xy9lDQaPbEuVb-0iI5qPRJj1ALU"

curl \
  -X GET \
  http://localhost:8282/posts/bj8nck0t8743105asvk0 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTc0OTAxMjIsImlkIjoiYmo4bjRjb3Q4NzQyb21uZWpqZDAifQ.hlYfQ-u-5ATeYLe8xy9lDQaPbEuVb-0iI5qPRJj1ALU"


curl \
  -X POST \
  http://localhost:8282/devices \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTc1NzMzMDcsImlkIjoiYmo4bjRjb3Q4NzQyb21uZWpqZDAifQ.F6s4u409GmYSlaChJWlPiy_rCXZrxYPiECBA-0THKXc" \
  -H "Content-Type: application/json" \
  -d '{"name" : "client2",
    "ipv4" : "192.168.0.1",
    "vnc_port" : "5000",
    "apps" : [{"name":"testApp","version":"1.0.0.1"}]
}'

curl \
  -X GET \
  http://localhost:8282/posts/bj8nck0t8743105asvk0 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTc1NzMzMDcsImlkIjoiYmo4bjRjb3Q4NzQyb21uZWpqZDAifQ.F6s4u409GmYSlaChJWlPiy_rCXZrxYPiECBA-0THKXc"


curl \
  -X POST \
  http://localhost8282/login \
  -H "Content-Type: application/json" \
  -d '{"email":"jon@labstack.com","password":"shhh!"}'

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTc1NzMzMDcsImlkIjoiYmo4bjRjb3Q4NzQyb21uZWpqZDAifQ.F6s4u409GmYSlaChJWlPiy_rCXZrxYPiECBA-0THKXc
