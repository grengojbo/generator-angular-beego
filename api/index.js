'use strict';
var path = require('path');
var util = require('util');
var _ = require('lodash'),
    _s = require('underscore.string'),
    pluralize = require('pluralize'),
    asciify = require('asciify');
// var gitconfig = require('git-config');

var ScriptBase = require('../script-base.js');
var angularUtils = require('../util.js');

module.exports = Generator;

function Generator() {
  ScriptBase.apply(this, arguments);
}

util.inherits(Generator, ScriptBase);

Generator.prototype.askFor = function askFor() {
  var cb = this.async();  

  this.modelName = this.name;
  this.cameledName = this._.camelize(this.name);  // jbo-app -> jboApp
  this.classedName = this._.classify(this.name); // jbo-app -> JboApp
  this.sname = this._.slugify(this.name);
  
  var prompts = [
    {
      type: 'input',
      name: 'pRouter',
      message: 'Routing url',
      default: '/api/v1/' + this.sname + '/'
    }, {
      type: 'input',
      name: 'pClassName',
      message: 'Prefix Controller Name',
      default: this.classedName
    }, {
      type: 'input',
      name: 'pTableName',
      message: 'SQL Table Name',
      default: this.sname
    }
  ];

  this.prompt(prompts, function (props) {
    this.routerName = props.pRouter;
    this.classedName = props.pClassName;
    this.tableName = props.pTableName;

    this.dequote = function (str) {
      return str.replace(/\"/gm, '\\"');
    };

    cb();
  }.bind(this));
};
Generator.prototype.createControllerFiles = function createControllerFiles() {
  var controllerDir = 'controllers/';
  var modelsDir = 'models/';
  this.modelName = this.name;
  this.cameledName = this._.camelize(this.name);  // jbo-app -> jboApp
  this.classedName = this._.classify(this.name); // jbo-app -> JboApp
  this.autorName = 'Oleg Dolya';
  this.sname = this._.slugify(this.name);
  this.template('controllers/_example.go', controllerDir + this.sname + '.go');
  this.template('models/_example.go', modelsDir + this.sname + '.go');
  // this.appTemplate('service/factory', 'scripts/services/' + this.name);
  // this.testTemplate('spec/service', 'services/' + this.name);
  this.addApiToRoute(this.sname, this.classedName);
};