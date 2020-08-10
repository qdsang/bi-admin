
var express = require('express');
var path = require('path')
var proxyMiddleware = require('http-proxy-middleware')

var app = express();

// app.use(proxyMiddleware('/api', {
// 	target: 'http://iot.api.niu.local/iot-dashboard',
//     changeOrigin: true, ws:true
// }))

app.use('/',express.static(path.join(__dirname, '../docs')))

app.get('/api', function (req, res) {
  var j = {msg:'Hello From Server'}
  res.end(JSON.stringify(j));
})

const port = 18082;
const host = '127.0.0.1';
const url = 'http://127.0.0.1:' + port + '/#/screenshot/?code=22&id=2c63c05c-2e60-47fe-8f2b-5162ebf876e3';

app.disable('x-powered-by');

var server = app.listen(port, host, function () {
    var host = server.address().address
    var port = server.address().port
    console.log("listening at http://%s:%s", host, port)
})

// console.log(server.close())

const puppeteer = require('puppeteer');
(async () => {
  const browser = await puppeteer.launch();
  const page = await browser.newPage();
  page.setDefaultTimeout(3 * 1000);
  try {
    await page.goto(url, {waitUntil: 'networkidle2'});
  } catch (ex) { }
  await page.screenshot({path: 'example.png'});
  console.log('ok');
  await browser.close();
})();

