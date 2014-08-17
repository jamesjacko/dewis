module.exports = function($scope, $http, GetData){
  $scope.data = {
    Username: null,
    Password: null
  };
  $scope.login = function(){
    console.log($scope.data);
    $scope.json = GetData.fetch(
      'Login', 
      'Login', 
      $scope.data
    ).success(function(dataReturned){
        console.log(dataReturned);
        $scope.json = dataReturned;
    });



  }
}
