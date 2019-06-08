# mac2company
A simple go client that speaks to https://macaddress.io/ and retrieves the company name in which the MAC address is registered.
The client is built with golang and using docker as an example on containerizing toolsets. It uses a two step build process to try
to keep image size small.

# What you need to get started!
* Docker installed on your machine
* Git > 1.8.x
* An API key from https://macaddress.io

# Build instructions
With API key in hand, use the following docker command to build the tool container.
<pre>
$ docker build --build-arg APIKEY=API_KEY_HERE https://github.com/Mudpuppy12/mac2company.git#master:docker -t mac2company
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

# Doing a company lookup

<pre>
$ docker run --rm -ti mac2company lookup --mac=08:00:27:6f:f7:15
</pre>

# Version history

<pre>
$ docker run --rm -ti mac2company version
</pre>

# Security
The API key is used during the image build process and is embedded into the config.toml file that the program uses. Keep this in mind as each run
will use the same APIKEY in which was used at build time.


# Other comments

This project was built as an assessment on coding and docker knowledge. The attempt was made to make the tool configurable down the road using a config.toml file should API endpoint change and also as a template for anyone else interested in learning coding.