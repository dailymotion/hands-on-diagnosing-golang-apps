# Standard Unix Tools

There are some really useful unix tools, that are most of the time the first ones we use when diagnosing an application:
- `ps`
- `top`
- `lsof`
- ...

## ps

The `ps` command - which stands for `Process Status` - displays the currently-running processes. Everybody knows `ps`... or a least knows a little about it, because it's a complex tool, with different behaviour on different OS.

If we run the `ps` command with the "classic" `aux` options, we should see the list of all running processes:

```
$ ps aux
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.0  0.0  38080  5392 ?        Ss   Apr05   0:40 /sbin/init
root         2  0.0  0.0      0     0 ?        S    Apr05   0:00 [kthreadd]
root         3  0.0  0.0      0     0 ?        S    Apr05   1:03 [ksoftirqd/0]
root         5  0.0  0.0      0     0 ?        S<   Apr05   0:00 [kworker/0:0H]
root         7  0.0  0.0      0     0 ?        S    Apr05  20:10 [rcu_sched]
root         8  0.0  0.0      0     0 ?        S    Apr05   0:00 [rcu_bh]
root         9  0.0  0.0      0     0 ?        S    Apr05   0:22 [migration/0]
root        10  0.0  0.0      0     0 ?        S    Apr05   0:28 [watchdog/0]
...
```

If you want to know how a process is behaving, the most interesting columns are the `%CPU` and `%MEM` - but this is "relative" information, expressed in percentage. If you want to know how much memory a process is using (in bytes), you should really look at the `RSS` column, which is the `Resident Set Size` - basically, it represents the amount of memory that is really in RAM, and had not been swaped-out (or paged-out).

If you want to get that information for a specific process, you can use the `ps l PID` command. The `l` option has the advantage of displaying roughly the same informations on different systems:

- on Linux:

  ```
  $ ps l 23929
  F   UID   PID  PPID PRI  NI    VSZ    RSS WCHAN  STAT TTY        TIME COMMAND
  0  3140 23929     1  20   0 859768 526472 futex_ Sl  ?        193:46 /opt/deploy/builds/console-api/console-api
  ```

- on MacOS:

  ```
  $ ps l 86113
  UID   PID  PPID CPU PRI NI       VSZ    RSS WCHAN  STAT   TT       TIME COMMAND
  502 86113 76310   0  31  0 556619472   4316 -      S+   s005    0:00.01 ./hands-on-diagnosing-golang-apps
  ```

Of course there are a lot more things to know about `ps` - this is just basic knowledge to get you started, but remember that `man` is your friend ;-)

## top

Another very useful command to get a quick look at what processes are doing, is the `top` command.

People are used to just type `top`, and see all processes:

```
$ top
Processes: 422 total, 2 running, 420 sleeping, 2139 threads
Load Avg: 1.59, 1.95, 2.11  CPU usage: 1.88% user, 2.12% sys, 95.99% idle  SharedLibs: 138M resident, 34M data, 38M linkedit.
MemRegions: 179803 total, 5853M resident, 105M private, 1738M shared. PhysMem: 15G used (2652M wired), 1264M unused.
VM: 3804G vsize, 627M framework vsize, 192472713(0) swapins, 196455544(0) swapouts. Networks: packets: 217947936/131G in, 181904723/144G out.
Disks: 54149815/1676G read, 40676473/1812G written.

PID    COMMAND      %CPU TIME     #TH   #WQ  #PORT MEM    PURG   CMPRS  PGRP  PPID  STATE    BOOSTS            %CPU_ME %CPU_OTHRS UID  FAULTS     COW      MSGSENT
98825  applessdstat 0.0  00:00.01 3     2    33    8192B  0B     740K   98825 1     sleeping *0[1]             0.00000 0.00000    0    758        123      90
97556  zsh          0.0  00:03.59 1     0    16    680K   0B     5756K  97556 97555 sleeping *0[1]             0.00000 0.00000    502  87063      19641    1000
97555  login        0.0  00:00.06 2     1    29    8192B  0B     1004K  97555 97554 sleeping *0[9]             0.00000 0.00000    0    939        159      156
97554  iTerm2       0.0  00:01.00 2     1    29    220K   0B     3608K  97554 282   sleeping *0[1]             0.00000 0.00000    502  24191      649      80
96908  Tunnelblick  0.0  03:37.06 5     1    230   13M    0B     14M    96908 1     sleeping *0[13447]         0.00000 0.00000    502  419718     964      1432433
94861  top          5.4  00:00.96 1/1   0    20    4952K  0B     0B     94861 86116 running  *0[1]             0.00000 0.00000    0    5473+      115      480561+
94672  Google Chrom 0.0  00:00.11 15    1    111   14M    0B     0B     44599 44599 sleeping *0[1]             0.00000 0.00000    502  7985       1682     329
...
```

