// @ts-ignore
const tessel = require("tessel");

const http = require("http");
const https = require("https");
const axios = require("axios").default;

// Tesla Powerwall IP
const pwURL = "https://192.168.91.1";

// HTTP client
const client = axios.create({
  baseURL: pwURL,
  httpsAgent: new https.Agent({
    rejectUnauthorized: false,
  }),
});

// Request handler
const reqHandler = async function (req, res) {
  try {
    var url;
    switch (req.url) {
      case "/aggregates":
        url = "/api/meters/aggregates";
        break;
      case "/status":
        url = "/api/system_status/soe";
        break;
      default:
        res.writeHead(404);
        res.end();
    }

    const { data } = await client.get(url);
    res.writeHead(200, { "Content-Type": "application/json" });
    res.end(JSON.stringify(data, null, 2));
  } catch (error) {
    // TODO: improve error handling
    res.writeHead(500);
    res.end(String(error));
  }
};

const server = http.createServer(reqHandler);
// Server listening on port 80
server.listen(80);
