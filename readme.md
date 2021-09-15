This is rest api CRUD service. <br>
The service has 4 APIs, first api is CreateUser api.<br>
It accepts the request:<br>
{<br>
    "first_name": "Joe",<br>
    "last_name": "Parker",<br>
    "age": 29,<br>
    "mail":"joe@gmail.com"<br>
}<br>
and returns response:<br>
{<br>
    "msg":"user created"<br>
}<br>
Second api is UpgradeUserByUUID accepts:<br>
{<br>
    "id": 1,<br>
    "first_name": "Joe",<br>
    "last_name": "Parker",<br>
    "age": 29,<br>
    "mail":"joe@gmail.com"<br>
}<br>
response is:<br>
{<br>
    "msg": "user is upgraded"<br>
}<br>
The last api is GetUserByUUID, request object of this api is:<br>
{<br>
    "id": 1<br>
}<br>
and the response is:<br>
{<br>
    "id": 1,<br>
    "first_name": "Joe",<br>
    "last_name": "Parker",<br>
    "age": 29,<br>
    "mail":"joe@gmail.com"<br>
}<br>
