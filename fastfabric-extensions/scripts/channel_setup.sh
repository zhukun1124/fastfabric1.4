#!/bin/bash
source base_parameters.sh

export CORE_PEER_MSPCONFIGPATH=./crypto-config/peerOrganizations/${PEER_DOMAIN}/users/Admin@${PEER_DOMAIN}/msp
export CORE_PEER_ADDRESS=${FAST_PEER_ADDRESS}:7051

peer channel create -o ${ORDERER_ADDRESS}:7050 -c fastfabric -f ./channel-artifacts/channel.tx
peer channel join -b fastfabric.block

export CORE_PEER_ADDRESS=${ENDORSER_ADDRESS}:7051
peer channel join -b fastfabric.block

export CORE_PEER_ADDRESS=${FAST_PEER_ADDRESS}:7051
peer channel update -o ${ORDERER_ADDRESS}:7050 -c fastfabric -f ./channel-artifacts/anchor_peer.tx
