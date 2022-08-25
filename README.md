# Supply-Chain-using-Hyperledger

## Supply Chain Management for APMC Using Hyperledger in Blockchain

![image](https://user-images.githubusercontent.com/46996519/170439070-94b66ce6-5b5e-4765-a598-cf727b087208.png)




Developed by
Ankita Shelke,Smit Bhatt, Jainam Shah for our final year engineering project.
Project Description

A peer to peer  decentralized system for tracebility of farmer produce from the producer to the consumer involving all stakeholders of the supply chain. we have Built a permissioned blockchain in Hyperledger fabric for managing the supply chain of APMC market. Integrate with an Android mobile application built in flutter to track the fruit produced , and check its authenticity. 


Setup Blockchain Network in Hyperledger Fabric

- Create organization, Peers, Users and Orderers with either cryptogen or CA (edit file create-artifacts.sh and cryoto-config.yaml ) 
- Set policies for Org (Signature for ORG or implicit meta policy for application)
  Anchor VS Leader peer?
  configtx.yaml -> contains who is anchor in each, policies for each etc  
  channel/config/configtx.yaml ->
- Create genesis channel anchor peers (edit create-artifacts.sh) 
- Up the network docker-compose up -d (might have to stop all service in docker to avoid host taken error)
- Create a channel and add peers to the channel, make changes to the anchor. 
  Leader peer - to get order and distribute it among the peers in its organization. Can we choose statically(meaning only one peer will be assigned at start) or dynamically chosen among available peers.
  Anchor peer - To introduce the peers in different orgs so later can communicate directly without the anchor peer. 
- Chaincode Cycle- > Deploy Chaincode
	Go file Imports 
  go mod init github.com/fabcar/go ->create go.mod 
  go mod tidy  -> add required import in go.mod, will also generate go.sum



Start the Server

- Navigate to the api-1.4 directory and run the following command to start the node server
  nodemon app.js
- Use fabric client to make api’s (running on 4000 port)
- In a separate terminal use localhost.run to host the server and get url to make calls from any device.
   ssh-keygen -t rsa -b 2048 -C "<comment>"
   ssh -R 80:localhost:4000 localhost.run
 
 
Building Flutter Android Application 
 
- To install all the modules run the command 
  flutter pub get
- To built the application in android device run the command 
  flutter run 

