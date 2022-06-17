# ELN_file_watcher
*Version 0.1.1*

Once all files in a subdirectory <CMD arg -src> 
(or a file directly in <CMD arg -src>) have not been
modified for about exactly <CMD arg -duration> seconds,
the subdirectory is sent to a remote WebDAV server at <CMD arg -dst>.

**Important** this project has to be compiled with go version 1.10.8. Otherwise, it cannot be guaranteed to run on Win XP.

## Usage

efw -duration &lt;integer&gt; -src &lt;folder&gt; -dst &lt;url&gt;/ -user &lt;username&gt; -pass &lt;password&gt; [-zip]

    -crt [string]
        Path to server TLS certificate. Only needed if the server has a self signed certificate.
    
    -duration [int]
        Duration in seconds, i.e., how long a file must
        not be changed before sent. (default 300)
    
    -src [string]
        Source directory to be watched.
    
    -dst [string]
        WebDAV destination URL. If the destination is on the lsdf, the URL should be as follows:
        https://os-webdav.lsdf.kit.edu/<OE>/<inst>/projects/<PROJECTNAME>/
            <OE>-Organisationseinheit, z.B. kit.
            <inst>-Institut-Name, z.B. ioc, scc, ikp, imk-asf etc.
            <PROJRCTNAME>-Projekt-Name

    -pass [string]
        WebDAV Password

    -user [string]
        WebDAV user
  
    -zip
        Only considered if result are stored in a folder. 
        If zipped is set the result folder will be transferred as zip file.   



  

