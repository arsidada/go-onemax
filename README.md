# Go backend for OneMAX.org

## Development Instructions

- Ensure that you have Go installed
    ```bash
    $ go version
    go version go1.10.3 darwin/amd64
    ```
- Clone the repo to your `$GOPATH/src/github.com/arsidada`. For example:
    ```bash
    $ echo $GOPATH
    /Users/arsalan/dev/max/go

    $ pwd
    /Users/arsalan/dev/max/go/src/github.com/arsidada/go-onemax
    ```
- Add env variables for DB credentials. Env vars required are:
  1. `DB_USER`
  2. `DB_PASSWORD`
  3. `DB_ADDR`
  4. `DB_DB`
  
- Start the server locally by running `go run server.go`

- Switch to a new branch before making any changes: `git branch -D branch-name`

- Once changes are complete and tested, create a Pull Request and assign a reviewer for change review

## Deployment Instructions

- Ensure that the heroku app is configured as a remote git repo:

  `git remote add heroku https://git.heroku.com/fierce-everglades-88127.git`
  
- Push to the heroku app by performing the following:

  `git push heroku master`
