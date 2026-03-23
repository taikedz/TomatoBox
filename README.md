# Fyne Pomodoro

Unsatisfied with existing pomodoro apps, I'd like to write my own.

Create a pomodoro driven by a sequence table. By default the table is hidden, and is instead configured with "Focus", "Break" and "Long break" mode settings. Each of these have a duration, and a lead-up time-bell to alert user to perform wrap-up actions before the next mode starts.

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
+--------------------------------------------------------+
Run-type = repeat
```

It cycles through the table, each time displaying the title.

**Custom table**

A table creation mode can be entered allowing the user to add their own entries.

* they can define "Activity types" in one space which give them default values
* and then add "Activities" to the table, each with a custom title
* run mode then can be "single" or "repeat"
* each activity can ask to confirm start (default false - just starts)

This can allow something like this to be quickly built

```
+========================================================+
| Title           | Wrap-up  |   Duration |  Auto-start  |
+========================================================+
| Do flash cards  |    00:30 |   25:00    | False        |
+--------------------------------------------------------+
| Review notes    |    00:05 |   05:00    | False        |
+--------------------------------------------------------+
| Break           |    00:30 |   05:00    | True         |
+--------------------------------------------------------+
| Tidy desk       |    00:10 |   05:00    | False        |
+--------------------------------------------------------+
| Research travel |    00:30 |   25:00    | False        |
+--------------------------------------------------------+
| Have lunch      |    05:00 |   45:00    | True         |
+--------------------------------------------------------+
Run-type = single
```

