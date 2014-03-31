###global app, ###
'use strict'

app.controller 'HomeCtrl', ['$scope', ($scope) ->
  $scope.content = 'AngularJS controller HomeCtrl!'
  return false
]