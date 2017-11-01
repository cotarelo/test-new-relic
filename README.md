# New Relic Infrastructure - SRE/Build engineer

## Prerequisites
Virtualbox >= 5.1
Vagrant >= 2.0

### Exercise 1
I have chosen to install Ubuntu 16.04 as a happy path. Run the command below to start the virtual machine.

```vagrant up```

Run the following command to provision the virtual machine with New Relic Infraestructure software

```vagrant provision```

The software is installed according the instructions in https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure/installation/install-infrastructure-linux

It will use Vagrant shell provisioner to deploy and configure the software. Note that I used shell provisioner for simplicity taking into consideration that is only one host, when deploying the software in multiple hosts I would use a puppet or ansible provisioner.

After installing the sofware it can be seen how the data is sent to New Relic under the user Acme_125 https://rpm.newrelic.com/accounts/1785110
