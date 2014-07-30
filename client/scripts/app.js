var dewisApp = angular.module('dewis', ['ngRoute'])
  .controller('loginCtrl', ['$scope', '$http', function($scope, $http){
    $scope.data = {
      request: 'test',
      status: 'HelloWorld'
    };
    $scope.login = function(){
//      console.log($scope.data);
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
        headers: {'Content-Type': 'application/x-www-form-urlencoded'}
      }).success(function(data){
        console.log(data.Data);
      }).error(function(data, status, headers, config){
        console.log(data, status, headers, config);
      });
    }
  }])
  .controller('timelineCtrl', ['$scope', '$http', function($scope, $http){
    Object.size = function(obj) {
      var size = 0, key;
      for (key in obj) {
        if (typeof obj[key] == 'object') size++;
        console.log(key);
      }
      return size;
    };
    $scope.json = {
      status: "ok",
      curUser: "James",
      data: {
          messages : {
            message1 : {
              user: "James",
              content: "Wow it works",
              time: "1406717220000"
            },
            message2 : {
              user: "Thiago",
              content: "Yes it does, Hello World!",
              time: "1406717220000"
            },
            message3 : {
              user: "Liam",
              content: "This is cool!",
              time: "1406716967000"
            },
            message4 : {
              user: "James",
              content: "Wow it works",
              time: "1406717220000"
            },
            message5 : {
              user: "Thiago",
              content: "Yes it does, Hello World!",
              time: "1406717220000"
            },
            message6 : {
              user: "James",
              content: "This is cool!",
              time: "1406716967000"
            }
          }
        }
    };
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
    $scope.addMessage = function(){
      var newName = "message" + Object.size($scope.json.data.messages) + 1;
      $scope.json.data.messages[newName] = { };
      $scope.json.data.messages[newName].user = $scope.json.curUser;
      $scope.json.data.messages[newName].content = $scope.newMessage;
      $scope.json.data.messages[newName].time = new Date().getTime();
      $scope.newMessage = "";
      console.log($scope.json.data.messages);
    };
    $scope.getMoreMessages = function(){
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
