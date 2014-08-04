var dewisApp = angular.module('dewis', ['ngRoute'])
  .controller('loginCtrl', ['$scope', '$http', function($scope, $http){
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
    }])
  .controller('timelineCtrl', ['$scope', '$http', function($scope, $http){
    $scope.curUser = "53dbafcff66029358cd113a1";
    Object.size = function(obj) {
      var size = 0, key;
      for (key in obj) {
        if (typeof obj[key] == 'object') size++;
        console.log(key);
      }
      return size;
    };

    $scope.data = {
      Request: 'Timeline',
      Action: 'GetRecords',
      Data: {
        Quantity: '0'
      }
    };


    $http({
      method: 'post',
      url: '/api',
      data: $scope.data,
      headers: {'Content-Type': 'application/json'}
    }).success(function(data){
      $scope.json = data;
      console.log(data);
    }).error(function(data, status, headers, config){
      console.log(status);
    });

    $scope.users = {
      status: "ok",
      data: {
        users : {
          James : {
            avatar : "https://fbcdn-profile-a.akamaihd.net/hprofile-ak-xap1/t1.0-1/p50x50/10447731_10152145074977541_5235408576063947935_n.jpg"
          },
          Thiago : {
            avatar : "https://fbcdn-profile-a.akamaihd.net/hprofile-ak-xap1/t1.0-1/p50x50/734141_396611893754134_474245279_n.jpg"
          },
          Liam : {
            avatar : "https://fbcdn-profile-a.akamaihd.net/hprofile-ak-xap1/t1.0-1/p50x50/1016905_10152172564666288_6498009010658710214_n.jpg"
          }
        }
      }
    }


    $scope.addRecord = function(){

      var record = {
        Request: "Timeline",
        Action: "AddRecord",
        Data: {
          User: $scope.curUser,
          Content: $scope.newRecord
        }
      }

      

      console.log(record);
      $http({
        method: 'post',
        url: '/api',
        data: record,
        headers: {'Content-Type': 'application/json'}
      }).success(function(data){
        //$scope.json = data;
        console.log(data);
      }).error(function(data, status, headers, config){
        console.log(status);
      });


      var newName = "record" + Object.size($scope.json.Records) + 1;
      var newOne = {
        User: $scope.curUser,
        Content: $scope.newRecord,
        Time: new Date().getTime()
      }
      if($scope.json.Records == null) $scope.json.Records = Array();
      $scope.json.Records.push(newOne);
      $scope.newRecord = "";
    };
    $scope.getMoreRecords = function(){
      alert("hello");
    };
  }])
  .config(['$routeProvider', function($routeProvider){
    $routeProvider
      .when("/timeline", {templateUrl: "partials/timeline.html", controller: "timelineCtrl"})
      .when("/login", {templateUrl: "partials/login.html", controller: "loginCtrl"})
      .otherwise({redirectTo: '/login'});
  }])
  .directive('ngEnter', function(){
    return function(scope, element, attrs) {
      element.bind("keydown keypress", function(event){
        if(event.which === 13){
          scope.$apply(function(){
            scope.$eval(attrs.ngEnter, {'event': event});
          });
          event.preventDefault();
        }
      });
    };
  })
  .directive('scroller', function () {
    return { 
      restrict: 'A',
      scope: {
          loadingMethod: "&"
      },
      link: function (scope, elem, attrs) {
        rawElement = elem[0];
        elem.bind('scroll', function () {
          if((rawElement.scrollTop + rawElement.offsetHeight+5) >= rawElement.scrollHeight){
            scope.$apply(scope.loadingMethod); 
          }
        });
      }
    };
  });