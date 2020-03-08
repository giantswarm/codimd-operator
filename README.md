# codimd-operator

This operator is currently for testing & teaching purposes.

The idea is to read markdown from [codiMD](https://github.com/hackmdio/codimd) pages and create any `yaml` descriptions of kubernetes deployments that can be found in code blocks.


### Super quick test guide

1. Create `kind` cluster locally
2. `make install`
3. `make run` (in separate commandline window)
4. `kubectl apply -f config/samples/`

### Operator workshop

Link to notes: https://hackmd.okfn.de/vAnaZtK7RuCMcf9xU8Ok7A
