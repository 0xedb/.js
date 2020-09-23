const http = require("http");

const server = http.createServer((req, res) => {
  res.end("hello from server");
  res.writeHead(200, )
});


server.listen(9090) 