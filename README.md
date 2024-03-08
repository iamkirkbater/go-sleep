# Go Sleep

`gosleep` is a go replacement for the `sleep` command that allows you to put time-durations into your sleep commands.

The _main_ goal is to be able to use this as a somewhat backwards-compatible drop-in for `sleep`, via `alias sleep=gosleep`, while adding some usability features. The main system sleep would still be able to be used with `\sleep` if aliased.

## Examples

* `gosleep 1` - sleeps for 1 second
* `gosleep 1m` - sleeps for 1 minute
* `gosleep 1h30m` - sleeps for 1 hour and 30 minutes

## Benchmarks

Benchmark data is from my M1 Macbook Pro.

```bash
> time sleep 10
sleep 10  0.00s user 0.00s system 0% cpu 10.014 total

> time bin/gosleep 10
bin/gosleep 10  0.00s user 0.00s system 0% cpu 10.008 total
```

```bash
> time sleep $((2 * 60))
sleep $((2 * 60))  0.00s user 0.00s system 0% cpu 2:00.01 total

> time bin/gosleep 2m
bin/gosleep 2m  0.00s user 0.00s system 0% cpu 2:00.01 total
```

## Roadmap

* add a spinner or some type of progress notification
