###global app, prettyparams, ###
'use strict'

window.app = angular.module('<%= baseName %>', [
  'ngRoute',
  # 'ngMockE2E',
  'ngCookies',
  'ngResource',
  'ngSanitize',
  # 'ui.bootstrap',
  # 'ui.router',
  'ngRoute',
  'chieffancypants.loadingBar',
  'ngAnimate',
  'ui.date'
  ])


# app.config(['$routeProvider', ($routeProvider) ->
#   $routeProvider
#     .when '/aaa',
#       redirectTo: '/'
#     .otherwise redirectTo: '/'
#   return false
# ])

app.config ['$routeProvider', ($routeProvider) ->
  $routeProvider
    .when '/',
      templateUrl: '/static/views/home/home.html',
      controller: 'HomeCtrl'
  return false
]

#VideoController.$inject = ['$scope', 'Video'];