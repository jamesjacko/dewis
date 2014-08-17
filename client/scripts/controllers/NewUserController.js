module.exports = function($scope, $http, GetData){
  $scope.data = {};
  $scope.data.Isadmin = "false";
  $scope.processUser = function(){
    console.log($scope.data);
    $scope.data.Isadmin = $scope.data.Isadmin.toString();
    $scope.json = GetData.fetch(
      'User', 
      'addUser', 
      $scope.data
    ).success(function(dataReturned){
        console.log(dataReturned);
        $scope.json = dataReturned;
    });
  }
}
