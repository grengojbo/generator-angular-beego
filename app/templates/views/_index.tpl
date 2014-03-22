<!doctype html>
<html lang="en" ng-app="<%= baseName %>">
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title><%= _.capitalize(baseName) %></title>
  <link href="static/lib/bootstrap/dist/css/bootstrap.css" rel="stylesheet">
  <link href="static/lib/jquery-ui/themes/smoothness/jquery-ui.css" rel="stylesheet"/>
  <link href="static/css/app.css" rel="stylesheet">
</head>
<body>
{{<template "header.tpl" .>}}
  <div class="container" ng-view></div>
{{<template "footer.tpl" .>}}
  <script src="static/lib/jquery/dist/jquery.js"></script>
  <script src="static/lib/jquery-ui/ui/jquery-ui.js"></script>
  <script src="static/lib/lodash/dist/lodash.js"></script>
  <script src="static/lib/angular/angular.js"></script>
  <script src="static/lib/angular-resource/angular-resource.js"></script>
  <script src="static/lib/angular-route/angular-route.js"></script>
  <script src="static/lib/angular-bootstrap/ui-bootstrap-tpls.js"></script>
  <script src="static/lib/angular-ui-date/src/date.js"></script>

  <script src="static/js/app.old.js"></script>
  <script src="static/js/home/home-controller.js"></script>
  <% _.each(entities, function (entity) { %>
  <script src="static/js/<%= entity.name %>/<%= entity.name %>-controller.js"></script>
  <script src="static/js/<%= entity.name %>/<%= entity.name %>-router.js"></script>
  <script src="static/js/<%= entity.name %>/<%= entity.name %>-service.js"></script>
  <% }); %>
</body>
</html>
