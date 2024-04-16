# perfect28

perfect28 is a very simple command line benchmark that uses brute force to find the first 8 perfect numbers. This tool should do a decent job as a benchmark since the workload is linear and consitent.

## Benchmarks

Benchmarks are pre-configured workloads that aim not to be too quick nor too long.

- `--pn8` use all available CPU threads to reach perfect number 8. (see `About --pn8` before use)
- `--m24` Benchmark: Multi thread 2024.
- `--s24` Benchmark: Single thread 2024.

### Custom loop

You can always execute a custom loop if the available pre-configurations don't fit your needs.

- `--loop=<COUNT>` allows you to define the workload.
- `--repeat=<COUNT>` if COUNT=0, restarts the workload every time it completes until you kill it. If COUNT>1, restarts until COUNT is reached.
- `--thread=<COUNT>` let you define how many threads you wish to use. By default, perfect28 will use all CPU threads.

### About --pn8

The `--pn8` pre-configuration sets the workload to **2305843008139952128** which is the value of the 8th perfect number. While a 5950x find the first 7 perfect numbers in arround 2 minutes when single threaded, it would take more than 4 years for it to reach the perfect number 8 (when using all 32 CPU threads). If you decide to give it a try, be aware that it will take **a very long time** for the program to complete.

## Maintenance

I would be very happy to fix typos, add new valuable features and fix bugs if you find any.

## Special thanks

- [GoReleaser](https://github.com/goreleaser/goreleaser) for making the releases easier and faster.

