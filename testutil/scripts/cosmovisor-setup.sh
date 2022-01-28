# Setting up cosmovisor

# download the cosmovisor binary
go get github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor
# check the version (it should print DAEMON_NAME is not set)
cosmovisor version

echo "# Setup Cosmovisor" >> ~/.profile
echo "export DAEMON_NAME=chihuahuad" >> ~/.profile
echo "export DAEMON_HOME=$HOME/.chihuahua" >> ~/.profile
source ~/.profile

# verify by running the following
echo $DAEMON_NAME

mkdir -p ~/.chihuahua/cosmovisor/upgrades
mkdir -p ~/.chihuahua/cosmovisor/genesis/bin/
cp $(which chihuahuad) ~/.chihuahua/cosmovisor/genesis/bin/
# verify by running the following (it should output the same version is running chihuahuad version):
cosmovisor version

export DAEMON_RESTART_AFTER_UPGRADE=true


# Preparing for v1.1.1 upgrade at block 535000

cd ~/chihuahua
git fetch --tags
git checkout v1.1.1
make install

# check the version - should return v1.1.1
chihuahuad version

# create the upgrade directory if you haven't
mkdir -p $DAEMON_HOME/cosmovisor/upgrades/angryandy/bin

# if you are uysing cosmovisor you then need to copy this new binary
cp /home/<YOUR_USER>/go/bin/chihuahuad $DAEMON_HOME/cosmovisor/upgrades/angryandy/bin

cosmovisor run start




# reset everything to test again

chihuahuad unsafe-reset-all
rm -rf ~/.chihuahua
chihuahuad config chain-id test

# run single-validator-node.sh