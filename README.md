# [![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-logo-small.png)](http://www.babygopher.org) http-flooder

You probably don't want to use this, but it starts X amount of requests to the server, doing Y at a time. It's sort of like `ab`, but much more limited and error prone.

# Usage

After a git clone;

```
$ go build
$ ./http-flooder 50 100000 http://yoursite.tld
```

This starts `100000` requests to `http://yoursite.tld`, doing `50` requests at a time.

# Examples

Here's what it looks like for [ma.ttias.be](https://ma.ttias.be) (please don't run this against my server.).

![HTTP flood](https://github.com/mattiasgeniar/http-flooder/raw/master/assets/http-flood.gif)

# Disclosure & liability

This was an experiment. Only launch this against websites or applications you have permissions to. This will stresstest any HTTP(s) server and will likely overwhelm it.

Do not abuse.
