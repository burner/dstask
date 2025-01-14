[![CircleCI](https://circleci.com/gh/naggie/dstask.svg?style=svg)](https://circleci.com/gh/naggie/dstask)

[![Go Report Card](https://goreportcard.com/badge/github.com/naggie/dstask)](https://goreportcard.com/report/github.com/naggie/dstask)

# dstask

A personal task tracker designed to help you focus.

Dstask is currently in beta -- the interface, data format and commands may
change before version 1.0. That said, it's unlikely that there will be a
breaking change as things are nearly finalised.

Features:

 * Powerful context system
 * **Git powered sync**/undo/resolve (passwordstore.org style) which means no need to set up a sync server, and sync between devices is easy!
 * Task listing won't break with long task text
 * `open` command -- **open URLs found in specified task** in the browser
 * `note` command -- edit a **full markdown note** for a task
 * zsh/bash completion for speed

Non-features:

 * Collaboration. This is a personal task tracker. Use another system for
   projects that involve multiple people. Note that it can still be beneficial
   to use dstask to track what you are working on in the context of a
   multi-person project tracked elsewhere.

Requirements:

* Git
* A 256-color capable terminal

<p align="center">
  <img src="https://github.com/naggie/dstask/raw/master/etc/dstask.png">
</p>

# Installation

1. Copy the executable (from the [releases page][1]) to somewhere in your path, named `dstask` and mark it executable. `/usr/local/bin/` is suggested.
1. Enable bash completions by copying `.bash-completion.sh` into your home directory and sourcing it from your `.bashrc`. There's also a zsh completion script.
1. Set up an alias in your `.bashrc`: `alias task=dstask` or `alias n=dstask` to make task management slightly faster.

# Moving from Taskwarrior

Before installing dstask, you may want to export your taskwarrior database:

    task export > taskwarrior.json

After un-installing taskwarrior and installing dstask, to import the tasks to dstask:

    dstask import-tw < taskwarrior.json


Commands and syntax are deliberately very similar to taskwarrior. Here are the exceptions:

  * The command is (nearly) always the first argument. Eg, `task eat some add bananas` won't work, but `task add eat some bananas` will. If there's an ID, it can proceed the command but doesn't have to.
  * Priorities are added by the keywords `P0` `P1` `P2` `P3`. Lower number is more urgent. Default is `P2`. For example `task add eat some bananas P1`. The keyword can be anywhere after the command.
  * Action is always the first argument. Eg, `task eat some add bananas` won't work, but `task add eat some bananas` will.
  * Contexts are defined on-the-fly, and are added to all new tasks if set. Use `--` to ignore current context in any command.

[1]: https://github.com/naggie/dstask/releases/latest

# Major things missing

There are a few things missing at the moment. That said I use dstask day to day and trust it with my work.

* Recurring tasks
* Subtask implementation (github issue style or otherwise)
* Deferring tasks
* Due dates
* Advanced reports
* Task dependencies

# Usage

```
Usage: dstask [id...] <cmd> [task summary/filter]

Where [task summary] is text with tags/project/priority specified. Tags are
specified with + (or - for filtering) eg: +work. The project is specified with
a project:g prefix eg: project:dstask -- no quotes. Priorities run from P3
(low), P2 (default) to P1 (high) and P0 (critical). Text can also be specified
for a substring search of description and notes.

Cmd and IDs can be swapped, multiple IDs can be specified for batch
operations.

run "task help <cmd>" for command specific help.

Add -- to ignore the current context. / can be used when adding tasks to note
any words after.

Available commands:

next           : Show most important tasks (priority, creation date -- truncated and default)
add            : Add a task
log            : Log a task (already resolved)
start          : Change task status to active
note           : Append to or edit note for a task
stop           : Change task status to pending
done           : Resolve a task
context        : Set global context for task list and new tasks
modify         : Set attributes for a task
edit           : Edit task with text editor
undo           : Undo last action with git revert
sync           : Pull then push to git repository, automatic merge commit.
open           : Open all URLs found in summary/annotations
git            : Pass a command to git in the repository. Used for push/pull.
show-projects  : List projects with completion status
show-tags      : List tags in use
show-active    : Show tasks that have been started
show-paused    : Show tasks that have been started then stopped
show-open      : Show non-resolved tasks (without truncation)
show-resolved  : Show resolved tasks
import-tw      : Import tasks from taskwarrior via stdin
help           : Get help on any command or show this message
```

# Syntax


## Priority

| Symbol | Name      | Note                                                    |
|--------|-----------|---------------------------------------------------------|
| `P0`   | Critical  | Must be resolved immediately                            |
| `P1`   | High      |                                                         |
| `P2`   | Normal    | Default priority                                        |
| `P3`   | Low       | Shown at bottom and faded.                              |


## Operators

| Symbol      | Syntax               | Description                                          | Example                                     |
|-------------|----------------------|------------------------------------------------------|---------------------------------------------|
| `+`         | `+<tag>`             | Include tag. Filter/context, or when adding task.    | `task add fix server +work`                 |
| `-`         | `-<tag>`             | Exclude tag. Filter/context only.                    | `task next -feature`                        |
| `--`        | `--`                 | Ignore context. When listing or adding tasks.        | `task --`, `task add -- +home do guttering` |
| `/`         | `/`                  | When adding a task, everything after will be a note. | `task add check out ipfs / https://ipfs.io` |
| `project:`  | `project:<project>`  | Set project. Filter/context, or when adding task.    | `task context project:dstask`               |
| `-project:` | `-project:<project>` | Exclude project, filter/context only.                | `task next -project:dstask -work`           |


# State

| State    | Description                                   |
|----------| ----------------------------------------------|
| Pending  | Tasks that have never been started            |
| Active   | Tasks that have been started                  |
| Paused   | Tasks that have been started but then stopped |
| Resolved | Tasks that have been done/close/completed     |


# A note on performance

Currently I'm using dstask to manage thousands of tasks and the interface still
appears instant.

Dstask currently loads and parses every non-resolved task, each task being a
single file. This may sound wasteful, but it allows for a simple design and is
actually performant thanks to modern OS disk caches and SSDs.

If it starts to slow down as my number of non-resolved tasks increases, I'll
look into indexing and other optimisations such as archiving really old tasks.
I don't believe that this will be necessary, as the number of open tasks is
(hopefully) bounded.

# Issues

As you've probably noticed, I don't use the github issues. Currently I use
dstask itself to track dstask bugs in my personal dstask repository. I've left
the issues system enabled to allow people to report bugs or request features.
As soon as dstask is used by more than a handful of people, I'll probably
import the dstask issues to github.


# General tips

* Overwhelmed by tasks? Try focussing by prioritising (set priorities) or narrowing the context. The `show-tags` and `show-projects` commands are useful for creating a context.
* Use dstask to track things you might forget, rather than everything. SNR is important.
* Spend regular time reviewing tasks. You'll probably find some you've already resolved, and many you've forgotten.
* Try to work through tasks from the top of the list. Dstask sorts by priority then creation date -- the most important tasks are at the top.
