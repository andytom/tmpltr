Writing a Template
==================

Initial Setup
-------------

Create a new directory for the purposes of this example we will use the name
`example`.

```bash
mkdir example
cd example
```

Next create a file called `config.yaml`. The config file contains metadata
about the template and along with questions that will be asked when using the
template. For now just add the following:

```yaml
meta:
  description: An example template
```

Add your first file
-------------------

Create a directory called `template`. All files and sub-directories of this
directroy are created in the target directory when the template is used. The
files are treated as templates and will be executed and the contents writted to
the newly created files. The paths for all the sub directories will also be
treated as templates allowing the creation of custom paths.

```bash
mkdir template
```

In that directory create a file called `hello.txt` with the contents.

```
Hello World!
```

If you used this template now it would just create a file called `hello.txt` in
directory where you call the template. The contents will be the same everytime
you call it.

Ask a question
--------------

To make your template dynamic you need to ask the user some questions. This is
set up in the `config.yaml`. To start with we will add the following to your
`config.yaml`.

```yaml
questions:
  - name: username
    type: input
    message: What is your name?
    required: true
```

The `name` is the name that this value will be made available in the template.
In this example it will be available as `.username` (note the prefix `.`).

The `type` is the type of question to ask, in this case it is a basic text
input.

The `message` is the prompt to display when asking the question.

Finaly there is `required` that indicates that this question is required. This
is optional an defaults to `false`.

Update the file
---------------

Now that we know the user's name we can greet them.

Change the `template/hello.txt` and replace `World` with `{{ .username }}` so
that the contents of the file now look like

```
Hello {{ .username }}!
```

If you executed the template now it will prompt you for your name and create a
file called `hello.txt` that greets you by name.


Install your template
---------------------

You can now install your template using the `install` command:

```bash
$ tmpltr install example .
Installing template from "." as "example"
```

You can now list the template using the `list` command:

```bash
$ tmpltr list
+---------+---------------------+
|  NAME   |     DESCRIPTION     |
+---------+---------------------+
| example | An example template |
+---------+---------------------+
```

The final version of this template can be found in `examples/example` in the
git repo.
