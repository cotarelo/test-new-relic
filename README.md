# New Relic Infrastructure - SRE/Build engineer

## Prerequisites
Virtualbox >= 5.1
Vagrant >= 2.0

### Exercise 1
I have chosen to install Ubuntu 16.04 as a happy path. . Go to the Exercise1 directory and run the command below to start the virtual machine

```$ vagrant up```

The following command to provision the virtual machine with New Relic Infraestructure software however it provisions by default on the first run  

```$ vagrant provision```

The software is installed according the instructions in https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure/installation/install-infrastructure-linux

It will use Vagrant shell provisioner to deploy and configure the software. **Note** that I used shell provisioner for simplicity taking into consideration that is only one host, when deploying the software in multiple hosts I would use a puppet or ansible provisioner.

After installing the sofware it can be seen how the data is sent to New Relic under the user Acme_125 https://rpm.newrelic.com/accounts/1785110

### Exercise 2

I have followed the redis tutorial of the custom for Redis and adapted it to calculate the size of a folder. The code and readme is on the Exercise2 folder.

### Exercise 3

I used the fpm tool to generate a debian package. After the integration was done is enough to copy the integration files into a empty directory replicating the structure of the files. This could be done with a shell script

```
#!/bin/sh
mkdir package
cd package
mkdir -p var/db/newrelic-infra/custom-integrations
cp /var/db/newrelic-infra/custom-integrations/test-folderSize-definition.yaml var/db/newrelic-infra/custom-integrations/test-folderSize-definition.yaml
mkdir -p var/db/newrelic-infra/custom-integrations/bin
cp /var/db/newrelic-infra/custom-integrations/bin/test-folderSize var/db/newrelic-infra/custom-integrations/bin/test-folderSize
mkdir -p etc/newrelic-infra/integrations.d/
cp /etc/newrelic-infra/integrations.d/test-folderSize-config.yaml etc/newrelic-infra/integrations.d/test-folderSize-config.yaml
```

And then you can run the tool from the package directory. I used the same package name and version used in the integration to keep the structure. The command to generate

```$ fpm -s dir -t deb --name acme-128 --version 1.0.0 --iteration 1 --description "Folder Size Integration for New Relic Infra" .```

The binary package can be found in the Exercise3 directory

### Exercise 4

To deploy the file into the virtual machine I used the file provisioner, but for deploying it was better using the Ansible local provisioner for Vagrant letting Ansible doing the work of installing the package and refreshing the service.

Since I had chosen Ubuntu 16.04 and the the Ansible local provisioner relies on python2 I needed to add python-minimal before using ansible. With Ubuntu 14.04 this installation would not be needed since python2 is the default

### Exercise 5

#### Docker image

I will be only building a docker container to run the test-folderSize binary every 2 seconds. I am not not deploying newrelic-infra software for this container for this exercise as we did with the previous exercises.

The command to build the docker image is below

```$ docker build -t foldersize:0.1.0 .```

Once the process finishes, we can run docker containers  by running

```$ docker run foldersize:0.1.0```

On each docker run the output of the integration will be on the standard output. I only need to run each container every two seconds. Containers should be as ephemeral as possible so an easy way of doing it would be using the watch command.

```$ watch sudo docker run foldersize:0.1.0 ```

We can see that indeed they were running every two seconds

```
$ docker ps -a
CONTAINER ID        IMAGE               COMMAND                   CREATED              STATUS                           PORTS               NAMES
5ae126efb8f5        foldersize:0.1.0    "/bin/sh -c /var/db/n"    3 seconds ago        Exited (0) 2 seconds ago                             tiny_stonebraker
922b8bd36e09        foldersize:0.1.0    "/bin/sh -c /var/db/n"    5 seconds ago        Exited (0) 4 seconds ago                             focused_yonath
9d6a343134c3        foldersize:0.1.0    "/bin/sh -c /var/db/n"    7 seconds ago        Exited (0) 6 seconds ago                             pensive_agn
```

#### Docker daemon



I set up a Docker Swarm cluster in the past. With this solution you could run docker images remotely by setting up the environment variable docker host pointing to the remote host.

```export DOCKER_HOST="tcp://1.2.3.4:2375"```

That would be a good way to go for a long term solution.
