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
