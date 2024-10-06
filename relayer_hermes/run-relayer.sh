#!/bin/sh
rm -rf .hermes

hermes --config config.toml health-check

#hermes --config config.toml keys delete --chain wasmapp --key-name alice
#hermes --config config.toml keys delete --chain localosmosis-b --key-name bob

hermes --config config.toml keys add --key-name alice --chain wasmapp --mnemonic-file <(echo "lonely nest leisure nose arrange apple burst manage asset arctic property economy nest wall few royal ritual october omit earn base anchor globe chase")
hermes --config config.toml keys add --key-name bob --chain localosmosis-b --mnemonic-file <(echo "increase bread alpha rigid glide amused approve oblige print asset idea enact lawn proof unfold jeans rabbit audit return chuckle valve rather cactus great")

hermes --config config.toml keys list --chain wasmapp
hermes --config config.toml keys list --chain localosmosis-b

hermes --config config.toml create connection --a-chain wasmapp --b-chain localosmosis-b
hermes --config config.toml create channel --a-port interquery --b-port icqhost --a-chain wasmapp --channel-version icq-1 --a-connection connection-0

hermes --config config.toml start

