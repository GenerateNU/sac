
# Package Managers

- Mac - [Homebrew](https://brew.sh/)
- Windows - get [WSL](https://docs.microsoft.com/en-us/windows/wsl/install-win10)
- Linux (you know what you're doing)

# Prerequisites

- [Git](https://git-scm.com/)
- [Node.js](https://nodejs.org/en/)
- [Yarn](https://yarnpkg.com/)
- [Go](https://golang.org/)
  - Do not install through brew, use the official website
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)
  - Install through brew: `brew install postgresql@15`
  - It requires you to add all the exports to path so read the end of the installation carefully!
- [Trunk](https://marketplace.visualstudio.com/items?itemName=Trunk.io) (Recommended!)
  - Visual Studio Code extension for linting/formatting
- [gofumpt](https://github.com/mvdan/gofumpt)
  - A stricter gofmt
- [golangci-lint](https://golangci-lint.run/welcome/install/)
  - A Go linters aggregator

# Setup

1. **Clone the repository**

   ```console
   git clone git@github.com:GenerateNU/sac.git
   ```

2. **Install dependencies**

   ```console
   cd frontend/* 
   yarn install
   ```

   - If you get an error about `expo-cli` not being installed, run `yarn global add expo-cli` and then run `yarn install` again.

   ```console
   cd server
   go get ./...
   ```

   - If this doesnt work, try running `go mod tidy` and then `go get ./...` again or delete the go.mod and go.sum files and then run `go mod init backend` and `go mod tidy` again.

## React Native Builds

1. **Create client build**

   ```console
   cd frotend/sac-mobile
   eas login
   eas build:configure
   # ios
   eas build -p ios --profile development
   # android
   eas build -p android --profile development

   ```

2. **Download the build and drag into simulator**

3. **Start the client**

   ```console
   cd frontend/sac-mobile
   npx expo start --dev-client
   ```

   - You can then open the app in the Expo app in the simulator.

### Postgresql Setup

1. **Turn on postgresql**

   - MacOS

   ```console
   brew services start postgresql@15
   ```

   - Windows

   ```console
   pg_ctl -D /usr/local/var/postgres start
   ```

2. **Create a user**

   ```console
   createdb
   ```

3. **Create a database**

   ```console
   psql // opens psql shell
   CREATE DATABASE sac;
   ```

4. **Create a user**

   ```console
   createuser postgres -U <your username>
   ```

# Commands

## React Native

  ```console
   npx expo start --dev-client // runnning dev client
   npx expo start --dev-client --ios // specific platform
   yarn format // format code
   yarn lint // lint code
   yarn test // run tests
   ```

## Go

   ```console
   go run main.go // run server
   go test ./... // run tests
   gofumpt -l -w . // format code
   golangci-lint run --fix // lint code
   ```

### SAC CLI

   To install use `./install.sh` and then run `sac-cli` to see all commands.

# Git Flow

1. **Create a new branch**

   ```console
   git checkout -b <branch-name> // this is determined by your ticket name
   ```

2. **Make changes and commit changes:**

   - **Commit changes**

     ```console
     git add .
     git commit
     ```

   - We use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for commit messages. (READ!)

3. **Push changes to GitHub**

   ```console
   git push
   ```

   or

   ```console
   git push origin <branch-name>
   ```

4. **Create a pull request**
   - Go to the [repository](https://github.com/GenerateNU/sac) on GitHub
   - Click on the `Pull requests` tab
   - Click on the `New pull request` button
   - Select the `base` branch as `main`
   - Select the `compare` branch as `<branch-name>`
   - Click on the `Create pull request` button

5. **Issues**

   Use the Issues tab to create issues if you find any bugs during development or can't find a feature you working on.
