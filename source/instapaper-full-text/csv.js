const csvToObj = require("csv-to-js-parser").csvToObj;
var assert = require("assert");
var expect = require("chai").expect;

const DataInterface = {
  URL: { type: "string", group: 1 },
  Title: { type: "string" },
  Selection: { type: "string" },
  Folder: { type: "string" },
  Timestamp: { type: "number" },
};

const parseCsv = (data) => csvToObj(data, ",", DataInterface);

module.exports = parseCsv;
