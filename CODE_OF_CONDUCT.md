# Getting Started

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   ```

2. **Create a new branch**

   ```bash
   git checkout -b feature/<branch-name>
   ```

3. **Make changes and commit changes:**

   - **Commit changes**

     ```bash
     git add .
     git commit
     ```

   - We use pre-commit hooks that allow us to format code before committing. This ensures that all code is formatted the same way. If your code gets formatted you will need to run `git add .` again before committing to add the formatted code to the commit. You can also run `task format` to format all code.

   - More information on commit messages can be found [here](#commit-messages).

4. **Push changes to GitHub**

   ```bash
   git push
   ```

   or

   ```bash
   git push origin feature/<branch-name>
   ```

5. **Create a pull request**
   - Go to the [repository](https://github.com/GenerateNU/legacy) on GitHub
   - Click on the `Pull requests` tab
   - Click on the `New pull request` button
   - Select the `base` branch as `main`
   - Select the `compare` branch as `feature/<branch-name>`
   - Click on the `Create pull request` button

### Commit Messages

- Commit messages should be in the present tense.
- Keep them short and concise. If necessary, add a longer description in the body of the commit.
- Use the following format for commit messages:

  ```
  <type>: <subject>
  <BLANK LINE>
  <body>
  ```

- The `<type>` can be one of the following:
  - **feat**: A new feature
  - **fix**: A bug fix
  - **docs**: Documentation only changes
  - **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc.)
  - **refactor**: A code change that neither fixes a bug nor adds a feature
  - **perf**: A code change that improves performance
  - **test**: Adding missing tests
  - **chore**: Changes to the build process or auxiliary tools and libraries such as documentation generation

### Pull Requests

- Ensure your pull request has a clear title and a summary of the changes made.
- Describe the problem you're solving or the feature you're adding.
- Mention any related issues or dependencies.
- Ensure your changes don't break any existing functionality, add tests if necessary.
- Request reviews from fellow developers or team members.
