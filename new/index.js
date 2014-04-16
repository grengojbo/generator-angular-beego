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
  // this.autorName = 'Oleg Dolya';
  this.sname = this._.slugify(this.name);
  
  var prompts = [
    {
      type: 'input',
      name: 'pRouter',
      message: 'Routing url',
      default: '/' + this.sname + '/'
    },
    {
      type: 'input',
      name: 'pTitle',
      message: 'Page Title',
      default: 'page title'
    },
    {
      type: 'input',
      name: 'pMenu',
      message: 'Menu Name',
      default: 'Home'
    }
  ];

  this.prompt(prompts, function (props) {
    this.routerName = props.pRouter;
    // this.autorName = 'Oleg Dolya';
    this.tTitle = props.pTitle;
    this.tMenu = props.pMenu;

    this.dequote = function (str) {
      return str.replace(/\"/gm, '\\"');
    };

    cb();
  }.bind(this));
};

Generator.prototype.createControllerFiles = function createControllerFiles() {
  // var controllerDir = 'controllers/';
  // var modelsDir = 'models/';

  // this.mkdir(viewsDir);
  // this.mkdir(controllers);


  this.beegoTemplate('views/_new.tpl', this.sname);
  // this.beegoController('views/_new.tpl', this.sname);
  this.template('controllers/_new.go', this.controllerDir + this.sname + '.go');
  // this.template('models/_example.go', modelsDir + this.sname + '.go');
  // this.appTemplate('service/factory', 'scripts/services/' + this.name);
  // this.testTemplate('spec/service', 'services/' + this.name);
  this.addMenuItem(this.tMenu, this.routerName, 'nav.tpl')
  this.addToRoute(this.routerName, this.classedName);
  // console.log('\nUnable to find '.yellow + fullPath + '. Reference to '.yellow + script + '.js ' + 'not added.\n'.yellow);
};