<!DOCTYPE html>
<html lang="en">
<head>
  <title>LAM</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
<center>
    <h4>Update user's information</h4><br>
    <form action="/admin/create" method="post">
    Account:<input type="text" name="Username" value="{{.Username}}"/><br>
    Password:<input type="password" name="Password" value="{{.Password}}"/><br>
    Role:<input type="text" name="Role" value="{{.Role}}"/><br>
    Email:<input type="text" name="Email" value="{{.Email}}"/><br>
    <button type="submit">submit</button>
    </form>
</center>
</body>
</html>