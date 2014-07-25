var dewisApp = angular.module('dewis', [])
  .controller('login', ['$scope', '$http', function($scope, $http){
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
  }]);