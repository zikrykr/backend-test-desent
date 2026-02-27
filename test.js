const http = require('http');

const options = {
  hostname: 'localhost',
  port: 3001,
  path: '/api/echo',
  method: 'POST',
  headers: {
    'Content-Type': 'application/json'
  }
};

const req = http.request(options, (res) => {
  let data = '';
  res.on('data', (d) => {
    data += d;
  });
  res.on('end', () => {
    console.log(data);
  });
});

req.write('{"large": 1234567890123456789, "message": "hello", "nested": {"key": "val"}, "number": 42}');
req.end();
