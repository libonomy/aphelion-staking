#! /bin/bash
set -eu

ID=$1
DOCKER_IMAGE=$2
NODEID="$(docker run --rm -e TMHOME=/go/src/github.com/libonomy/aphelion-staking/test/p2p/data/mach$((ID-1)) $DOCKER_IMAGE aphelion show_node_id)"
echo "$NODEID@172.57.0.$((100+$ID))"
