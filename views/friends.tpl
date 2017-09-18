<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>好友</title>
{{template "inc/meta.tpl" .}}
</head>
<body>
<div class="container">
  <form class="form-signin" id="friends-form">
    <label for="userid" class="sr-only">userid</label>
    <input type="tel" name="userid" class="form-control" placeholder="userid" required autofocus>
    <button class="btn btn-lg btn-primary btn-block" type="submit">查看我的好友</button>
  </form>
</div>
{{template "inc/foot.tpl" .}}
</body>
</html>
