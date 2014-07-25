<!DOCTYPE html>
<html ng-app="dewis">
  <head>
    <script src="http://ajax.googleapis.com/ajax/libs/angularjs/1.2.6/angular.min.js"></script>
    <script src="scripts/app.js"></script>
  </head>
  <body>
    <h1>Login</h1>
    <form method="post" ng-controller="login" ng-submit="login()">
        Username:<br />
        <input type="text" name="username" value="<?php echo $submitted_username; ?>" ng-model="data.username" />
        <br /><br />
        Password:<br />
        <input type="password" name="password" value="" ng-model="data.password" />
        <br /><br />
        <input type="submit" value="Login" />
    </form>
    <a href="register.php">Register</a>
  </body>
</html>