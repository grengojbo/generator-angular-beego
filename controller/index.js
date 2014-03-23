'use strict';
var path = require('path');
var util = require('util');
// var ScriptBase = require('../script-base.js');
// var angularUtils = require('../util.js');
var _ = require('lodash'),
    _s = require('underscore.string'),
    pluralize = require('pluralize'),
    asciify = require('asciify');

var ScriptBase = require('../script-base.js');
var angularUtils = require('../util.js');

module.exports = Generator;

function Generator() {
  ScriptBase.apply(this, arguments);
}

util.inherits(Generator, ScriptBase);

Generator.prototype.createControllerFiles = function createControllerFiles() {
  var controllerDir = 'controllers/';
  this.modelName = this.name;
  this.autorName = 'Oleg Dolya';
  this.sname = this._.slugify(this.name);
  this.template('controllers/_example.go', controllerDir + this.sname + '.go');
  // this.appTemplate('service/factory', 'scripts/services/' + this.name);
  // this.testTemplate('spec/service', 'services/' + this.name);
  // this.addScriptToIndex('services/' + this.name);
};