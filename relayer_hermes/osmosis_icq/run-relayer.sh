#!/bin/sh
rm -rf ~/.hermes

hermes --config config.toml health-check

#hermes --config config.toml keys delete --chain wasmapp --key-name alice
#hermes --config config.toml keys delete --chain localosmosis-b --key-name bob

hermes --config config.toml keys add --key-name alice --chain wasmapp --mnemonic-file <(echo "lonely nest leisure nose arrange apple burst manage asset arctic property economy nest wall few royal ritual october omit earn base anchor globe chase")
hermes --config config.toml keys add --key-name bob --chain osmo-test-5 --mnemonic-file <(echo "picture switch doctor police industry drill insane law quick best festival suggest soul mammal mammal lava audit morning glove teach valve eyebrow hurt cause")

hermes --config config.toml keys list --chain wasmapp
hermes --config config.toml keys list --chain osmo-test-5

hermes --config config.toml create connection --a-chain wasmapp --b-chain osmo-test-5

hermes --config config.toml create channel --a-port  wasm.cosmos1rl8su3hadqqq2v86lscpuklsh2mh84cxqvjdew4jt9yd07dzekyqq0tv5k --b-port icqhost --a-chain wasmapp --channel-version icq-1 --a-connection connection-0

hermes --config config.toml start




