# New Relic Infrastructure - SRE/Build engineer

## Prerequisites
Virtualbox >= 5.1
Vagrant >= 2.0

### Exercise 1
I have chosen to install Ubuntu 16.04 as a happy path. . Go to the Exercise1 directory and run the command below to start the virtual machine

```vagrant up```

The following command to provision the virtual machine with New Relic Infraestructure software however it provisions by default on the first run  

```vagrant provision```

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

```fpm -s dir -t deb --name acme-128 --version 1.0.0 --iteration 1 --description "Folder Size Integration for New Relic Infra" .```

The binary package can be found in the Exercise3 directory

### Exercise 4
