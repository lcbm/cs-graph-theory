# ☸️ Graph Theory Course's Go Workspace

## 🐳 Docker installation and usage

Instead of having you go through the hassle of downloading a bunch of dependencies to your system and dealing with all sorts of conflicts, configurations and so on... we will be using Docker! If this is new to you, don't worry - we will guide you through the process 😜

>**_Feeling curious?_** Here's an [awesome list](https://github.com/veggiemonk/awesome-docker#what-is-docker) with a ton of Docker resources, for you to take a look!

### Requirements

Okay, maybe you will have go through _some_ software installation... but I promise that it will only be this one and that after installing `Docker`, you won't ever have worry about project dependencies conflicting with system ones anymore 🤩

To install `Docker`, please follow the instructions in the link below:

- [Docker](https://docs.docker.com/get-docker/)

> **_note_**: if you're using a Linux system, please take a look at [Docker's post-installation steps for Linux](https://docs.docker.com/engine/install/linux-postinstall/)!

### Building and running

Once you have `Docker` and `Docker-compose`, change your current working directory to `/go` workspace and build the `Docker image`:

```bash
# change current working directory
$ cd <path/to/cs-graph-theory/go>

# start the container in the background of your terminal
$ docker build . --tag cesarschool/graph-theory-go
```

Finally, run the container with the following command:

```bash
docker run --rm --volume `pwd`:/go --tty --interactive cesarschool/graph-theory-go /bin/sh
```

>**_note_**: the `--volume` flag binds and mounts `pwd` (your current working directory) to the container's `/go` directory. This means that the changes you make outside the container will be reflected inside (and vice-versa). You may use your IDE to make code modifications, additions, deletions and so on, and these changes will be persisted both in and outside the container.

## 📝 Exercises usage

### Running

Once you are inside the container, you may quickly run an exercise's (e.g. `graph`) code with the the following [`go` command](https://golang.org/pkg/cmd/go/):

```bash
# change current working directory
$ cd graph

# run the named main Go package
$ go run .
```

### Compiling

To compile an exercise's (e.g. `graph`) code, use the following [`go` commands](https://golang.org/pkg/cmd/go/) instead:

```bash
# change current working directory
$ cd graph

# compile package and its dependencies (force rebuild with -a)
$ go build -a -o executable .
```

To run the compiled binary from the previous step, simply execute it as shown below:

```bash
# execute the compiled binary
$ ./executable
```

## 👋 Wrapping up

Once you're done, you may remove what was created by `docker build`:

```bash
# remove the created container
$ docker rmi cesarschool/graph-theory-go --force
```
