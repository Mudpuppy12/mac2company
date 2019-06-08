# Coding Assessment


These positions are hands-on and competency in Docker is assumed. Looking for clean skills around the use of APIs and some basic development tooling.

Please follow these instructions exactly. Please write a program in the language of your choice to query the https://macaddress.io/ MAC address lookup API. Here’s the requirements:

* Command line only – no GUI wanted or needed – please don’t do a GUI
* Must run on linux at the command line, accepting a command line parameter and returning the data desired.
* Any language you choose. You may not use any libraries that implement this API for you – the challenge is to do a simple REST implementation of this protocol yourself. Consider security in your solution, even if it’s just discussing it in your README
* Program must accept a command line parameter of a MAC address
* Program must query the API for a result over the network
* Program must output the Company Name associated with that MAC address in some useful way to the user at the command line
* Program must be “dockerized” – i.e. the work to make it into a docker container must be present
* All source code for the program and Docker file must be in a git repository
* The Git repository must be in your github.com account and publicly accessible (github.com is free for public repos)
* You should assume we will clone your repo and test it personally
* You should “think unix philosophy” when you do this and consider how this kind of tool would build your toolbox and let that guide how you build it
* It’s expected that you provide some kind of README explaining your work
* This is an example of your work. Assume this is something you are delivering to your peers on a team for them to re-use





# Background 

In order to give you a sense of the work this team has done lately we are sharing the blurb below. You will see that the team is doing some very detailed and advanced work. Candidates to join this team need to be highly technical and still able to do hands on work. The example simple project above is something anyone on this team could do almost without thinking. It’s the most basic baseline of expected skills.

Please email us directly if you have any questions.

In recent assignments this team has focused on the “hard, fundamental” problems that cut across products, across the capability framework, across the platform. Things like security, authentication, authorization, log aggregation, alerting… and more. Our job is to define standard patterns for solving general problems and to provide re-usable code examples. Our deliverables are working code, not powerpoint. Here’s some of our recent accomplishments:

* https://kdf.csco.cloud/ - KDF is the standard platform independent tooling developers can use to have a standard CXDP k8s cluster on their notebook for development
* https://github.com/cisco-sso/kubeadm-one - Tool for deploying a remotely-accessible single node kubernetes cluster on linux using kubeadm – POSSIBLE use for IE remote nodes
* https://github.com/cisco-sso/framework-deploy - Core k8s deployment tech used in KDK
* https://github.com/cisco-sso/charts - Standardized helm charts for everything in our tech stack – databases, kafka, elk, etc – stuff your teams use every single day
* https://github.com/cisco-sso/mh - Multi-helm - templating tool to treat sets of helm charts as an atomic unit for a single deployment
* https://github.com/cisco-sso/snapshotpolicy - A Kubernetes controller to periodically back up disk volumes transparently (after SDP deleted it’s own volumes one day)
* https://github.com/cisco-sso/oauth2_proxy - part of a solution that allows users to individually authenticate for cluster admin instead of using shared certs
