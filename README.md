# The Angular-Beego generator

A [Yeoman](http://yeoman.io) generator for [AngularJS](http://angularjs.org) and [Beego](https://beego.me).

Beego is an open source framework to build and develop your applications in the Go way

## Installation

Install [Git](http://git-scm.com), Mercurial, Bzr, Subversion, [node.js](http://nodejs.org), and [Go](http://golang.org/).

Install Yeoman:

    npm install -g grunt-cli bower yo 

Install the Angular-Go-Martini generator:

    npm install -g generator-beego
    npm install -g ./generator-beego

## Creating a Beego service

In a new directory, generate the service:

    mkdir <myapp> && cd <myapp>
    yo beego
    git init
    git add -A
    git commit -a -m "first commit"
    git remote add origin git@github.com:<LOGIN>/<myapp>.git
    git push -u origin master

Get the dependencies:

    go get -u github.com/gpmgo/gopm
    gopm get -g

Run the service:

  Terminal 1
    grunt server

  Terminal 2
    bee run

Your service will run at [http://localhost:8080](http://localhost:8080).

The Grunt server supports hot reloading of client-side HTML/CSS/Javascript file changes.

## For developer

go get -v github.com/astaxie/beego
go get -v github.com/beego/bee
go get -v github.com/nsf/gocode
go get -v code.google.com/p/go.tools/godoc
go get -v code.google.com/p/go.tools/cmd/godoc
go get -v code.google.com/p/go.tools/cmd/goimports
go get -v code.google.com/p/go.tools/cmd/gotype
go get -v github.com/bytbox/golint


### Разработка с помощью Sublime Text 3 + GoSublime + goimports
Now opening the settings file will succeed. The only thing left is to set the "fmt_cmd" option to ["goimports"], and you're done!

