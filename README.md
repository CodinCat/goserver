# goserver

A minimum web server that serves your static files.

```
go get github.com/CodinCat/goserver
```

If you don't have [Go](https://golang.org/) installed and using Windows, here is the pre-built binary (only windows 64bit for now): 
https://github.com/CodinCat/goserver/releases/latest

Download it and put it in your `PATH`, rename it to anything you like.

Restart your terminal, then you can start to use the web server:

```
cd my/web/project
goserver
```
It will listen on 8080 port by default. Or you can specify a port:
```
goserver --port :7777
```

(don't forget the colon `:`)

The binary is portable, so you can even put it in your USB drive and use it anywhere.

And I recommend [LivePage](https://chrome.google.com/webstore/detail/livepage/pilnojpmdoofaelbinaeodfpjheijkbh) extension for live reloading, it's super lightweight and easy to use.

That's all!
