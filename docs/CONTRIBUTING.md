# :sparkles: Contributing to Aki

Hello stranger :wave:

> Thanks for taking the time to contribute to this project and improve it.
> I really appreciate this. :slightly_smiling_face:

The following is a set of guidelines for contributing to Aki. These are mostly guidelines, not rules, but still, it
will be great if you followed them. If you feel that something should be changed in this document (or in any other place
of the project) fill free to propose those changes in issue or pull request.

## :raised_eyebrow: What should I know before I get started?

### Main idea of Aki

This library was created to have one place for useful generic types, methods, functions, type constraints and maybe 
some other generally useful things.

### Always releasable

The goal is to have `main` branch always "releasable", that means that while new functional is added or old is updated
no changes should break existing code. If feature was working before update, after update it should work the same or
with expected changed behaviour.

This is main reason why this project doesn't have `dev` branch.

## :monocle_face: How can I contribute?

### Reporting bugs

You can help the project by reporting bugs or issues that you found.
Also, when submitting bugs, please provide as much information as possible.

### Suggesting new features

If you see that something is missing, or you think some functionality can be extended, feel free to propose that. Also,
any part of docs or comments can be improved, and you can create issue or pull request for that.

In case you want some new functionality, adding usage (or even implementation) examples will be great. A full detailed
description will also help a lot.

### Code contribution

Code contribution is also welcomed. You can pick unresolved issue or just create new feature or add documentation,
basically any help will be great. Still note that your code should meet some basic quality and style guidelines
(see below).

How to contribute step by step:

1. Fork repo
2. Clone `git clone https://github.com/<username>/aki.git`
3. Create new branch `git checkout -b my-new-feature`
4. Make your changes, then add them `git add .`
5. Commit `git commit -m ":fire: New feature added"`
6. Push `git push origin my-new-feature`
7. Create pull request in Aki repo

How to run tests & linter locally:

- Run tests: `make test`
- Run linter: `make lint`
    - Install linter: `make lint-install`
- Run both tests and linter: `make pre-commit`

To see full usage of [Makefile](../Makefile) use: `make help` or just `make`.

## :art: Style guidelines

### Commit messages

No specific requirements, but all commit messages should start with an emoji and a capital letter, and verbs should be 
in the past tense. Message should contain a brief description of what you've done, no need for full text, but just
`Fix` won't be enough.

#### What does emoji mean?

It's just for fun and cool looking commit history, but also, it helps to understand what kind of commit is it.

Explanation:

- :city_sunset: `:city_sunset:` Changes related to GitHub repo (README update, etc.)
- :fire: `:fire:` New feature or extending of functionality
- :speech_balloon: `:speech_balloon:` Documentation (or comments) update
- :space_invader: `:space_invader:` Bug fix
- :page_facing_up: `:page_facing_up:` Changes in configs
- :hammer_and_wrench: `:hammer_and_wrench:` Refactoring or reengineering
- :white_check_mark: `:white_check_mark:` Adding or fixing tests
- :envelope_with_arrow: `:envelope_with_arrow:` Updating dependencies
- :umbrella: `:umbrella:` Fixing CI checks
- :jigsaw: `:jigsaw:` Tiny change or fix

#### Commit message examples

Bad:

- `fix`
- `added tests`
- `move function`

Good:

- `👾 Fixed function X`
- `✅ Added unit tests for X`
- `🛠 Moved X from Y to Z`

### Code style

Your code should meet general Go standards, like camel case naming, capitalized abbreviations, Go style comments, etc.
All described can be read in [Effective Go](https://go.dev/doc/effective_go) and strongly recommended following.

Also, using panic is generally not allowed, but may be used in some cases like initialization of handler with nil handle
function or in unit tests. All errors should be handled (returned to the user or logged).

Documentation of new functionality is essential, and should be added before merging pull request.

### CI testing & linters

Your contribution should pass code quality gates, which means that new code should be covered with unit tests at list by
80%, no new code smells should be added and pass linters.

Unit tests should be independent of each other.
