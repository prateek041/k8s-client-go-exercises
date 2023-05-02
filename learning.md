# Understanding the Kubernetes Go client

## Section 1: Connecting to the cluster

The kubernetes API server is the central place that manages the state of the cluster and also exposes an HTTP REST API for the clients to interact with it. Since it is a REST API, we need a client to connect with it. That client is build using a config file. That config file contains the necessary information like name, location etc where the cluster API server can be found.

So, we use that config to build a clientSet that can be used to interact with the K8s. Now you need some basic understanding of how the API is planned and versioned to better understand all of listing and stuff.