module.exports = function($scope, $http){
  $scope.data = {
    request: 'test',
    status: 'HelloWorld'
  };
  $scope.login = function(){
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
      url: '/api',
      data: dataString,
      headers: {'Content-Type': 'application/json'}
    }).success(function(data){
      console.log(data.Data);
    }).error(function(data, status, headers, config){
      console.log(data, status, headers, config);
    });



  }
}
