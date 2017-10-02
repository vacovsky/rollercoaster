#!/bin/bash

./beads \
    --key-names="name,key" \
    --value-names="hostAddress,connectionString,value" \
    --exclude="Microsoft.,System." \
    --config="config/configuration.json"
