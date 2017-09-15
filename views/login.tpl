<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>登录</title>
{{template "inc/meta.tpl" .}}
</head>
<body>
<div class="container">
  <form class="form-signin" id="login-form">
    <h2 class="form-signin-heading">请登录</h2>
    <label for="inputEmail" class="sr-only">Username</label>
    <input type="tel" name="name" class="form-control" placeholder="Username" required autofocus>
    <label for="inputPassword" class="sr-only">Password</label>
    <input type="password" name="password" class="form-control" placeholder="Password" required>
    <div class="checkbox">
      <label>
      <input type="checkbox" value="remember-me">
      Remember me </label>
    </div>
    <button class="btn btn-lg btn-primary btn-block" type="submit">登录</button>
    <button class="btn btn-sm btn-warning btn-block" type="button">找回密码</button>
    <button class="btn btn-sm btn-danger  btn-block" type="button">销号</button>
  </form>
</div>
{{template "inc/foot.tpl" .}}
</body>
</html>
