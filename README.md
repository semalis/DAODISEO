# Achilles 1 TESTNET

DAODISEO Sandbox Blockchain Environment

## Achilles app-chain binaries installation (achillesd)

```
go: go version go1.22.4 linux/amd64
name: achilles
```

## Prerequisites

### Install go

```
sudo rm -rvf /usr/local/go/
wget https://golang.org/dl/go1.23.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.23.6.linux-amd64.tar.gz
rm go1.23.6.linux-amd64.tar.gz
```

### Put PATH to ~/.profile

```
nano .profile
```

```
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
```

### Use source ~/.profile

```
source .profile
```

### Check go

```
go version
```

### Install packages

```
sudo apt-get update
sudo apt-get upgrade
sudo apt install mc htop screen git gcc make
```

## Binary building

### Clone source from repo

```
git clone https://github.com/daodiseomoney/Achilles.git
```

### Open source directory

```
cd Achilles/achilles
```

### Build binary

```
make install
```

## Network launch

### Init

```bash:
achillesd init "<moniker-name>" --chain-id ithaca-1
```

### Set minimum-gas-prices = "" in app.toml to minimum-gas-prices = "0.25uodis"

```
sed -i -e "s|^minimum-gas-prices *=.*|minimum-gas-prices = \"0.25uodis\"|" $HOME/.achilles/config/app.toml
```

### Generate keys

```bash:
# To create new keypair - make sure you save the mnemonics!
achillesd keys add <key-name>
```

or

```
# Restore existing odin wallet with mnemonic seed phrase.
# You will be prompted to enter mnemonic seed.
achillesd keys add <key-name> --recover
```

or

```
# Add keys using ledger
achillesd keys show <key-name> --ledger
```

Check your key:

```
# Query the keystore for your public address
achillesd keys show <key-name> -a
```

### Create account to genesis

```
achillesd genesis add-genesis-account <key-name> 1000000uodis
```

### Create GenTX

```
# Create the gentx.
# Note, your gentx will be rejected if you use any amount greater than 1000000uodis.
# Make sure that all participants built their gentx files without typos.

achillesd genesis gentx <key-name> 1000000uodis \
  --pubkey=$(achillesd tendermint show-validator) \
  --chain-id=ithaca-1 \
  --moniker="my-moniker" \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01"
```

### Add all accounts to genesis

```
# Add account addresses of all participants before generating genesis.
# (whose Gentx files you're using to generate genesis)
achillesd genesis add-genesis-account <account-address> 1000000uodis
```

### Generate genesis

```
achillesd genesis collect-gentxs
```

### Start network

```
achillesd start
```

### ****Set Up achilles Service****

Set up a service to allow binary to run in the background as well as restart automatically if it runs into any problems:
```
sudo tee /etc/systemd/system/achilles.service > /dev/null << EOF
[Unit]
Description=Achilles app chain daemon
After=network-online.target
[Service]
Environment="DAEMON_NAME=achillesd"
Environment="DAEMON_HOME=${HOME}/.achilles"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=false"
Environment="DAEMON_LOG_BUFFER_SIZE=512"
Environment="UNSAFE_SKIP_BACKUP=true"
User=$USER
ExecStart=${HOME}/go/bin/achillesd start
Restart=always
RestartSec=3
LimitNOFILE=infinity
LimitNPROC=infinity
[Install]
WantedBy=multi-user.target
EOF
```

And start service:
```
sudo systemctl daemon-reload
sudo systemctl enable achilles 
sudo systemctl restart achilles
```

How you can check the logs
```
sudo journalctl -u achilles -f --output cat
```

How you can check blocks sync
```
curl http://localhost:26657/status | jq -r ".result.sync_info"
```
