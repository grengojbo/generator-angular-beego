'use strict';
var util = require('util'),
    path = require('path'),
    yeoman = require('yeoman-generator'),
    _ = require('lodash'),
    _s = require('underscore.string'),
    pluralize = require('pluralize'),
    asciify = require('asciify');

var AngularGoMartiniGenerator = module.exports = function AngularGoMartiniGenerator(args, options, config) {
  yeoman.generators.Base.apply(this, arguments);

  this.on('end', function () {
    this.installDependencies({ skipInstall: options['skip-install'] });
  });

  this.pkg = JSON.parse(this.readFileAsString(path.join(__dirname, '../package.json')));
};

util.inherits(AngularGoMartiniGenerator, yeoman.generators.Base);

AngularGoMartiniGenerator.prototype.askFor = function askFor() {

  var cb = this.async();

  console.log('\n'+
'                                                                                \n' +
'                  ,===============           :==============~                   \n' +
'                 :=============~               :==============                  \n' +
'                 =============,                  =============,                 \n' +
'                 ============,                    ============,                 \n' +
'                 ============        =====,       :===========,                 \n' +
'                 ===========,      ,=======~       ===========,                 \n' +
'                 ===========       =========       ===========,                 \n' +
'                 ======~,          =========           :======,                 \n' +
'                 ====   ~======~===========:     ~=====,  ~===,                 \n' +
'                 ==, :===================~       :========  ==,                 \n' +
'                 =  =================,           ==========  =,                 \n' +
'                 , =======~,,~========           ,,,:=======  ,                 \n' +
'                  ~=====,      :=====,                 ======                   \n' +
'                  =====:        :=====   :=====         =====                   \n' +
'                  =====,        ,=====   :=====         =====,                  \n' +
'                  =====~        ,=~:=~   ,=====         =====                   \n' +
'                  :=====~                 ======,     :======                   \n' +
'                 : ~=======~~=           :=========~========  ,                 \n' +
'                 =, ~=========            =================  =,                 \n' +
'                 ==~  ========       ====================  ,==,                 \n' +
'                 ====:  ,~====     ,========~,~:,~===~,  ,====,                 \n' +
'                 =======~,         =========,        ,:=======,                 \n' +
'                 ===========       =========       ===========,                 \n' +
'                 ===========:       =======:       ===========,                 \n' +
'                 ============        ,===~        :===========,                 \n' +
'                 ============~                   ,============,                 \n' +
'                 =============~                 ,=============,                 \n' +
'                 ,==============,              ==============~                  \n' +
'                   :==============~         :==============~                    \n' +
'                                                                                \n' +
'                      :~~~,                                                     \n' +
'                   :=========                                                   \n' +
'                  =====~:~====~===== =====:=======: ==========,                 \n' +
'                 ====      ,======== =====:=========~=========,                 \n' +
'                 ===:           ===   ===, ===   ====== ==~,==,                 \n' +
'                 ===,           ===   ===, ===   ===,  ,===                     \n' +
'                 ====       :::,===   ===, ========~   ,===                     \n' +
'                 ,====:   :==== ===   ===, =======     ,===                     \n' +
'                   ===========  ~======== ==== :====~ ~=====                    \n' +
'                     ======~     :======  ====  ,===~ ~=====                    \n' +
'                                                                                \n' +
  	'\n');
  console.log('\n' +
    '+-+-+-+-+-+-+-+ +-+-+ +-+-+-+-+-+-+-+ +-+-+-+-+-+-+-+-+-+\n' +
    '|   angular    | go  |     beego     |    generator     |\n' +
    '+-+-+-+-+-+-+-+ +-+-+ +-+-+-+-+-+-+-+ +-+-+-+-+-+-+-+-+-+\n' +
    '\n');

  var prompts = [{
    type: 'input',
    name: 'baseName',
    message: 'What is the name of your application?',
    default: 'myapp'
  }];

  this.prompt(prompts, function (props) {
    this.baseName = props.baseName;
    this.goBin = process.env.GOROOT+'/bin';
    this.baseDir = './';
    this.appPort = 8080;
    this.modelName = 'Example';
    this.autorName = 'Oleg Dolya';
    cb();
  }.bind(this));
};

