# <%= _.camelize(baseName) %>

## Installation <%= _.camelize(baseName) %>
 Download mercurial http://mercurial.selenic.com/downloads
 Install mercurial 

```shell
go get github.com/mattn/gom
npm install && bower install && gom install
grunt shell:get
yo beego:api Post

```

### Inject your Bower dependencies right into your HTML from Grunt.
Call the Grunt task:
```
grunt bowerInstall
```

You're in business!
```html
<!-- bower:css -->
<script src="bower_components/jquery/jquery.js"></script>
<!-- endbower -->
...
<!-- bower:js -->
<script src="bower_components/jquery/jquery.js"></script>
<!-- endbower -->
```
