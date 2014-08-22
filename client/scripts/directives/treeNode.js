module.exports = function($compile, $log){
  return {
    restrict: 'E',
    template: '<li ng-click="hi()">'+
              ' <span>'+
              '   <i class="glyphicon glyphicon-{{item.type}} text-success"></i>'+
              '     {{item.name}}'+
              '   <i class="glyphicon glyphicon-plus text-primary"></i>'+
              ' </span>'+
              '</li>',
    replace: true,
    link: function(scope, elem, attrs){
      if(scope.item.items.length > 0){
        var children = $compile('<tree items="item.items"></tree>')(scope);
        elem.append(children);
      }
      scope.hi - function(){ alert("hi"); }
    }
  };
}