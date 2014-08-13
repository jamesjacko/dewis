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