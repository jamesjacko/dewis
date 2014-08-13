(function e(t,n,r){function s(o,u){if(!n[o]){if(!t[o]){var a=typeof require=="function"&&require;if(!u&&a)return a(o,!0);if(i)return i(o,!0);throw new Error("Cannot find module '"+o+"'")}var f=n[o]={exports:{}};t[o][0].call(f.exports,function(e){var n=t[o][1][e];return s(n?n:e)},f,f.exports,e,t,n,r)}return n[o].exports}var i=typeof require=="function"&&require;for(var o=0;o<r.length;o++)s(r[o]);return s})({1:[function(require,module,exports){
angular.module('dewis', ['ngRoute'])


  .config(['$routeProvider', function($routeProvider){
    $routeProvider
      .when("/login", {templateUrl: "partials/login.html", controller: "LoginController"})
      .when("/projects", {templateUrl: "partials/projects.html", controller: "ProjectsController"})
      .when("/project/:id", {templateUrl: "partials/single-project/dash.html", controller: "SingleProjectController"})
      .when("/project/:id/timeline", {templateUrl: "partials/timeline.html", controller: "TimelineController"})
      .when("/project/:id/files", {templateUrl: "partials/single-project/files.html", controller: "SingleProjectFilesController"})
      .when("/developer", {templateUrl: "partials/developer.html", controller: "DeveloperController"})
      .otherwise({redirectTo: '/login'});
  }])


  .controller({
    TimelineController: require('./controllers/TimelineController'),
    LoginController: require('./controllers/LoginController'),
    DeveloperController: require('./controllers/DeveloperController'),
    ProjectsController: require('./controllers/ProjectsController'),
    SingleProjectController: require('./controllers/single-project/ProjectController'),
    SingleProjectFilesController: require('./controllers/single-project/ProjectFilesController'),
  })


  .directive({
    ngEnter: require('./directives/ngEnter'),
    scroller: require('./directives/scroller'),
    tree: require('./directives/tree'),
    treeNode: require('./directives/treeNode')
  })

  
  .factory({
    GetData: require('./services/GetData')
  });
},{"./controllers/DeveloperController":2,"./controllers/LoginController":3,"./controllers/ProjectsController":4,"./controllers/TimelineController":5,"./controllers/single-project/ProjectController":6,"./controllers/single-project/ProjectFilesController":7,"./directives/ngEnter":8,"./directives/scroller":9,"./directives/tree":10,"./directives/treeNode":11,"./services/GetData":12}],2:[function(require,module,exports){
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
},{}],3:[function(require,module,exports){
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

},{}],4:[function(require,module,exports){
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
},{}],5:[function(require,module,exports){
module.exports = function($scope, $routeParams, GetData){
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
      console.log(dataReturned);
      $scope.json = dataReturned;
  });

  if($scope.json == null){
    $scope.json = {
      Status: true,
      Records: Array()
    }
  }
  //console.log($scope.json);

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

},{}],6:[function(require,module,exports){
module.exports = function($scope, $routeParams, GetData){
  $scope.projectID = $routeParams.id;
  console.log($scope.data);
}
},{}],7:[function(require,module,exports){
module.exports = function($scope, $parse, GetData){
  function Item(name, type, items){
    this.name = name;
    this.type = type;
    this.items = items || [];
  }
  $scope.data = {
    status: true,
    data: [
      new Item('Project', 'folder-open', [
        new Item('topLevel', 'folder-open', [
          new Item('file_1', 'file'),
          new Item('file_2', 'file'),
          new Item('2ndLevel', 'folder-open', [
            new Item('file_3', 'file'),
            new Item('file_4', 'file'),
          ]),
        ]),
        new Item('file_6', 'file'),
        new Item('file_7', 'file'),
      ])
      ]
  };
  
  console.log($scope.data);
}
},{}],8:[function(require,module,exports){
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
},{}],9:[function(require,module,exports){
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
},{}],10:[function(require,module,exports){
module.exports = function(){
  return {
    template: '<ul><tree-node ng-repeat="item in items"></tree-node></ul>',
    restrict: 'E',
    replace: true,
    scope: {
      items: '=items'
    }
  };
}
},{}],11:[function(require,module,exports){
module.exports = function($compile){
  return {
    restrict: 'E',
    template: '<li><span><i class="glyphicon glyphicon-{{item.type}} text-primary"></i>{{item.name}}</span></li>',
    replace: true,
    link: function(scope, elem, attrs){
      if(scope.item.items.length > 0){
        var children = $compile('<tree items="item.items"></tree>')(scope);
        elem.append(children);
      }
    }
  };
}
},{}],12:[function(require,module,exports){
module.exports= function($http){
  
  function fetch(req, act, data){
    var json = null;

    var requestObj = {
      Request: req,
      Action: act,
      Data: data
    };

    console.log("Here",requestObj);

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