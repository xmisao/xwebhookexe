# xwebhookexe

Tiny HTTP server execute command when received webhook.
This web server is implemented by golang `http` module.

~~~~
xwebhookexe -b 203.0.113.0 -p 8888 -u /webhook -e "cd /path/to/repos && git pull origin master"
~~~~

In this example, xwebhookexe boot up at address 203.0.133.0 and await connection on port 8888.
If it will be accessed to `/webhook`, then execute `cd /path/to/repos && gitpull origin master`.

Usage is here. This is displayed by `xwebhookexe --help`.

~~~~
Usage of ./xwebhookexe:
  -b string
        Bind address.
  -e string
        Execute command by $SHELL. (default "true")
  -p int
        Listen port. (default 8080)
  -u string
        Handle URL. (default "/")
~~~~
