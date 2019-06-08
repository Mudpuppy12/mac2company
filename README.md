# mac2company
A simple go client that speaks to https://macaddress.io/ and retrieves the company name in which the MAC address is registered.
The client is built with golang and using docker as an example on containerizing toolsets.


# What you need to get started!
* Docker installed on your machine
* Git > 1.8.x
* An API key for https://macaddress.io

# Build instructions
With API key in hand, use the following docker command to build the tool container.
<pre>
$ docker build --build-arg APIKEY=<API_KEY_HERE> https://github.com/Mudpuppy12/mac2company.git#master:docker -t mac2company
</pre>

If you have an older git version (Centos 7.5 had 1.8.X) less than 2.X you may
get an error on the above command. Update your git! 

See: 
* https://computingforgeeks.com/how-to-install-latest-version-of-git-git-2-x-on-centos-7 and
* https://github.com/moby/moby/issues/36965

# Running the tool 
<pre>
$ docker run --rm -ti mac2company 
</pre>
# Getting help

<pre>
$ docker run --rm -ti mac2company -h
</pre>

# Doing a company lookup lookup

<pre>
$ docker run --rm -ti mac2company lookup --mac=08:00:27:6f:f7:15
</pre>

# Version history

<pre>
$ docker run --rm -ti mac2company version
</pre>
