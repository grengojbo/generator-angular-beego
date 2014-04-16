{{<template "base.tpl" .>}}
{{<define "meta">}}
  <title>{{<.Website>}}</title>
{{<end>}}
{{<template "header.tpl" .>}}
{{<define "body">}}
  <section class="hero-unit">
    <div class="container" ng-view></div>
  </section>
{{<end>}}
{{<template "footer.tpl" .>}}