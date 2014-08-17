angular.module('dewis', ['ngRoute'])


  .config(['$routeProvider', function($routeProvider){
    $routeProvider
      .when("/login", {templateUrl: "partials/login.html", controller: "LoginController"})
      .when("/new-user", {templateUrl: "partials/new-user.html", controller: "NewUserController"})
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
    NewUserController: require('./controllers/NewUserController'),
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