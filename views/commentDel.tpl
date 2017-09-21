<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>删除评论</title>
{{template "inc/meta.tpl" .}}
</head>
<body>
<div class="container">
  <form class="form-signin" id="commentDel-form">
    <h2 class="form-signin-heading">删除评论~</h2>

    <label for="inputUsername" class="sr-only">Username</label>
    <input type="text" name="commentID" class="form-control" placeholder="你想删的评论ID~" required autofocus>

    <div class="checkbox">
    </div>
    <button class="btn btn-lg btn-primary btn-block" type="submit">Delete</button>
  </form>
</div>
{{template "inc/foot.tpl" .}}
</body>
</html>
