var ServerRequest = function(method, path, data, callback) {
  var request = new XMLHttpRequest();
  request.open(method, path);
  request.onreadystatechange = function () {
    if (request.readyState === 4) {
      return callback(JSON.parse(request.responseText))
    }
  }

  return request.send(JSON.stringify(data));
}

