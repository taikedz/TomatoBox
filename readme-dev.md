# Readme - devs

## Run the project

Structured so that there only need to run against one top file:

```sh
go run pomodoro.go
```

## Build distributable

Use the `./build.sh` script to produce a trimmed binary. It will typically depend on libc as Fyne has this particular dependency.
