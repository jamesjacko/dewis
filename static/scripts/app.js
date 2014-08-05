(function e(t,n,r){function s(o,u){if(!n[o]){if(!t[o]){var a=typeof require=="function"&&require;if(!u&&a)return a(o,!0);if(i)return i(o,!0);throw new Error("Cannot find module '"+o+"'")}var f=n[o]={exports:{}};t[o][0].call(f.exports,function(e){var n=t[o][1][e];return s(n?n:e)},f,f.exports,e,t,n,r)}return n[o].exports}var i=typeof require=="function"&&require;for(var o=0;o<r.length;o++)s(r[o]);return s})({1:[function(require,module,exports){
angular.module('dewis', ['ngRoute'])


  .config(['$routeProvider', function($routeProvider){
    $routeProvider
      .when("/timeline", {templateUrl: "partials/timeline.html", controller: "TimelineController"})
      .when("/login", {templateUrl: "partials/login.html", controller: "LoginController"})
      .otherwise({redirectTo: '/login'});
  }])


  .controller({
    TimelineController: require('./controllers/TimelineController'),
    LoginController: require('./controllers/LoginController')
  })


  .directive({
    ngEnter: require('./directives/ngEnter'),
    scroller: require('./directives/scroller')
  })

  
  .factory({
    GetData: require('./services/GetData')
  });
},{"./controllers/LoginController":2,"./controllers/TimelineController":3,"./directives/ngEnter":4,"./directives/scroller":5,"./services/GetData":6}],2:[function(require,module,exports){
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
},{}],3:[function(require,module,exports){
module.exports = function($scope, $http, GetData){
  $scope.curUser = "53dbafcff66029358cd113a1";
  Object.size = function(obj) {
    var size = 0, key;
    for (key in obj) {
      if (typeof obj[key] == 'object') size++;
      console.log(key);
    }
    return size;
  };



  $scope.json = GetData.fetch(
    'Timeline', 
    'GetRecords', 
    {
      Quantity: "0"
    }
  ).success(function(dataReturned){
      $scope.json = dataReturned;
    });

  if($scope.json == null){
    $scope.json = {
      Status: true,
      Records: Array()
    }
  }
  console.log($scope.json);

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

    $scope.added = GetData.fetch(
      'Timeline', 
      'AddRecord', 
      {
        User: $scope.curUser,
        Content: $scope.newRecord
      }
    ).success(function(data){
      var newName = "record" + Object.size($scope.json.Records) + 1;
      var newOne = {
        User: $scope.curUser,
        Content: $scope.newRecord,
        Time: new Date().getTime()
      }
      if($scope.json.Records == null) $scope.json.Records = Array();
      $scope.json.Records.push(newOne);
      $scope.newRecord = "";
    });
  };
  $scope.getMoreRecords = function(){
    alert("hello");
  };
};
},{}],4:[function(require,module,exports){
module.exports = function(){
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
}
},{}],5:[function(require,module,exports){
module.exports = function () {
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
};
},{}],6:[function(require,module,exports){
module.exports= function($http){
  
  function fetch(req, act, data){
    var json = null;

    var requestObj = {
      Request: req,
      Action: act,
      Data: data
    };

    //console.log(requestObj);
    return $http({
      method: 'post',
      url: '/api',
      data: requestObj,
      headers: {'Content-Type': 'application/json'}
    });
  }

  return{
    fetch: fetch
  }
}
},{}]},{},[1]);