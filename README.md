# Productivity Counter

I built this CLI stopwatch in GO for recording my productive work sessions and charting them like screen time graphs.

## Installation

To install the productivity-counter, clone the repository and build the project:

```sh
git clone https://github.com/GauravC4/productivity-counter.git
cd productivity-counter
go build -o productivity-counter
```

## Usage

Run the `productivity-counter` executable to start the REPL (Read-Eval-Print Loop) interface:
```./productivity-counter```

This will generate WAL.csv file which is the database of your sessions. Delete it to clear history.

## Commands

The following commands are available in the REPL:

- `start`: Start the clock.
- `analyse <n> <v>`: Displays analysis for the last n days with sessions if v is provided.
- `help`: Display a help message.
- `compress`: Compress the WAL file by removing redundant sessions.
- `exit`: Exit the productivity clock.

## Command Details

### `start`
Starts a new productive work session. The session will be recorded in the WAL (Write-Ahead Log) file. Press 'z' and Enter to stop the session.

### `analyse <n> <v>`
Displays analysis for the last n days. If v is provided, it will display verbose output with session details.
e.g. `analyse 2 v` will show analysis for today and yesterday.

Example output:
```
07:00:24AM [######..........................................] 1h30m16s / 12h
10:10:43AM [###.............................................] 50m22s / 12h
02:31:07PM [##############..................................] 3h30m8s / 12h
10:01:50PM [####............................................] 1h0m17s / 12h
2025-02-24 [###########################.....................] 6h51m4s / 12h

06:00:24AM [##########......................................] 2h30m16s / 12h
09:00:43AM [########........................................] 2h0m22s / 12h
02:01:07PM [################................................] 4h0m8s / 12h
09:01:50PM [########........................................] 2h0m17s / 12h
2025-02-25 [##########################################......] 10h31m4s / 12h
```

### `help`
Displays a help message with a list of available commands.

### `compress`
Compresses the WAL file by removing redundant sessions. This happens automatically after every session stop.

### `exit`
Exits the productivity clock.

## Future Improvements
- ability to cycle through commands in repl using arrow keys
- tag productivity session to classify work
- more analyse options like weekly and monthly
- compare particular days

## License

This project is licensed under the MIT License.
