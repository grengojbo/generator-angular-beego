// Declare app level module which depends on filters, and services
angular.module('<%= baseName %>', ['ngResource', 'ngRoute', 'ui.bootstrap', 'ui.date'])
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: '/static/views/home/home.html',
        controller: 'HomeController'});
  }]);
