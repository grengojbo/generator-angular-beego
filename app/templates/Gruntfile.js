'use strict';

var lrSnippet = require('grunt-contrib-livereload/lib/utils').livereloadSnippet;
var proxySnippet = require('grunt-connect-proxy/lib/utils').proxyRequest;
var mountFolder = function(connect, dir) {
    return connect.static(require('path').resolve(dir));
};

module.exports = function (grunt) {
  // require('load-grunt-tasks')(grunt);
  require('matchdep').filterDev('grunt-*').forEach(grunt.loadNpmTasks);
  require('time-grunt')(grunt);

  grunt.initConfig({
    yeoman: {
      // configurable paths
      // app: require('./bower.json').appPath || 'public',
      app: 'app',
      gen: grunt.file.readJSON('generator.json'),
      // lib: grunt.file.readJSON('conf/lib.json'),
      dist: 'public'
    },
    livereload: {
      port: 35728
    },
    sync: {
      dist: {
        files: [{
          cwd: '<%%= yeoman.app %>',
          dest: '<%%= yeoman.dist %>',
          src: '**'
        }]
      }
    },
    shell: {
      get: {
        command: 'go get',
        options: {
          stdout: true,
          stderr: true
          // execOptions: {
          //   encoding: 'windows-1251',
          //   env: {
          //     Gopath: 'F:/home/jbo/src/repo-git/sssua-app'
          //   }
          // }
        }
      },
      wget: {
        command: [
          'go get github.com/astaxie/beego/validation',
          'go get github.com/beego/i18n',
          'go get github.com/beego/wetalk/modules/utils',
          'go get github.com/astaxie/beego/context',
          'go get -u github.com/beego/social-auth',
          'go get github.com/astaxie/beego/cache',
          'go get github.com/grengojbo/beego/modules/utils'
          ].join(';'),
        options: {
          stdout: true,
          stderr: true
          // execOptions: {
          //   encoding: 'windows-1251',
          //   env: {
          //     Gopath: 'F:/home/jbo/src/repo-git/sssua-app'
          //   }
          // }
        }
      }
    },
    watch: {
      // options: {
      //  livereload: 35729
      // },
      // src: {
      //  files: [
      //    '<%%= yeoman.app %>/*.html',
      //    '<%%= yeoman.app %>/css/**/*',
      //    '<%%= yeoman.app %>/js/**/*',
      //    '<%%= yeoman.app %>/views/**/*',
   //        '<%%= yeoman.app %>/views/*.tpl'
      //  ]
      //  //tasks: ['sync:dist']
      // },
      coffee: {
        tasks: ['coffee:dist'],
        files: ['<%%= yeoman.app %>/static/js/{,*/}*.coffee']
      },
      less: {
        // tasks: ['less:dev', 'autoprefixer'],
        tasks: ['less:dev'],
        files: ['<%%= yeoman.app %>/static/less/{,*/}*.less']
      },
      livereload: {
        tasks: ['livereload'],
        files: [
          '<%%= yeoman.app %>/views/*.{html,htm,tpl}',
          'views/*.{html,htm,tpl}',
          'static/css/{,*/}*.css',
          'static/js/{,*/}*.js',
          'static/img/{,*/}*.{png,jpg,jpeg,gif,webp}'
        ]
      },
    },
    connect: {
      proxies: [
        {
          context: '/',
          port: 8080,
          https: false,
          // changeOrigin: false,
          host: 'localhost'
        }
      ],
      options: {
        port: 3000,
        debug: true,
        livereload: 35729,
        // Change this to '0.0.0.0' to access the server from outside.
        hostname: 'localhost'
      },
      livereload: {
        options: {
          // open: true,
          // base: [
            // '.tmp'
            // ''
          // ],
          middleware: function (connect) {
            // return [lrSnippet, mountFolder(connect, '.tmp'), mountFolder(connect, '.')];
            // return [lrSnippet, proxySnippet, mountFolder(connect, '.tmp'), mountFolder(connect, '.')];
            // return [proxySnippet, connect.static(require('path').resolve('.tmp'))];
            return [lrSnippet, proxySnippet, connect.static(require('path').resolve('.'))];
          }
        }
      },
      /*
      dist: {
        options: {
          base: '<%%= yeoman.dist %>'
        }
      }
      */
    },
    less: {
      dev: {
        files: [
          // {'.tmp/static/css/bootstrap.css': '/static/less/bootstrap.less'},
          {'static/css/aplication.css': '<%%= yeoman.app %>/static/less/aplication.less'}
        ],
        options: {
          paths: ['/static/less']
        }
      },
      dist: {
        options: {
          paths: ['/static/less'],
          yuicompress: true
        },
        files: [
          {'.tmp/static/css/aplication.css': '/static/less/aplication.less'},
          {'.tmp/static/css/bootstrap.css': '/static/less/bootstrap.less'}
        ]
      }
    },
    coffee: {
      options: {
        sourceMap: true,
        sourceRoot: ''
      },
      dist: {
        files: [{
          // rather than compiling multiple files here you should
          // require them into your main .coffee file
          expand: true,
          cwd: '<%%= yeoman.app %>/static/js',
          src: '{,*/}*.coffee',
          dest: 'static/js',
          ext: '.js'
        }]
      },
      test: {
        files: [{
          expand: true,
          cwd: '.tmp/spec',
          src: '{,*/}*.coffee',
          dest: 'test/spec'
        }]
      }
    },
    // Put files not handled in other tasks here
    copy: {
      dist: {
        files: [{
          expand: true,
          dot: true,
          cwd: '<%%= yeoman.app %>',
          dest: '<%%= yeoman.dist %>',
          src: '**'
        }]
      },
    },
    // Test settings
    karma: {
      unit: {
        configFile: 'test/config/karma.conf.js',
        singleRun: true
      }
    },
    bowerInstall: {
      target: {
        src: ['views/index.tpl'],
        ignorePath: '../'
      }
    },
    bowercopy: {
      options: {
        // Bower components folder will be removed afterwards
        // clean: true,
        destPrefix: '<%%= yeoman.app %>'
      },
      test: {
        files: {
          'test/lib/angular-mocks': 'angular-mocks',
          'test/lib/angular-scenario': 'angular-scenario'
        }
      }
    },
    open: {
      server: {
        path: 'http://localhost:<%%= connect.options.port %>'
      }
    },
    // jshint: {
    //   options: {
    //             jshintrc: '.jshintrc'
    //         },
    //         dev: [
    //             'Gruntfile.js',
    //             './tmp/static/js/{,*/}*.js'
    //         ],
    //         all: [
    //             'Gruntfile.js',
    //             './tmp/static/js/{,*/}*.js',
    //             '<%%= yeoman.app %>/static/js/{,*/}*.js',
    //             '!<%%= yeoman.app %>/static/js/vendor/*',
    //             'test/spec/{,*/}*.js'
    //         ]
    // },
    // useminPrepare: {
    //         html: 'dist/base.html',
    //         options: {
    //             dest: 'dist'
    //         }
    //     },
    //     usemin: {
    //         // html: ['dist/{,*/}*.html'],  // INFO minimize all files
    //         html: ['dist/tpl-home.html', 'dist/base.html'],  // INFO minimize only base.html
    //         // html: ['dist/base.html'],  // INFO minimize only base.html
    //         css: ['dist/static/css/{,*/}*.css'],
    //         options: {
    //             // dirs: ['dist'] // NO REV
    //             assetsDirs: ['dist'] // REVISON big time
    //         }
    //     },
    //     imagemin: {
    //         dist: {
    //             files: [{
    //                 expand: true,
    //                 cwd: '<%%= yeoman.app %>/static/img',
    //                 src: '{,*/}*.{png,jpg,jpeg}',
    //                 dest: '<%%= yeoman.dist %>/img'
    //             }]
    //         }
    //     },
    clean: {
      tmp: ['.tmp/*.html', '.tmp/tpl', '.tmp/<%%= yeoman.mobile %>'],
      dist: ['.tmp', 'dist', '*.dmp'],
      minifier: [
        '<%%= yeoman.app %>/css/*.{style,bootstrap,font-awesome}.css',
        '<%%= yeoman.app %>/js/*.{jquery.min,vendor,main,angular-lib}.js'
      ],
      server: '.tmp'
    }
  });

  // grunt.renameTask('regarde', 'watch');

  grunt.registerTask('server', function (target) {
    grunt.task.run([
      'clean:server',
      // 'clean:minifier',
      'coffee:dist',
      'less:dev',
      'livereload-start',
      'configureProxies',
      'connect:livereload',
      'open',
      'watch'
    ]);
  });

};
