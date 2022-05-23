# ELN_file_watcher
Version 0.1.1

Once all files in a subdirectory of &lt;CMD arg -src&gt; 
(or a file directly in &lt;CMD arg -src&gt;) are not changed
for about exactly &lt;CMD -duration&gt; seconds, the 
subdirectory is sent to &lt;CMD arg -url&gt; via HTTP.

## Usage

efw -duration &lt;integer&gt; -src &lt;folder path&gt; -url &lt;http url&gt;/ [-zip]

    -duration [int]
            Duration in seconds, i.e., how long a file must not be changed before sent (default 300)

    -src [string]
            Src directory to be watched

    -url [string]
            HTTP url to the file network storage. For example: http://&lt;ip address&gt;:&lt;port&gt;/&lt;upload path&gt;/

    -zip 
            Only considered if result are stored in a folder. If zipped is set the result folder will be transferred as zip file

```