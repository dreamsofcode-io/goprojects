# Todo App

## Goal

Create an cli application for managing tasks in the terminal.

```
$ tasks
```

## Requirements

Should be able to perform crud operations via a cli on a data file of tasks. The operations should be as follows:

```
$ tasks add "My new task"
$ tasks list
$ tasks complete 
```

### Add

The add method should be used to create new tasks in the underlying data store. It should take a positional argument with the task description

```
$ tasks add <description>
```

for example:

```
$ tasks add "Tidy my desk"
```

should add a new task with the description of "Tidy my desk"

### List

This method should return a list of all of the **uncompleted** tasks, with the option to return all tasks regardless of whether or not they are completed.

for example:

```
$ tasks list
ID    Task                                                Created
1     Tidy up my desk                                     a minute ago
3     Change my keyboard mapping to use escape/control    a few seconds ago
```

or for showing all tasks, using a flag (such as -a or --all)

```
$ tasks list -a
ID    Task                                                Created          Done
1     Tidy up my desk                                     2 minutes ago    false
2     Write up documentation for new project feature      a minute ago     true
3     Change my keyboard mapping to use escape/control    a minute ago     false
```


### Complete

To mark a task as done, add in the following method

```
$ tasks complete <taskid>
```

### Delete

The following method should be implemented to delete a task from the data store

```
$ tasks delete <taskid>
```

## Notable Packages Used

- `encoding/csv` for writing out as a csv file
- `strconv` for turning types into strings and visa versa
- `text/tabwriter` for writing out tab aligned output
- `os` for opening and reading files
- `github.com/spf13/cobra` for the command line interface
- `github.com/mergestat/timediff` for displaying relative friendly time differences (1 hour ago, 10 minutes ago, etc)
 
## Custom Resources

### Example Application

You can find an [example version](https://github.com/dreamsofcode-io/goprojects/releases/tag/0.1.0) of this todo list on the releases tab of this repo.

### Example Data File

Additionally, an example CSV looks like as follows:

```
ID,Description,CreatedAt,IsComplete
1,My new task,2024-07-27T16:45:19-05:00,true
2,Finish this video,2024-07-27T16:45:26-05:00,true
3,Find a video editor,2024-07-27T16:45:31-05:00,false
```

## Technical Considerations

### Stderr vs Stdout

Make sure to write any diagnostics or errors to stderr stream and write output to stdout.

### File Locking
One major consideration is that the underlying data file should be locked by the process to prevent concurrent read/writes. This can
be achieved using the flock system call in unix like systems to obtain an exclusive lock on the file.

You can achieve this in go using the following code:

```go
func loadFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

    // Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}
```

Then to unlock the file, use the following:

```go
func closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}
```

## Extra Features

- Change the IsComplete property of the Task data model to use a timestamp instead, which gives further information.
- Change from CSV to JSON, JSONL or SQLite
- Add in an optional due date to the tasks