But you can also call it directly with a PID:

- on Linux:

  ```
  $ top -p 23929
  top - 13:04:56 up 110 days,  2:16,  3 users,  load average: 0.05, 0.08, 0.06
  Tasks:   1 total,   0 running,   1 sleeping,   0 stopped,   0 zombie
  %Cpu(s):  0.2 us,  0.2 sy,  0.0 ni, 99.7 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
  KiB Mem :  8173488 total,   811308 free,  1148280 used,  6213900 buff/cache
  KiB Swap:        0 total,        0 free,        0 used.  6458336 avail Mem

    PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND
  23929 deploy    20   0  859768 717732  20956 S   0.3  8.8 201:26.15 console-api
  ```

- on MacOS:

  ```
  $ top -pid 86113
  Processes: 422 total, 2 running, 420 sleeping, 2148 threads
  Load Avg: 2.03, 1.89, 1.96  CPU usage: 4.48% user, 3.53% sys, 91.98% idle  SharedLibs: 140M resident, 34M data, 39M linkedit.
  MemRegions: 178117 total, 5706M resident, 106M private, 1818M shared. PhysMem: 14G used (2656M wired), 1671M unused.
  VM: 3801G vsize, 627M framework vsize, 192473161(0) swapins, 196455544(0) swapouts. Networks: packets: 217989050/131G in, 181917220/144G out.
  Disks: 54154122/1676G read, 40687445/1812G written.

  PID    COMMAND      %CPU TIME     #TH  #WQ  #POR MEM    PURG CMPR PGRP  PPID  STATE    BOOSTS     %CPU_ME %CPU_OTHRS UID  FAULT COW  MSGS MSGR SYSB SYSM CSW  PAGE IDLE
  86113  hands-on-dia 0.0  00:00.01 8    0    33   1516K  0B   0B   86113 76310 sleeping *0[1]      0.00000 0.00000    502  1289  126  44   21   425  305  305  0    52
  ```

Notice that you need to use the `-p PID` option on Linux, and `-pid PID` on MacOS...

## lsof

Yet another very useful command is `lsof`, which lists the open file descriptors. Same as `ps` and `top`, it can display information for all processes, or only a specific one - with the `-p PID` option.

For example, on Linux:

```
$ lsof -p 23929
COMMAND     PID   USER   FD      TYPE   DEVICE SIZE/OFF     NODE NAME
console-a 23929 deploy  cwd       DIR    202,1     4096  2901852 /opt/deploy/builds/console-api_vmaster_1532338773
console-a 23929 deploy  rtd       DIR    202,1     4096        2 /
console-a 23929 deploy  txt       REG    202,1 44220713  2901860 /opt/deploy/builds/console-api_vmaster_1532338773/console-api
console-a 23929 deploy    0r     FIFO     0,10      0t0 81927771 pipe
console-a 23929 deploy    1w      REG    202,1  4522504   137397 /var/log/console-api/start-stop-daemon.log
console-a 23929 deploy    2w      REG    202,1  4522504   137397 /var/log/console-api/start-stop-daemon.log
console-a 23929 deploy    3u     IPv4 81928392      0t0      UDP localhost:38524->localhost:8125
console-a 23929 deploy    4u  a_inode     0,11        0     8124 [eventpoll]
console-a 23929 deploy    5u     IPv6 81928396      0t0      TCP *:8888 (LISTEN)
console-a 23929 deploy    6u     IPv6 82720252      0t0      TCP ip-10-0-58-200.us-west-2.compute.internal:http-alt->ip-10-0-69-16.us-west-2.compute.internal:49258 (ESTABLISHED)
console-a 23929 deploy    7u     IPv4 81928092      0t0      TCP ip-10-0-58-200.us-west-2.compute.internal:33096->ip-10-0-75-170.us-west-2.compute.internal:postgresql (ESTABLISHED)
console-a 23929 deploy    9u     IPv6 82720281      0t0      TCP ip-10-0-58-200.us-west-2.compute.internal:http-alt->ip-10-128-11-16.us-west-2.compute.internal:55639 (ESTABLISHED)
console-a 23929 deploy   10u     IPv6 82719333      0t0      TCP ip-10-0-58-200.us-west-2.compute.internal:http-alt->ip-10-0-58-240.us-west-2.compute.internal:41440 (ESTABLISHED)
console-a 23929 deploy   13u     IPv4 81928425      0t0      TCP ip-10-0-58-200.us-west-2.compute.internal:48902->ip-10-0-93-10.us-west-2.compute.internal:9092 (ESTABLISHED)
console-a 23929 deploy   14u     IPv4 81928101      0t0      TCP ip-10-0-58-200.us-west-2.compute.internal:35594->ip-10-0-93-11.us-west-2.compute.internal:9092 (ESTABLISHED)
console-a 23929 deploy   15u     IPv4 81928428      0t0      TCP ip-10-0-58-200.us-west-2.compute.internal:45802->ip-10-0-69-10.us-west-2.compute.internal:9092 (ESTABLISHED)
console-a 23929 deploy   16u     IPv4 81928105      0t0      TCP ip-10-0-58-200.us-west-2.compute.internal:48908->ip-10-0-93-10.us-west-2.compute.internal:9092 (ESTABLISHED)
console-a 23929 deploy   17u     IPv6 81928429      0t0      TCP *:http-alt (LISTEN)
console-a 23929 deploy   18u     IPv4 82704126      0t0      TCP ip-10-0-58-200.us-west-2.compute.internal:39148->ip-10-0-75-170.us-west-2.compute.internal:postgresql (ESTABLISHED)
console-a 23929 deploy   20u     IPv4 82705028      0t0      TCP ip-10-0-58-200.us-west-2.compute.internal:39295->ip-10-0-75-170.us-west-2.compute.internal:postgresql (ESTABLISHED)
console-a 23929 deploy   22u     IPv4 82705561      0t0      TCP ip-10-0-58-200.us-west-2.compute.internal:39294->ip-10-0-75-170.us-west-2.compute.internal:postgresql (ESTABLISHED)
console-a 23929 deploy   26u     IPv4 82705030      0t0      TCP ip-10-0-58-200.us-west-2.compute.internal:39302->ip-10-0-75-170.us-west-2.compute.internal:postgresql (ESTABLISHED)
```

