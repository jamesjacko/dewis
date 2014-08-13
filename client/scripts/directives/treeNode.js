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