# ELN_file_watcher
Version 0.1.1

As soon as all files in a subdirectory of &lt;CMD arg -src&gt; 
(or a file directly in &lt;CMD arg -src&gt;) are  not changed 
for almost exactly &lt;CMD -duration&gt; seconds,
the subdirectory will be sent via HTTP to &lt;CMD arg -url&gt;. 

## Usage

USAGE:
efw [OPTIONS] --root-path <ROOT_PATH> --ip <IP> --user <USER> --target-path <TARGET_PATH>

OPTIONS:
-d, --duration-in-sec <DURATION_IN_SEC>
duration how long a file must not be changed before sent [default: 300]

    -h, --help
            Print help information

    --ip <IP>
            ip address of target file server

    -p, --password <PASSWORD>
            TFP - user password

        --port <PORT>
            Network port to use [default: 21]

    -r, --root-path <ROOT_PATH>
            Root directory to be watched

    -t, --target-path <TARGET_PATH>
            FTP target path

    -u, --user <USER>
            TFP - Network user

    -v, --verbose
            set -v to get verbose logging

    -V, --version
            Print version information

    -z, --zipped
            Only considered if result are stored in a folder. If zipped is set the result folder
            will be transferred as zip file
```