package dstask

import (
	"fmt"
	"os"
)

func Help(cmd string) {
	var helpStr string

	switch cmd {
	case CMD_NEXT:
		helpStr = `Usage: dstask next [filter] [--]
Usage: dstask [filter] [--]
Example: dstask +work +bug --

Display list of non-resolved tasks in the current context, most recent last,
optional filter. It is the default command, so "next" is unnecessary.

Bypass the current context with --.
`
	case CMD_ADD:
		helpStr = `Usage: dstask add [task summary] [--]
Example: dstask add Fix main web page 500 error +bug P1 project:website

Add a task, returning the git commit output which contains the task ID, used
later to reference the task.

Tags, project and priority can be added anywhere within the task summary.

Add -- to ignore the current context. / can be used when adding tasks to note
any words after.

`

	case CMD_LOG:
		helpStr = `Usage: dstask log [task summary] [--]
Example: dstask log Fix main web page 500 error +bug P1 project:website

Add an immediately resolved task. Syntax identical to add command.

Tags, project and priority can be added anywhere within the task summary.

Add -- to ignore the current context.

`
	case CMD_START:
		helpStr = `Usage: dstask <id...> start
Usage: dstask start [task summary] [--]
Example: dstask 15 start
Example: dstask start Fix main web page 500 error +bug P1 project:website

Mark a task as active, meaning you're currently at work on the task.

Alternatively, "start" can add a task and start it immediately with the same
syntax is the "add" command.  Tags, project and priority can be added anywhere
within the task summary.

Add -- to ignore the current context.
`
	case CMD_NOTE:
		fallthrough
	case CMD_NOTES:
		helpStr = `Usage: dstask note <id>
Usage: dstask note <id> <text>
Example task 13 note problem is faulty hardware

Edit or append text to the markdown notes attached to a particular task.
`
	case CMD_STOP:
		helpStr = `Usage: dstask <id...> stop [text]
Example: dstask 15 stop
Example: dstask 15 stop replaced some hardware

Set a task as inactive, meaning you've stopped work on the task. Optional text
may be added, which will be appended to the note.
`
	case CMD_RESOLVE:
		fallthrough
	case CMD_DONE:
		helpStr = `Usage: dstask <id...> done [text]
Example: dstask 15 done
Example: dstask 15 done replaced some hardware

Resolve a task. Optional text may be added, which will be appended to the note.
`
	case CMD_CONTEXT:
		helpStr = `Usage: dstask context <filter>
Example: dstask context +work -bug

Set a global filter consisting of a project, tags or antitags. Subsequent new
tasks and most commands will then have this filter applied automatically.

For example, if you were to run "task add fix the webserver," the given task
would then have the tag "work" applied automatically.
`
	case CMD_MODIFY:
		helpStr = `Usage: dstask <id...> modify <filter>
Example: dstask 34 modify -work +home project:workbench -project:website

Modify the attributes of a task.
`
	case CMD_EDIT:
		helpStr = `Usage: dstask <id...> edit

Edit a task in your text editor.
`
	case CMD_UNDO:
		helpStr = `Usage: dstask undo

Undo the last command that changed the repository. This uses git revert on one
or more commits.
`
	case CMD_SYNC:
		helpStr = `Usage: dstask sync

Synchronise with the remote git server. Runs git pull then git push. If there
are conflicts that cannot be automatically resolved, it is necessary to
manually resolve them in  ~/.dstask or with the "task git" command.
`
	case CMD_GIT:
		helpStr = `Usage: dstask git <args...>
Example: dstask git status

Run the given git command inside ~/.dstask
`
	case CMD_SHOW_RESOLVED:
		helpStr = `Usage: dstask resolved

Show a report of last 1000 resolved tasks.
`
	case CMD_OPEN:
		helpStr = `Usage: dstask <id...> open

Open all URLs found within the task summary and notes. If you commonly have
dozens of tabs open to later action, convert them into tasks to open later with
this command.
`
	case CMD_SHOW_PROJECTS:
		helpStr = `Usage: dstask show-projects

Show a breakdown of projects with progress information
`
	case CMD_IMPORT_TW:
		helpStr = `Usage: cat export.json | task import-tw

Import tasks from a taskwarrior json dump. The "task export" taskwarrior
command can be used for this.
`
	default:
		helpStr = `Usage: dstask [id...] <cmd> [task summary/filter]

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

Task table key:

`
	}
	fmt.Fprintf(os.Stderr, helpStr)

	colourPrintln(0, FG_PRIORITY_CRITICAL, BG_DEFAULT_2, "Critical priority")
	colourPrintln(0, FG_PRIORITY_HIGH, BG_DEFAULT_2, "High priority")
	colourPrintln(0, FG_DEFAULT, BG_DEFAULT_1, "Normal priority")
	colourPrintln(0, FG_PRIORITY_LOW, BG_DEFAULT_2, "Low priority")
	colourPrintln(0, FG_ACTIVE, BG_ACTIVE, "Active")
	colourPrintln(0, FG_DEFAULT, BG_PAUSED, "Paused")

	os.Exit(0)
}

func colourPrintln(mode, fg, bg int, line string) {
	line = FixStr(line, 25)
	fmt.Fprintf(os.Stderr, "\033[%d;38;5;%d;48;5;%dm%s\033[0m\n", mode, fg, bg, line)
}
