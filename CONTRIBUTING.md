## Contributing to PingMe
We want to make contributing to this project as easy and transparent as
possible.

## Project structure

- `main` - Contains definitions for the underlying notification services.
  - `cmd/discord` - Discord notification service.
  - `cmd/email` - Email notification service.
  - `cmd/msteams` - Microsoft Teams notification service.
  - `cmd/rocketchat` - RocketChat notification service.
  - `cmd/slack` - Slack notification service.
  - `cmd/telegram` - Telegram notification service
  
- Documentation
  `docs` - Contains the documentation in markdown format.
  - `services.md` If you are adding a new service please add documentaiton to `services.md`.
  - `home.md` Is the main page rendered when docs website is loaded.
  - `install.md` Contains the install instructions for different packages.

  - Checking Locally
    - Docsify is used for documentation rendering from markdown, you can download
    the cli and test locally before opening a pull request.

    Install
    ```bash
    npm i docsify-cli -g
    # yarn global add docsify-cli
    ```
    Serve locally
    ```bash
    docsify serve docs
    ```


## Commits

Commit messages should be well formatted, and to make that "standardized", we
are using Conventional Commits.

```shell

  <type>[<scope>]: <short summary>
     │      │             │
     │      │             └─> Summary in present tense. Not capitalized. No period at the end. 
     │      │
     │      └─> Scope (optional): eg. common, compiler, authentication, core
     │                                                                                          
     └─> Type: chore, docs, feat, fix, refactor, style, or test.
     
```

You can follow the documentation on
[their website](https://www.conventionalcommits.org).

## Pull Requests
We actively welcome your pull requests.

1. Fork the repo and create your branch from `master`.
2. If you've added code that should be tested, add tests.
3. If you've changed APIs, update the documentation.
4. Ensure the test suite passes (`make test`).
5. Make sure your code lints (`make lint`).
6. Make sure your code is well formatted (`make fmt`).

## Issues
We use GitHub issues to track public bugs. Please ensure your description is
clear and has sufficient instructions to be able to reproduce the issue.

## License
By contributing to PingMe, you agree that your contributions will be licensed
under the LICENSE file in the root directory of this source tree.
