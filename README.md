# Sense

Sense (working name) is a system which used to manage a network of sensing Nodes which have been deployed to the environment.

## Server

Written in Go. Uses a PostgreSQL database. 

* Serves a RESTful JSON api over HTTP
* Subscribes to MQTT topics from an MQTT broker called Mosquitto

