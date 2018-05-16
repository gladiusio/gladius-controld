# Gladius Control Daemon

See the main [gladius-node](https://github.com/gladiusio/gladius-node) repository to see more.

## Cross compile
To compile for other systems you need to install [xgo](https://github.com/karalabe/xgo).
This is because of the Ethereum CGO bindings.

Run `make dependencies`

Then run `xgo --targets="windows/*,darwin/*,linux/*" --dest="./build/" ./cmd/gladius-controld`
from the project root. You can change the target to be whatever system you want.
