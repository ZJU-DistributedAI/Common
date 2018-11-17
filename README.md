# Common client for distributed AI on blockchain

There are three types of users in our platform. They are: data providers, model providers and execution resource 
providers. This project provides general API services that could be consumed by any of the three types of users. For 
instance, IPFS file upload APIs, data quality management APIs and many more.

## Run as a server using Docker

- (CAUTION) Run `./_clean.sh` to prune all Docker images, containers, networks, volumes.
- Run `./_build.sh` to build the docker image
- Run `./start.sh` to start the sample servers

`start.sh` will firstly start two IPFS nodes (ipfs1 and ipfs2). Then it will wait until all ports are available on 
the two ipfs nodes. Next it will start two cluster nodes (cluster1 and cluster2), and connect the two cluster nodes 
(using raft). Lastly it will run (common1 and common2).

## PORTS
- IPFS1

      - "4001:4001" # ipfs swarm
      - "5001:5001" # expose if needed/wanted
      - "6001:8080" # exposes if needed/wanted
      
- CLUSTER1

      - "9094:9094" # API
      - "9096:9096" # Cluster IPFS Proxy endpoint
      
- COMMON1

      - "3001:3001" # distributed storage server with swagger

- IPFS2

      - "4101:4001" # ipfs swarm
      - "5101:5001" # expose if needed/wanted
      - "6101:8080" # exposes if needed/wanted
      
- CLUSTER2

      - "9194:9094" # API
      - "9196:9096" # Cluster IPFS Proxy endpoint
      
- COMMON2

      - "3101:3001" # distributed storage server with swagger
      
      
# SWAGGER-UI

Friendly swagger-ui is available at `common1:3001/swagger/index.html` and `common2:3101/swagger/index.html`