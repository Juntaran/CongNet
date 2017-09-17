<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>注册</title>
{{template "inc/meta.tpl" .}}
</head>
<body>
<div class="container">
  <form class="form-signin" id="cancel-form">
    <h2 class="form-signin-heading">注销账号，如需注销请先登录</h2>

    <label for="inputEmail" class="sr-only">Email</label>
    <input type="tel" name="email" class="form-control" placeholder="Email" required autofocus>

    <label for="inputPassword" class="sr-only">Password</label>
    <input id="password"  type="password" name="password" class="form-control" placeholder="Password" required>

    <label for="inputUsername" class="sr-only">Username</label>
    <input type="tel" name="name" class="form-control" placeholder="Username" required autofocus>

    <div class="checkbox">

    </div>
    <button class="btn btn-lg btn-danger btn-block" type="submit">Cancel</button>
  </form>
</div>
{{template "inc/foot.tpl" .}}
</body>
</html>
