---
Title: Set Up Docker
Query: true
---

First you will need to have created a <https://hub.docker.com> account.

Then install the `docker` application.

```sh
sudo apt install docker.io
```

That should grab all the dependencies.

Confirm that you have `docker` installed.

```sh
docker -v
```

```{.out}
Docker version 19.03.6, build 369ce74a3c
```


Test the command by logging in and pulling down and running a common image.

```sh
docker login
docker run -ti arch
```
