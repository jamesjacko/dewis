module.exports = function($scope, $parse, GetData){
  $scope.submit = function(){
    var ret = "";
    console.log(angular.fromJson($scope.DataObj));
    GetData.fetch(
      $scope.Request, 
      $scope.Action, 
      angular.fromJson($scope.DataObj)
    ).success(function(dataReturned){
        console.log(dataReturned);
        console.log($scope);
        $scope.ReturnedData = angular.toJson(dataReturned);
        //alert(angular.toJson(dataReturned));
    });
    
  }
}