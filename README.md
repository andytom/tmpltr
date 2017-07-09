tmpltr
======
[![Build Status](https://travis-ci.org/andytom/tmpltr.svg?branch=master)](https://travis-ci.org/andytom/tmpltr)

Overview
--------

A command line tool to create files and directories from a template.

### Features
* **Single Binary** - No dependencies to install just a single binary to add to
  your PATH and run
* **Powerful Templates** - Templates are written in the powerful [Golang
  Template language](https://golang.org/pkg/text/template/)
* **Template Paths and Contents** - You can template both the path and the
  content of files
* **Simple Interactive UI** - Easy to use UI based on
  [AlecAivazis/survey](https://github.com/AlecAivazis/survey)

Getting Started
---------------

### Install

Download the pre-compiled binary for your OS from the
[release page](https://github.com/andytom/tmpltr/releases/latest) and copy it
into your path.

### Usage

You can use the built in help to list the usage:
```bash
tmpltr help
```

Writing Templates
-----------------

See [Writing a Template](docs/writing_templates.md) for a guide to write a
template.

There are example templates in the `examples` directory.

Getting Involved
----------------

### Spotted a Bug? Need a Features?

If you notice a bug or have a feature suggestion please create a github issue.

### Contributing

If you want to help contribute please see `CONTRIBUTING.md` for more
information on requirements for PRs and other useful information.
