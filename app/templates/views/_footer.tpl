{{<define "footer">}}
<footer id="footer">
  <div class="container navbar-fixed-bottom jbo-line">
    <p>{{<i18n .Lang "copyright">}} © 2014</p>
    <p class="desc"><!-- TODO loadtimes ??? => PageStartTime|loadtimes -->
      {{<i18n .Lang "time_load">}} {{<.PageStartTime|loadtimes>}}ms. {{<i18n .Lang "powered_by">}} <a target="_blank" href="http://beego.me">Beego</a>. {{<i18n .Lang "based_on">}} <a target="_blank" href="http://getbootstrap.com/">Bootstrap</a>. {{<i18n .Lang "icons_from">}} <a target="_blank" href="http://fortawesome.github.io/Font-Awesome/">Font Awesome</a>.
    </p>
  </div>
</footer>
{{<end>}}