There are a lot of information here:

- the `current working directory` (`cwd`) of the application is `/opt/deploy/builds/console-api_vmaster_1532338773`
- the `program` (`txt`, for `program text`) is `/opt/deploy/builds/console-api_vmaster_1532338773/console-api`
- the `stdin` (FD `0`) is a pipe
- the `stdout` (FD `1`) is `/var/log/console-api/start-stop-daemon.log`
- the `stderr` (FD `2`) is `/var/log/console-api/start-stop-daemon.log` too

So if you want to know where you can find the output of an application, or from where it was started, `lsof` is your friend.

You can also see on which ports the application is listening: in our example, on ports `8888` and `http-alt` - which is `8080`, as returned by `grep http-alt /etc/services`. And we can see the incoming/outgoing TCP connections, for example here we see that:
- there are 5 connections to postgres
- there are 4 connections to 3 different kafka brokers
- there are 3 connections from external clients to the `http-alt` port

## Proc files on Linux

On Linux, you can learn a great deal of information about a process, from looking at the virtual `/proc` filesystem. If you know the PID of the process, you can list the files under `/proc/PID`:

```
$ ls -lh /proc/23929
total 0
dr-xr-xr-x  2 deploy deploy 0 Jul 24 14:20 attr
-rw-r--r--  1 deploy deploy 0 Jul 24 14:20 autogroup
-r--------  1 deploy deploy 0 Jul 24 14:20 auxv
-r--r--r--  1 deploy deploy 0 Jul 24 14:20 cgroup
--w-------  1 deploy deploy 0 Jul 24 14:20 clear_refs
-r--r--r--  1 deploy deploy 0 Jul 23 09:39 cmdline
-rw-r--r--  1 deploy deploy 0 Jul 24 14:20 comm
-rw-r--r--  1 deploy deploy 0 Jul 24 14:20 coredump_filter
-r--r--r--  1 deploy deploy 0 Jul 23 09:39 cpuset
lrwxrwxrwx  1 deploy deploy 0 Jul 24 13:08 cwd -> /opt/deploy/builds/console-api_vmaster_1532338773
-r--------  1 deploy deploy 0 Jul 24 14:20 environ
lrwxrwxrwx  1 deploy deploy 0 Jul 23 09:39 exe -> /opt/deploy/builds/console-api_vmaster_1532338773/console-api
dr-x------  2 deploy deploy 0 Jul 23 09:39 fd
dr-x------  2 deploy deploy 0 Jul 24 13:08 fdinfo
-rw-r--r--  1 deploy deploy 0 Jul 24 14:20 gid_map
-r--------  1 deploy deploy 0 Jul 23 09:39 io
-r--r--r--  1 deploy deploy 0 Jul 24 14:20 limits
-rw-r--r--  1 deploy deploy 0 Jul 24 14:20 loginuid
dr-x------  2 deploy deploy 0 Jul 24 14:20 map_files
-r--r--r--  1 deploy deploy 0 Jul 24 13:08 maps
-rw-------  1 deploy deploy 0 Jul 24 14:20 mem
-r--r--r--  1 deploy deploy 0 Jul 24 14:20 mountinfo
-r--r--r--  1 deploy deploy 0 Jul 24 14:20 mounts
-r--------  1 deploy deploy 0 Jul 24 14:20 mountstats
dr-xr-xr-x  5 deploy deploy 0 Jul 24 14:20 net
dr-x--x--x  2 deploy deploy 0 Jul 24 14:20 ns
-r--r--r--  1 deploy deploy 0 Jul 24 14:20 numa_maps
-rw-r--r--  1 deploy deploy 0 Jul 24 14:20 oom_adj
-r--r--r--  1 deploy deploy 0 Jul 24 14:20 oom_score
-rw-r--r--  1 deploy deploy 0 Jul 24 14:20 oom_score_adj
-r--------  1 deploy deploy 0 Jul 24 14:20 pagemap
-r--------  1 deploy deploy 0 Jul 24 14:20 personality
-rw-r--r--  1 deploy deploy 0 Jul 24 14:20 projid_map
lrwxrwxrwx  1 deploy deploy 0 Jul 24 13:08 root -> /
-rw-r--r--  1 deploy deploy 0 Jul 24 14:20 sched
-r--r--r--  1 deploy deploy 0 Jul 24 14:20 schedstat
-r--r--r--  1 deploy deploy 0 Jul 24 14:20 sessionid
-rw-r--r--  1 deploy deploy 0 Jul 24 14:20 setgroups
-r--r--r--  1 deploy deploy 0 Jul 24 14:20 smaps
-r--------  1 deploy deploy 0 Jul 24 14:20 stack
-r--r--r--  1 deploy deploy 0 Jul 23 09:39 stat
-r--r--r--  1 deploy deploy 0 Jul 23 09:39 statm
-r--r--r--  1 deploy deploy 0 Jul 23 09:39 status
-r--------  1 deploy deploy 0 Jul 24 14:20 syscall
dr-xr-xr-x 13 deploy deploy 0 Jul 24 14:20 task
-r--r--r--  1 deploy deploy 0 Jul 24 14:20 timers
-rw-r--r--  1 deploy deploy 0 Jul 24 14:20 uid_map
-r--r--r--  1 deploy deploy 0 Jul 24 11:34 wchan
```

