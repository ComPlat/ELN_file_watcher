# ELN_file_watcher
*Version 1.1*

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

## Setup the efw on a Windows system
1) Download the **run_example.bat**, the **efw.exe** and the **task_example.vbs** for your system [here](https://github.com/ComPlat/ELN_file_watcher/releases/tag/latest)
    - In the following introduction we will assume a 64 bit Windows system.
2) copy the efw.exe and the run_example.bat to the target directory on your target machine
    - In the following we use the example "C:\Program Files\file_exporter".
3) Replace in the **task_example.vbs**:
    - "&lt;Full path to run_.bat&gt;" with "C:\Program Files\file_exporter\run64_example.bat"
4) Replace in the **run64_example.bat**:
    - &lt;Path to efw_win64.exe&gt; with "C:\Program Files\file_exporter\"
    - Setup all parameter (hint: use _efw_win64.exe -h_):
    - -dst, -src, -crt, -duration, -user, -pass, -crt, -zip
5) copy the **task_example.vbs** int the startup directory
    - Hint to open startup folder:
    1) press: **Win + R**
    2) type: **shell:startup**
    3) press: **enter**



  

