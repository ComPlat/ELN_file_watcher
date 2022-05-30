# ELN_file_watcher
*Version 0.1.1*

Once all files in a subdirectory of &lt;CMD arg -src&gt; 
(or a file directly in &lt;CMD arg -src&gt;) are not changed
for about &lt;CMD arg -duration&gt; seconds, the 
subdirectory is sent to &lt;CMD arg -url&gt; via HTTP.

**Important** this project has to be compiled with go version 1.10.8. Otherwise, it cannot be guaranteed to run on Win XP.

## Usage

efw -duration &lt;integer&gt; -src &lt;folder path&gt; -post-name &lt;post-field-name&gt; -url &lt;http url&gt;/ [-zip]


    -duration [int]
        Duration in seconds, i.e., how long a file must
        not be changed before sent. (default 300)
    
    -post-name [string]
        The post field name by which the file will be sent. (default "file")
    
    -src [string]
        Source directory to be watched.
    
    -url [string]
        HTTP url to the file network storage. 
        For example: http://<ip address>:<port>/<upload path>/
    
    -zip
        Only considered if result are stored in a folder. 
        If zipped is set the result folder will be transferred as zip file.   

## Receiver

Additionally, this repo contains also a responding file receiver sever. The server is contained in the subdirectory *efw_receiver*.

efw_receiver -dst &lt;destination&gt; -url &lt;upload path&gt; -post &lt;post-field-name&gt; -port &lt;port&gt;

    -dst [string]
        Destination directory where received files are stored.

    -url [string]
        Fixed URL path to upload files <upload path> (default "upload")

    -post [string]
        The post field name by which the file will be sent (default "file")

    -port [string]
        "Server address port. Starts with leading (default ":8080")

  

