# Profiling with pprof

* Example program: [https://github.com/miku/demo/profiling]

## Cpu profile

```
$ 
$  go tool pprof cpu.pprof
File: main
Type: cpu
Time: Jun 16, 2021 at 10:53pm (CEST)
Duration: 4.81s, Total samples = 4.96s (103.04%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 4.62s, 93.15% of 4.96s total
Dropped 54 nodes (cum <= 0.02s)
Showing top 10 nodes out of 38
      flat  flat%   sum%        cum   cum%
     2.02s 40.73% 40.73%      3.59s 72.38%  runtime.slicerunetostring
     1.48s 29.84% 70.56%      1.48s 29.84%  runtime.encoderune
     0.80s 16.13% 86.69%      4.68s 94.35%  main.main
     0.17s  3.43% 90.12%      0.17s  3.43%  runtime.memclrNoHeapPointers
     0.04s  0.81% 90.93%      0.14s  2.82%  runtime.sweepone
     0.03s   0.6% 91.53%      0.04s  0.81%  runtime.pageIndexOf (inline)
     0.02s   0.4% 91.94%      0.09s  1.81%  runtime.(*mspan).sweep
     0.02s   0.4% 92.34%      0.04s  0.81%  runtime.(*spanSet).push
     0.02s   0.4% 92.74%      0.06s  1.21%  runtime.greyobject
     0.02s   0.4% 93.15%      0.34s  6.85%  runtime.mallocgc
```

Web rendering.

```
$ go tool pprof -web cpu.pprof
```

Export to a file:

```
$ go tool pprof -png cpu.pprof > cpu.png
```

Other example:

* https://raw.githubusercontent.com/miku/span/master/docs/span-tagger.png


## Heap profile

