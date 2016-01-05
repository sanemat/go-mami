'use strict';
var execSync = require('child_process').execSync;
var URI = require('urijs');

var version = execSync('go run cmd/mami/mami.go -v').toString('utf-8').trim();
var homepageUrl = 'https://github.com/sanemat/go-mami';
var url = new URI(homepageUrl);
var host = url.protocol() + '://' + url.authority();
var owner = url.pathname().split('/')[1];
var repository = url.pathname().split('/')[2];

module.exports = {
  version: version,
  host: host,
  owner: owner,
  repository: repository
};
