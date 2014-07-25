var dewisApp = angular.module('dewis', ['ngRoute'])
  .controller('loginCtrl', ['$scope', '$http', function($scope, $http){
    $scope.data = {
      status: 'Hello World'
    };
    $scope.login = function(){
      console.log($scope.data);
      var dataString = "";
      for (var key in $scope.data) {
          if (dataString != "") {
              dataString += "&";
          }
          dataString += key + "=" + $scope.data[key];
      }
      console.log(dataString);
      $http({
        method: 'post',
        url: '/test.go',
        data: dataString,
        headers: {'Content-Type': 'application/x-www-form-urlencoded'}
      }).success(function(data){
        console.log(data);
      });
    }
  }])
  .controller('timelineCtrl', ['$scope', '$http', function($scope, $http){
    $scope.json = {
      status: "ok",
      curUser: "James",
      data: {
          messages : {
            message1 : {
              user: "James",
              content: "Wow it works"
            },
            message2 : {
              user: "Thiago",
              content: "Yes it does, Hello World!"
            },
            message3 : {
              user: "Liam",
              content: "This is cool!"
            }
          }
        }
    };
  }])
  .config(['$routeProvider', function($routeProvider){
    $routeProvider
      .when("/timeline", {templateUrl: "partials/timeline.html", controller: "timelineCtrl"})
      .when("/login", {templateUrl: "partials/login.html", controller: "loginCtrl"})
      .otherwise({redirectTo: '/'});
  }]);