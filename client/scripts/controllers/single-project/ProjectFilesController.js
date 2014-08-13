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