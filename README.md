# The Angular-Beego generator

A [Yeoman](http://yeoman.io) generator for [AngularJS](http://angularjs.org) and [Beego](https://beego.me).

Beego is an open source framework to build and develop your applications in the Go way

## Installation

Install [Git](http://git-scm.com), [node.js](http://nodejs.org), and [Go](http://golang.org/).

Install Yeoman:

    npm install -g yo

Install the Angular-Go-Martini generator:

    npm install -g generator-beego

## Creating a Beego service

In a new directory, generate the service:

    yo angular-beego

Get the dependencies:

    go get

Run the service:

  Terminal 1
    grunt server

  Terminal 2
    bee run

Your service will run at [http://localhost:8080](http://localhost:8080).

The Grunt server supports hot reloading of client-side HTML/CSS/Javascript file changes.

