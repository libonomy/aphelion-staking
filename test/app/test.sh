#! /bin/bash
set -ex

#- kvstore over socket, curl
#- counter over socket, curl
#- counter over grpc, curl
#- counter over grpc, grpc

# TODO: install everything

export PATH="$GOBIN:$PATH"
export TMHOME=$HOME/.aphelion_app

function kvstore_over_socket(){
    rm -rf $TMHOME
    aphelion init
    echo "Starting kvstore_over_socket"
    abci-cli kvstore > /dev/null &
    pid_kvstore=$!
    aphelion node > aphelion.log &
    pid_aphelion=$!
    sleep 5

    echo "running test"
    bash kvstore_test.sh "KVStore over Socket"

    kill -9 $pid_kvstore $pid_aphelion
}

# start aphelion first
function kvstore_over_socket_reorder(){
    rm -rf $TMHOME
    aphelion init
    echo "Starting kvstore_over_socket_reorder (ie. start aphelion first)"
    aphelion node > aphelion.log &
    pid_aphelion=$!
    sleep 2
    abci-cli kvstore > /dev/null &
    pid_kvstore=$!
    sleep 5

    echo "running test"
    bash kvstore_test.sh "KVStore over Socket"

    kill -9 $pid_kvstore $pid_aphelion
}


function counter_over_socket() {
    rm -rf $TMHOME
    aphelion init
    echo "Starting counter_over_socket"
    abci-cli counter --serial > /dev/null &
    pid_counter=$!
    aphelion node > aphelion.log &
    pid_aphelion=$!
    sleep 5

    echo "running test"
    bash counter_test.sh "Counter over Socket"

    kill -9 $pid_counter $pid_aphelion
}

function counter_over_grpc() {
    rm -rf $TMHOME
    aphelion init
    echo "Starting counter_over_grpc"
    abci-cli counter --serial --abci grpc > /dev/null &
    pid_counter=$!
    aphelion node --abci grpc > aphelion.log &
    pid_aphelion=$!
    sleep 5

    echo "running test"
    bash counter_test.sh "Counter over GRPC"

    kill -9 $pid_counter $pid_aphelion
}

function counter_over_grpc_grpc() {
    rm -rf $TMHOME
    aphelion init
    echo "Starting counter_over_grpc_grpc (ie. with grpc broadcast_tx)"
    abci-cli counter --serial --abci grpc > /dev/null &
    pid_counter=$!
    sleep 1
    GRPC_PORT=36656
    aphelion node --abci grpc --rpc.grpc_laddr tcp://localhost:$GRPC_PORT > aphelion.log &
    pid_aphelion=$!
    sleep 5

    echo "running test"
    GRPC_BROADCAST_TX=true bash counter_test.sh "Counter over GRPC via GRPC BroadcastTx"

    kill -9 $pid_counter $pid_aphelion
}

cd $GOPATH/src/github.com/evdatsion/aphelion-dpos-bft/test/app

case "$1" in 
    "kvstore_over_socket")
    kvstore_over_socket
    ;;
"kvstore_over_socket_reorder")
    kvstore_over_socket_reorder
    ;;
    "counter_over_socket")
    counter_over_socket
    ;;
"counter_over_grpc")
    counter_over_grpc
    ;;
    "counter_over_grpc_grpc")
    counter_over_grpc_grpc
    ;;
*)
    echo "Running all"
    kvstore_over_socket
    echo ""
    kvstore_over_socket_reorder
    echo ""
    counter_over_socket
    echo ""
    counter_over_grpc
    echo ""
    counter_over_grpc_grpc
esac