AngularGoMartiniGenerator.prototype.app = function app() {

  this.entities = [];
  this.resources = [];
  this.generatorConfig = {
    "baseName": this.baseName,
    "goBin": this.goBin,
    "baseDir": this.baseDir,
    "entities": this.entities,
    "resources": this.resources,
    "appPort": this.appPort
  };
  this.generatorConfigStr = JSON.stringify(this.generatorConfig, null, '\t');

  this.template('_README.md', 'README.md');
  this.template('_LICENSE', 'LICENSE');
  this.template('_generator.json', 'generator.json');
  this.template('_package.json', 'package.json');
  this.template('_bower.json', 'bower.json');
  this.template('bowerrc', '.bowerrc');
  this.template('Gruntfile.js', 'Gruntfile.js');
  this.template('_gitignore', '.gitignore');

  var confDir = 'conf/'
  var controllerDir = 'controllers/'
  var modelsDir = 'models/'
  var publicDir = 'static/'
  var routesDir = 'routers/'
  var testsDir = 'tests/'
  var viewsDir = 'views/'
  var appDir = 'app/'
  var appStatic = 'app/static/'
  var appLess = 'app/static/less/'
  var appJs = 'app/static/js'
  this.mkdir(confDir);
  this.mkdir(controllerDir);
  this.mkdir(modelsDir);
  this.mkdir(publicDir);
  this.mkdir(routesDir);
  this.mkdir(testsDir);
  this.mkdir(viewsDir);
  this.mkdir(appDir);
  this.mkdir(appStatic);
  this.mkdir(appLess);
  this.mkdir(appJs);

  this.template('_main.go', 'main.go');
  this.template('_bee.json', 'bee.json');
  this.template('conf/_app.conf', confDir + 'app.conf');
  this.template('conf/_locale_en-US.ini', confDir + 'locale_en-US.ini');
  this.template('conf/_locale_ua-UA.ini', confDir + 'locale_ua-UA.ini');
  this.template('controllers/_default.go', controllerDir + 'default.go');
  this.template('controllers/_example.go', controllerDir + 'example.go');
  this.template('models/_user.go', modelsDir + 'user.go');
  this.template('models/_example.go', modelsDir + 'example.go');
  this.template('routers/_router.go', routesDir + 'router.go');
  this.template('views/_index.tpl', viewsDir + 'index.tpl');

  var publicCssDir = publicDir + 'css/';
  var publicLessDir = publicDir + 'less/';
  var publicJsDir = publicDir + 'js/';
  var publicViewDir = publicDir + 'views/';
  this.mkdir(publicCssDir);
  this.mkdir(publicJsDir);
  this.mkdir(publicViewDir);
  this.template('views/_index.tpl', publicViewDir + 'index.tpl');
  this.copy('static/css/app.css', publicCssDir + 'app.css');
  this.copy('static/less/variables.less', publicLessDir + 'variables.less');
  this.copy('static/less/mixins.less', publicLessDir + 'mixins.less');
  this.copy('static/less/aplication.less', appLess + 'aplication.less');
  this.template('static/js/_app.js', publicJsDir + 'app.old.js');
  this.template('static/js/_app.coffee', appJs + 'app.coffee');
  this.template('static/js/home/_home-controller.js', publicJsDir + 'home/home-controller.js');
  this.template('static/views/home/_home.html', publicViewDir + 'home/home.html');
};

AngularGoMartiniGenerator.prototype.projectfiles = function projectfiles() {
  this.copy('editorconfig', '.editorconfig');
  this.copy('jshintrc', '.jshintrc');
};
