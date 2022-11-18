### Project boilerplate

1. Run `git init`
2. Add the rest of the git credentials
3. Run `make init` to set up the project operational dependencies
4. Create a `go.mod` file
5. Run `make deps` to update dependencies if any

#### Running on the fly
`make run`

#### Generate the bin
`make build` - Will generate a file name like the folder name in `cmd` folder where main files is

`./cmd/app/main.go` -> `./bin/app`