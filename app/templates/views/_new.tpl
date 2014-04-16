{{<template "base.tpl" .>}}
{{<define "meta">}}
  <title>{{<.Title>}} : {{<.Website>}}</title>
{{<end>}}
{{<template "header.tpl" .>}}
{{<define "body">}}
  <section class="container" ui-view >
  Loading...
  </section>
{{<end>}}
{{<template "footer.tpl" .>}}