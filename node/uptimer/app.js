const http = require("http");

const server = http.createServer((req, res) => {
  res.end("hello from server");
  console.log(req.method, req.headers)
  res.writeHead(200, )
});


server.listen(9090) 