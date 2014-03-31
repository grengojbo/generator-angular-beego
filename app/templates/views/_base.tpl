<!doctype html>
<html lang="en" ng-app="<%= baseName %>">
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta name="author" content="Oleg Dolya"/>
  <!-- <link rel="shortcut icon" href="static/ico/favicon.png"> -->
  <!-- build:css({.tmp,app}) static/css/libs.css -->
  <!-- bower:css -->
  <link rel="stylesheet" href="static/lib/bootstrap/dist/css/bootstrap.css" />
  <link rel="stylesheet" href="static/lib/angular-loading-bar/src/loading-bar.css" />
  <link rel="stylesheet" href="static/lib/angular-formstamp/build/formstamp.css" />
  <!-- endbower -->
  <link href="static/lib/jquery-ui/themes/smoothness/jquery-ui.css" rel="stylesheet"/>
  <!-- endbuild -->
  <!-- build:css({.tmp,app}) static/css/aplication.css -->
  <link href="static/css/aplication.css" rel="stylesheet">
  <!-- endbuild -->
  <!-- <link title="RSS Feed" rel="alternate" type="application/atom+xml" href="/feeds/atom/"/> -->
  <!-- <link rel="self" type="application/rss+xml" title="RSS Feed" href="/feeds/rss/"/> -->
  {{<template "meta" .>}}
</head>
<body ng-controller="MainCtrl">
  {{<define "header">}}
    {{<template "base/header.tpl" .>}}
  {{<end>}}
  {{<template "body" .>}}
  {{<define "footer">}}
    {{<template "base/footer.tpl" .>}}
  {{<end>}}
  <!-- build:js({.tmp,app}) static/js/jquery.js -->
  <script src="static/lib/jquery/dist/jquery.js"></script>
  <script src="static/lib/jquery-ui/ui/jquery-ui.js"></script>
  <!-- endbuild -->
  <!-- build:js({.tmp,app}) static/js/angular.js -->
  <script src="static/lib/angular/angular.js"></script>
  <script src="static/lib/angular-resource/angular-resource.js"></script>
  <script src="static/lib/angular-route/angular-route.js"></script>
  <script src="static/lib/angular-ui-date/src/date.js"></script>
  <!-- endbuild -->
  <!-- build:js({.tmp,app}) static/js/lib.js -->
  <!-- bower:js -->
  <script src="static/lib/bootstrap/dist/js/bootstrap.js"></script>
  <script src="static/lib/underscore/underscore.js"></script>
  <script src="static/lib/angular-cookies/angular-cookies.js"></script>
  <script src="static/lib/angular-sanitize/angular-sanitize.js"></script>
  <script src="static/lib/angular-animate/angular-animate.js"></script>
  <script src="static/lib/angular-ui-router/release/angular-ui-router.js"></script>
  <script src="static/lib/angular-loading-bar/src/loading-bar.js"></script>
  <script src="static/lib/json3/lib/json3.min.js"></script>
  <script src="static/lib/lodash/dist/lodash.compat.js"></script>
  <script src="static/lib/angular-formstamp/build/formstamp.js"></script>
  <!-- endbower -->
  <!-- endbuild -->
  <!-- build:js({.tmp,app}) static/js/app.js -->
  <script src="static/js/app.js"></script>
  <script src="static/js/controllers/main.js"></script>
  <script src="static/js/controllers/home.js"></script>
  <!-- endbuildapp -->
  <!-- endbuild -->
</body>
</html>
