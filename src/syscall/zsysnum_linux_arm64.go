// mksysnum_linux.pl /usr/include/asm-generic/unistd.h
// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT

package syscall

const (
	SYS_IO_SETUP               = 0
	SYS_IO_DESTROY             = 1
	SYS_IO_SUBMIT              = 2
	SYS_IO_CANCEL              = 3
	SYS_IO_GETEVENTS           = 4
	SYS_SETXATTR               = 5
	SYS_LSETXATTR              = 6
	SYS_FSETXATTR              = 7
	SYS_GETXATTR               = 8
	SYS_LGETXATTR              = 9
	SYS_FGETXATTR              = 10
	SYS_LISTXATTR              = 11
	SYS_LLISTXATTR             = 12
	SYS_FLISTXATTR             = 13
	SYS_REMOVEXATTR            = 14
	SYS_LREMOVEXATTR           = 15
	SYS_FREMOVEXATTR           = 16
	SYS_GETCWD                 = 17
	SYS_LOOKUP_DCOOKIE         = 18
	SYS_EVENTFD2               = 19
	SYS_EPOLL_CREATE1          = 20
	SYS_EPOLL_CTL              = 21
	SYS_EPOLL_PWAIT            = 22
	SYS_DUP                    = 23
	SYS_DUP3                   = 24
	SYS_FCNTL                  = 25
	SYS_INOTIFY_INIT1          = 26
	SYS_INOTIFY_ADD_WATCH      = 27
	SYS_INOTIFY_RM_WATCH       = 28
	SYS_IOCTL                  = 29
	SYS_IOPRIO_SET             = 30
	SYS_IOPRIO_GET             = 31
	SYS_FLOCK                  = 32
	SYS_MKNODAT                = 33
	SYS_MKDIRAT                = 34
	SYS_UNLINKAT               = 35
	SYS_SYMLINKAT              = 36
	SYS_LINKAT                 = 37
	SYS_RENAMEAT               = 38
	SYS_UMOUNT2                = 39
	SYS_MOUNT                  = 40
	SYS_PIVOT_ROOT             = 41
	SYS_NFSSERVCTL             = 42
	SYS_STATFS                 = 43
	SYS_FSTATFS                = 44
	SYS_TRUNCATE               = 45
	SYS_FTRUNCATE              = 46
	SYS_FALLOCATE              = 47
	SYS_FACCESSAT              = 48
	SYS_CHDIR                  = 49
	SYS_FCHDIR                 = 50
	SYS_CHROOT                 = 51
	SYS_FCHMOD                 = 52
	SYS_FCHMODAT               = 53
	SYS_FCHOWNAT               = 54
	SYS_FCHOWN                 = 55
	SYS_OPENAT                 = 56
	SYS_CLOSE                  = 57
	SYS_VHANGUP                = 58
	SYS_PIPE2                  = 59
	SYS_QUOTACTL               = 60
	SYS_GETDENTS64             = 61
	SYS_LSEEK                  = 62
	SYS_READ                   = 63
	SYS_WRITE                  = 64
	SYS_READV                  = 65
	SYS_WRITEV                 = 66
	SYS_PREAD64                = 67
	SYS_PWRITE64               = 68
	SYS_PREADV                 = 69
	SYS_PWRITEV                = 70
	SYS_SENDFILE               = 71
	SYS_PSELECT6               = 72
	SYS_PPOLL                  = 73
	SYS_SIGNALFD4              = 74
	SYS_VMSPLICE               = 75
	SYS_SPLICE                 = 76
	SYS_TEE                    = 77
	SYS_READLINKAT             = 78
	SYS_FSTATAT                = 79
	SYS_FSTAT                  = 80
	SYS_SYNC                   = 81
	SYS_FSYNC                  = 82
	SYS_FDATASYNC              = 83
	SYS_SYNC_FILE_RANGE2       = 84
	SYS_SYNC_FILE_RANGE        = 84
	SYS_TIMERFD_CREATE         = 85
	SYS_TIMERFD_SETTIME        = 86
	SYS_TIMERFD_GETTIME        = 87
	SYS_UTIMENSAT              = 88
	SYS_ACCT                   = 89
	SYS_CAPGET                 = 90
	SYS_CAPSET                 = 91
	SYS_PERSONALITY            = 92
	SYS_EXIT                   = 93
	SYS_EXIT_GROUP             = 94
	SYS_WAITID                 = 95
	SYS_SET_TID_ADDRESS        = 96
	SYS_UNSHARE                = 97
	SYS_FUTEX                  = 98
	SYS_SET_ROBUST_LIST        = 99
	SYS_GET_ROBUST_LIST        = 100
	SYS_NANOSLEEP              = 101
	SYS_GETITIMER              = 102
	SYS_SETITIMER              = 103
	SYS_KEXEC_LOAD             = 104
	SYS_INIT_MODULE            = 105
	SYS_DELETE_MODULE          = 106
	SYS_TIMER_CREATE           = 107
	SYS_TIMER_GETTIME          = 108
	SYS_TIMER_GETOVERRUN       = 109
	SYS_TIMER_SETTIME          = 110
	SYS_TIMER_DELETE           = 111
	SYS_CLOCK_SETTIME          = 112
	SYS_CLOCK_GETTIME          = 113
	SYS_CLOCK_GETRES           = 114
	SYS_CLOCK_NANOSLEEP        = 115
	SYS_SYSLOG                 = 116
	SYS_PTRACE                 = 117
	SYS_SCHED_SETPARAM         = 118
	SYS_SCHED_SETSCHEDULER     = 119
	SYS_SCHED_GETSCHEDULER     = 120
	SYS_SCHED_GETPARAM         = 121
	SYS_SCHED_SETAFFINITY      = 122
	SYS_SCHED_GETAFFINITY      = 123
	SYS_SCHED_YIELD            = 124
	SYS_SCHED_GET_PRIORITY_MAX = 125
	SYS_SCHED_GET_PRIORITY_MIN = 126
	SYS_SCHED_RR_GET_INTERVAL  = 127
	SYS_RESTART_SYSCALL        = 128
	SYS_KILL                   = 129
	SYS_TKILL                  = 130
	SYS_TGKILL                 = 131
	SYS_SIGALTSTACK            = 132
	SYS_RT_SIGSUSPEND          = 133
	SYS_RT_SIGACTION           = 134
	SYS_RT_SIGPROCMASK         = 135
	SYS_RT_SIGPENDING          = 136
	SYS_RT_SIGTIMEDWAIT        = 137
	SYS_RT_SIGQUEUEINFO        = 138
	SYS_RT_SIGRETURN           = 139
	SYS_SETPRIORITY            = 140
	SYS_GETPRIORITY            = 141
	SYS_REBOOT                 = 142
	SYS_SETREGID               = 143
	SYS_SETGID                 = 144
	SYS_SETREUID               = 145
	SYS_SETUID                 = 146
	SYS_SETRESUID              = 147
	SYS_GETRESUID              = 148
	SYS_SETRESGID              = 149
	SYS_GETRESGID              = 150
	SYS_SETFSUID               = 151
	SYS_SETFSGID               = 152
	SYS_TIMES                  = 153
	SYS_SETPGID                = 154
	SYS_GETPGID                = 155
	SYS_GETSID                 = 156
	SYS_SETSID                 = 157
	SYS_GETGROUPS              = 158
	SYS_SETGROUPS              = 159
	SYS_UNAME                  = 160
	SYS_SETHOSTNAME            = 161
	SYS_SETDOMAINNAME          = 162
	SYS_GETRLIMIT              = 163
	SYS_SETRLIMIT              = 164
	SYS_GETRUSAGE              = 165
	SYS_UMASK                  = 166
	SYS_PRCTL                  = 167
	SYS_GETCPU                 = 168
	SYS_GETTIMEOFDAY           = 169
	SYS_SETTIMEOFDAY           = 170
	SYS_ADJTIMEX               = 171
	SYS_GETPID                 = 172
	SYS_GETPPID                = 173
	SYS_GETUID                 = 174
	SYS_GETEUID                = 175
	SYS_GETGID                 = 176
	SYS_GETEGID                = 177
	SYS_GETTID                 = 178
	SYS_SYSINFO                = 179
	SYS_MQ_OPEN                = 180
	SYS_MQ_UNLINK              = 181
	SYS_MQ_TIMEDSEND           = 182
	SYS_MQ_TIMEDRECEIVE        = 183
	SYS_MQ_NOTIFY              = 184
	SYS_MQ_GETSETATTR          = 185
	SYS_MSGGET                 = 186
	SYS_MSGCTL                 = 187
	SYS_MSGRCV                 = 188
	SYS_MSGSND                 = 189
	SYS_SEMGET                 = 190
	SYS_SEMCTL                 = 191
	SYS_SEMTIMEDOP             = 192
	SYS_SEMOP                  = 193
	SYS_SHMGET                 = 194
	SYS_SHMCTL                 = 195
	SYS_SHMAT                  = 196
	SYS_SHMDT                  = 197
	SYS_SOCKET                 = 198
	SYS_SOCKETPAIR             = 199
	SYS_BIND                   = 200
	SYS_LISTEN                 = 201
	SYS_ACCEPT                 = 202
	SYS_CONNECT                = 203
	SYS_GETSOCKNAME            = 204
	SYS_GETPEERNAME            = 205
	SYS_SENDTO                 = 206
	SYS_RECVFROM               = 207
	SYS_SETSOCKOPT             = 208
	SYS_GETSOCKOPT             = 209
	SYS_SHUTDOWN               = 210
	SYS_SENDMSG                = 211
	SYS_RECVMSG                = 212
	SYS_READAHEAD              = 213
	SYS_BRK                    = 214
	SYS_MUNMAP                 = 215
	SYS_MREMAP                 = 216
	SYS_ADD_KEY                = 217
	SYS_REQUEST_KEY            = 218
	SYS_KEYCTL                 = 219
	SYS_CLONE                  = 220
	SYS_EXECVE                 = 221
	SYS_MMAP                   = 222
	SYS_FADVISE64              = 223
	SYS_SWAPON                 = 224
	SYS_SWAPOFF                = 225
	SYS_MPROTECT               = 226
	SYS_MSYNC                  = 227
	SYS_MLOCK                  = 228
	SYS_MUNLOCK                = 229
	SYS_MLOCKALL               = 230
	SYS_MUNLOCKALL             = 231
	SYS_MINCORE                = 232
	SYS_MADVISE                = 233
	SYS_REMAP_FILE_PAGES       = 234
	SYS_MBIND                  = 235
	SYS_GET_MEMPOLICY          = 236
	SYS_SET_MEMPOLICY          = 237
	SYS_MIGRATE_PAGES          = 238
	SYS_MOVE_PAGES             = 239
	SYS_RT_TGSIGQUEUEINFO      = 240
	SYS_PERF_EVENT_OPEN        = 241
	SYS_ACCEPT4                = 242
	SYS_RECVMMSG               = 243
	SYS_ARCH_SPECIFIC_SYSCALL  = 244
	SYS_WAIT4                  = 260
	SYS_PRLIMIT64              = 261
	SYS_FANOTIFY_INIT          = 262
	SYS_FANOTIFY_MARK          = 263
	SYS_NAME_TO_HANDLE_AT      = 264
	SYS_OPEN_BY_HANDLE_AT      = 265
	SYS_CLOCK_ADJTIME          = 266
	SYS_SYNCFS                 = 267
	SYS_SETNS                  = 268
	SYS_SENDMMSG               = 269
	SYS_PROCESS_VM_READV       = 270
	SYS_PROCESS_VM_WRITEV      = 271
	SYS_KCMP                   = 272
	SYS_FINIT_MODULE           = 273
	SYS_SCHED_SETATTR          = 274
	SYS_SCHED_GETATTR          = 275
	SYS_RENAMEAT2              = 276
	SYS_SECCOMP                = 277
	SYS_GETRANDOM              = 278
	SYS_MEMFD_CREATE           = 279
	SYS_BPF                    = 280
	SYS_EXECVEAT               = 281

	SYS_EPOLL_CREATE = 1042
	SYS_EPOLL_WAIT = 1069
)
