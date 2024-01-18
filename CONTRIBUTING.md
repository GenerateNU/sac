
# Package Managers

- Mac - [Homebrew](https://brew.sh/)
- Windows - get [WSL](https://docs.microsoft.com/en-us/windows/wsl/install-win10)
- Linux (you know what you're doing)

# Prerequisites

- [Git](https://git-scm.com/)
- [Node.js](https://nodejs.org/en/)
- [Yarn](https://yarnpkg.com/)
- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)
  - Install through brew: `brew install postgresql@15`
  - It requires you to add all the exports to path so read the end of the installation carefully!
- [Trunk](https://marketplace.visualstudio.com/items?itemName=Trunk.io) (Recommended!)
  - Visual Studio Code extension for linting/formatting

# Setup

1. **Clone the repository**

   ```bash
   git clone git@github.com:GenerateNU/sac.git
   ```

2. **Install dependencies**

   ```bash
   cd frontend/* 
   yarn install
   ```

   - If you get an error about `expo-cli` not being installed, run `yarn global add expo-cli` and then run `yarn install` again.

   ```bash
   cd server
   go get ./...
   ```

   - If this doesnt work, try running `go mod tidy` and then `go get ./...` again or delete the go.mod and go.sum files and then run `go mod init backend` and `go mod tidy` again.

### React Native Builds

1. **Create client build**

   ```bash
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

   ```bash
   cd frontend/sac-mobile
   npx expo start --dev-client
   ```

   - You can then open the app in the Expo app in the simulator.

### Postgresql Setup

1. **Turn on postgresql**

   - MacOS

   ```bash
   brew services start postgresql@15
   ```

   - Windows

   ```bash
   pg_ctl -D /usr/local/var/postgres start
   ```

2. **Create a user**

   ```bash
   createdb
   ```

3. **Create a database**

   ```bash
   psql // opens psql shell
   CREATE DATABASE sac;
   ```

4. **Create a user**

   ```bash
   createuser postgres -U <your osusername>
   ```

# Commands

### React Native

  ```bash
   npx expo start --dev-client // runnning dev client
   npx expo start --dev-client --ios // specific platform
   yarn format // format code
   yarn lint // lint code
   yarn test // run tests
   ```

### Go

   ```bash
   go run main.go // run server
   go test ./... // run tests
   go fmt ./... // format code
   go vet ./... // lint code
   ```

### Others (WIP)

   ```bash
   sac-cli migrate // run migrations
   sac-cli reset // reset database
   sac-cli swagger // generate swagger docs
   sac-cli lint // lint code
   sac-cli format // format code
   sac-cli test // run tests
   ```

# Git Flow

1. **Create a new branch**

   ```bash
   git checkout -b <branch-name> // this is determined by your ticket name
   ```

2. **Make changes and commit changes:**

   - **Commit changes**

     ```bash
     git add .
     git commit
     ```

   - We use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for commit messages. (READ!)

   <!-- - We especially recommend [Trunk](https://marketplace.visualstudio.com/items?itemName=Trunk.io) for linting -->

3. **Push changes to GitHub**

   ```bash
   git push
   ```

   or

   ```bash
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
