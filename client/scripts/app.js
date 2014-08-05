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