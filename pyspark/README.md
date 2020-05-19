# ☸️ Graph Theory Course's PySpark Workspace

## 🐳 Docker installation and usage

Instead of having you go through the hassle of downloading a bunch of dependencies to your system and dealing with all sorts of conflicts, configurations and so on... we will be using Docker! If this is new to you, don't worry - we will guide you through the process 😜

>**_Feeling curious?_** Here's an [awesome list](https://github.com/veggiemonk/awesome-docker#what-is-docker) with a ton of Docker resources, for you to take a look!

### Requirements

Okay, maybe you will have go through _some_ software installation... but I promise that it will only be these two and that after installing `Docker` and `Docker-compose`, you won't ever have worry about project dependencies conflicting with system ones anymore 🤩

To install them, please follow the instructions in the links below:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker-compose](https://docs.docker.com/compose/install/)

> **_note_**: if you're using a Linux system, please take a look at [Docker's post-installation steps for Linux](https://docs.docker.com/engine/install/linux-postinstall/)!

### Building and running

Once you have `Docker` and `Docker-compose`, change your current working directory to `/pyspark` workspace then build and run the container:

```bash
# change current working directory
$ cd <path/to/cs-graph-theory/pyspark>

# start the container in the background of your terminal
$ docker-compose up --detach
```

At this point, [Jupyter Notebook](https://jupyter.org) will be running at: `http://localhost:8888`

## 📦 Installing new packages

There are a few ways you may install packages to the container. It'll depend on your goal and needs.

### Pip

If you need to do update or add packages via `pip`, execute the following command **inside your jupyter notebook**:

```bash
import sys

# install a pip package in the current Jupyter kernel
!{sys.executable} -m pip install <package>
```

> _**note**_: the `!` notation is used to run `pip` directly as a shell command from the notebook. Also, take a look [here](https://jakevdp.github.io/blog/2017/12/05/installing-python-packages-from-jupyter/) to see why you should NOT use `!pip install <package>`.

### Conda

If you need to do update or add packages via `conda`, execute the following command **inside your jupyter notebook**:

```bash
import sys

# install a conda package in the current Jupyter kernel
!conda install --yes --prefix {sys.prefix} <package>
```

> _**note**_: the `!` notation is used to run `conda` directly as a shell command from the notebook. Also, take a look [here](https://jakevdp.github.io/blog/2017/12/05/installing-python-packages-from-jupyter/) to see why you should NOT use `!conda install --yes <package>`.

### System

To add or update system packages, you will need `root` user permissions. To achieve this, use the following command:

```bash
# execute the container's shell
$ docker exec -it --user root tensorflow_notebook /bin/bash

# install a package to the system the container runs on
$ sudo apt install <package>
```

## 👋 Wrapping up

Once you're done, you may remove what was created by `docker-compose up`:

```bash
# change current working directory
$ cd <path/to/cs-graph-theory/pyspark>

# stop containers and removes containers, networks, volumes, and images created by `docker-compose up`
$ docker-compose down
```