- you can see the `cwd` and `exe`cutable, same as with `lsof`
- in the `fd` directory you can see the file descriptors, with `stdin` (`0`), `stdout` (`1`) and `stderr` (`2`):

  ```
  $ ls -lh /proc/23929/fd
  total 0
  lr-x------ 1 deploy deploy 64 Jul 24 13:08 0 -> pipe:[81927771]
  l-wx------ 1 deploy deploy 64 Jul 24 13:08 1 -> /var/log/console-api/start-stop-daemon.log
  lrwx------ 1 deploy deploy 64 Jul 24 13:08 13 -> socket:[81928425]
  lrwx------ 1 deploy deploy 64 Jul 24 13:08 14 -> socket:[81928101]
  lrwx------ 1 deploy deploy 64 Jul 24 13:08 15 -> socket:[81928428]
  lrwx------ 1 deploy deploy 64 Jul 24 13:08 16 -> socket:[81928105]
  lrwx------ 1 deploy deploy 64 Jul 24 13:08 17 -> socket:[81928429]
  lrwx------ 1 deploy deploy 64 Jul 24 13:15 19 -> socket:[82737378]
  l-wx------ 1 deploy deploy 64 Jul 24 13:08 2 -> /var/log/console-api/start-stop-daemon.log
  lrwx------ 1 deploy deploy 64 Jul 24 13:42 23 -> socket:[82736894]
  lrwx------ 1 deploy deploy 64 Jul 24 13:13 24 -> socket:[82737380]
  lrwx------ 1 deploy deploy 64 Jul 24 13:13 25 -> socket:[82736678]
  lrwx------ 1 deploy deploy 64 Jul 24 13:08 3 -> socket:[81928392]
  lrwx------ 1 deploy deploy 64 Jul 24 13:08 4 -> anon_inode:[eventpoll]
  lrwx------ 1 deploy deploy 64 Jul 24 13:08 5 -> socket:[81928396]
  lrwx------ 1 deploy deploy 64 Jul 24 13:08 7 -> socket:[81928092]
  ```
- in `/proc/23929/cmdline` you have... the command line used to start the process
- in `/proc/23929/environ` you can see the environment variables of the process
- in `/proc/23929/limits` you have all the limits that apply to the process (like file descriptors limit, ...)
- and so on...

## Next

You can now head over to the next section, on [sending signals to a running application](../signals/README.md).
