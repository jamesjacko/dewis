module.exports= function($http){
  
  function fetch(req, act, data){
    var json = null;

    var requestObj = {
      Request: req,
      Action: act,
      Data: data
    };

    console.log(requestObj);
    $http({
      method: 'post',
      url: '/api',
      data: requestObj,
      headers: {'Content-Type': 'application/json'}
    }).success(function(data){
      json = data;
    }).error(function(data, status, headers, config){
      json = {
        status: false
      }
    });
    return json;
  }

  return{
    fetch: fetch
  }
}