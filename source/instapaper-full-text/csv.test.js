const fs = require("fs");
const parseCsv = require("./csv");

var expect = require("chai").expect;
const data = fs.readFileSync("./tmp/instapaper-export.csv").toString();

describe("Read csv", function () {
  it("should return non-empty object array", function () {
    let obj = parseCsv(data);
    expect(obj).to.not.be.empty;
    expect(obj[0]).to.have.property("Title").that.is.not.empty;
  });
});
