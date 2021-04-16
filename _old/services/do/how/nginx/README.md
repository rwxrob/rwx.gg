---
Title: Setup Nginx Web Server
---

Here's how to setup an Nginx web server on the Digital Ocean service.

Login as a user with root sudo privileges.

## Install the `nginx` Server Software

```sh
sudo apt install nginx
```

```{.out}
rwxrob@skilstak:~$ sudo apt install nginx
[sudo] password for rwxrob:
Reading package lists... Done
Building dependency tree
Reading state information... Done
The following additional packages will be installed:
  fontconfig-config fonts-dejavu-core libfontconfig1 libgd3 libjbig0 libjpeg-turbo8 libjpeg8 libnginx-mod-http-image-filter
  libnginx-mod-http-xslt-filter libnginx-mod-mail libnginx-mod-stream libtiff5 libwebp6 libxpm4 nginx-common nginx-core
Suggested packages:
  libgd-tools fcgiwrap nginx-doc ssl-cert
The following NEW packages will be installed:
  fontconfig-config fonts-dejavu-core libfontconfig1 libgd3 libjbig0 libjpeg-turbo8 libjpeg8 libnginx-mod-http-image-filter
  libnginx-mod-http-xslt-filter libnginx-mod-mail libnginx-mod-stream libtiff5 libwebp6 libxpm4 nginx nginx-common nginx-core
0 upgraded, 17 newly installed, 0 to remove and 0 not upgraded.
Need to get 2431 kB of archives.
After this operation, 7891 kB of additional disk space will be used.
Do you want to continue? [Y/n] y
```

## Confirm Server Installed and Running

```sh
curl localhost
```

```{.out}
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```

## Create Initial Content

You probably want to change the default Nginx content to point to
something that does not require root access. The safest way to do this
is to create a Git repo and clone it in read-only mode on your new
server. Then you make changes to your site from anywhere (including
GitHub or GitLab using their editors) and update your server by logging
in and doing a `git pull` to pull down your updates.

Let's go ahead and create a temporary directory that will eventually
contain our Git content. We will assume that the name of the repo will
be `www` and that the `html` directory inside of it will contain the
top-level content of our web site.

```sh
mkdir -p www/html
```

Now we need to create a filler HTML page there.

```sh
cd www/html
touch index.html
vi index.html
```

Here's something simple.

```html
<!doctype html>
<html>
<head>
  <title>Coming Soon</title>
  <style>
    body {
      width: 35em;
      margin: 0 auto;
      font-family: Tahoma, Verdana, Arial, sans-serif;
    }
  </style>
</head>
<body>
  <h1>Coming Soon</h1>
  <p>Here is the future home of blah-blah-blah.</p>
</body>
</html>
```

## Change Content Directory in Nginx Configuration

Now we need to update our Nginx configuration to point to our new
temporary web page. 

You'll need `sudo` access for this. Don't proceed
until you have it. 

Also don't proceed until you have setup an additional
user. In other words, don't do this as root.


