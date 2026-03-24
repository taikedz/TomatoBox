# TomatoBox: a customizable Pomodoro/Timeboxing utility

Unsatisfied with existing pomodoro apps, I'd like to write my own. In other apps I don't have quite the control I would like around setting activity duration, and how breaks/long breaks are distributed. One app in particular seems to have gotten stuck in setting long breaks only... without a clue why.

I also tend to want to know when time is coming to close on a task so I can wrap up to be ready to switch to the next task.

And also sometimes I just need to timebox a series of unrelated, time-varied activities...

## Activity Driven Design

Create a pomodoro driven by a sequence table, a list of activities. By default the table is hidden, and is instead configured with "Focus", "Break" and "Long break" mode settings. Each of these have a duration, and a lead-up time-bell to alert user to perform wrap-up actions before the next mode starts.

This creates a table like this in the background (number of "Focus" items before long-break is configurable):

```
+========================================================+
| Title           | Wrap-up  |   Duration |  Auto-start  |
+========================================================+
| Focus           |    00:30 |   25:00    | True         |
+--------------------------------------------------------+
| Break           |    00:10 |   05:00    | True         |
+--------------------------------------------------------+
| Focus           |    00:30 |   25:00    | True         |
+--------------------------------------------------------+
| Break           |    00:10 |   05:00    | True         |
+--------------------------------------------------------+
| Focus           |    00:30 |   25:00    | True         |
+--------------------------------------------------------+
| Break           |    00:10 |   05:00    | True         |
+--------------------------------------------------------+
| Long break      |    05:00 |   30:00    | True         |
+--------------------+-----------------------------------+
| Run-type = repeat  |
+--------------------+
```

It cycles through the table, each time displaying the title.

The table can be viewed by expanding a "table view". It is by default collapsed with default tables.

## Custom activity list

A table creation mode can be entered allowing the user to add their own entries.

* they can add "Activities" to the table, each with a custom title
* run mode then can be "single" or "repeat"
* each activity can can either autostart, or prompt the user to start next activity timer
* tables can be saved to files for later re-loading. this allows creating recurring routine plans

This can allow something like this to be quickly built

```
+========================================================+
| Title           | Wrap-up  |   Duration |  Auto-start  |
+========================================================+
| Do flash cards  |    00:30 |   25:00    | False        |  [+] [-] [^] [v]
+--------------------------------------------------------+
| Review notes    |    00:05 |   05:00    | False        |  [+] [-] [^] [v]
+--------------------------------------------------------+
| Break           |    00:30 |   05:00    | True         |  [+] [-] [^] [v]
+--------------------------------------------------------+
| Tidy desk       |    00:10 |   05:00    | False        |  [+] [-] [^] [v]
+--------------------------------------------------------+
| Research travel |    00:30 |   25:00    | False        |  [+] [-] [^] [v]
+--------------------------------------------------------+
| Have lunch      |    05:00 |   45:00    | True         |  [+] [-] [^] [v]
+--------------------------------------------------------+
|                 |          |            |              |  [+]
+--------------------+-----------------------------------+
| Run-type = single  |
+--------------------+

Type `Ctrl + Enter` to add a new task below, type `Shift + Enter` to add a new task above
Type `Ctrl + Bksp` to remove the line, type `Ctrl + Shift + Bksp` to clear the table
Type `Ctrl + S` to save the table to a new file, Type `Ctrl + O` to load a saved table
Type `Ctrl + Shift + C` to select a sequence to copy, then press `Ctrl + Enter` to add a copy below, press `Esc` to cancel.
Use `Alt + Up/Down` to move current item up or down in the list
```

In the run view, the normal table can be reinstated by choosing to "Clear custom table."
