# Kohirens STDLIB

A collection of highly reusable code throughout Korhirens projects.

## Setup Local Development.

You can use the Docker environment to get going if you have Docker on your
computer. In fact, there is no documentation for any other way in this reading.

You will need to set the `$Env:HOME` environment on Windows before you can start
the dev environment. It should be the same value as `$Env:USERPROFILE`, or
what you need it to be. Once you have it set and it shows up in Powershell when
you run `Get-ChildItem Env:` you then do one of the following.

### Run Docker

1. Clone this repository.
2. In a command prompt, run: `docker-compose -f .docker/docker-compose.yml up`
3. In another command prompt, login to the container: 
   `docker exec -it stdlib_dev_1 sh`
4. Execute a command such as `go test`
   ```output
   ~/src/github.com/kohirens/stdlib $ go test
   PASS
   ok      github.com/kohirens/stdlib      0.004s
   ```

### Run with VS Code

1. Install the VS code extention "Remote Container".
2. Clonet this repository localy.
3. Open the project in VS Code, which should ask to open the folder in a
   remote container.
4. Open a termianl in VS Code and type `go test`.
5. Now got to "Run and Debug" in the left nav and run one of the launch
   configurations.
