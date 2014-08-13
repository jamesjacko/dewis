module.exports = function($scope, $parse, GetData){
  $scope.data = {
    status: true,
    data: {
      project1: {
        name: "Project1",
        id: "123"
      },
      project2: {
        name: "Project2",
        id: "456"
      }
    }
  };
  $scope.goto = function(id){
    $location.path("project/" + id);
  }
  console.log($scope.data);
}