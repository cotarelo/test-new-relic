#Download base image ubuntu 16.04 so we did with Vagrant
FROM ubuntu:16.04

# Copy to the container
COPY acme-128_1.0.0-1_amd64.deb /tmp/acme-128_1.0.0-1_amd64.deb

# Install the package to run the binary
RUN dpkg -i /tmp/acme-128_1.0.0-1_amd64.deb

# Run the binary file
CMD /var/db/newrelic-infra/custom-integrations/bin/test-folderSize
