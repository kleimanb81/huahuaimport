#!/bin/sh

# --chain-id test --broadcast-mode block --keyring-backend test --from validator

chihuahuad tx gov submit-proposal software-upgrade spirit-of-the-wolf --chain-id test --broadcast-mode block --keyring-backend test --from validator


sleep 3


chihuahuad tx gov vote 1 VOTE_OPTION_YES --chain-id test --broadcast-mode block --keyring-backend test --from validator