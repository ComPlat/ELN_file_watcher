# ELN_file_watcher
*Version 1.2*

Once all files in a subdirectory <CMD arg -src>
(or a file directly in <CMD arg -src>) have not been
modified for about exactly <CMD arg -duration> seconds,
the subdirectory is sent to a remote WebDAV or SFTP server at <CMD arg -dst>.

**Important**
1) If you want to run this tool on Windows XP it has to be compiled with go version 1.10.8.
2) If the tool is used with the SFTP protocol under Windows XP, you must also save a protable
   version of WinSCP (WinSCP)[https://winscp.net/download/WinSCP-5.21.5-Portable.zip] in the same folder as the efw.exe.

## Usage

efw -duration &lt;integer&gt; -src &lt;folder&gt; -dst &lt;url&gt;/ -user &lt;username&gt; -pass &lt;password&gt; [-zip]

    -name string
        Name of the EFW instance. This name is used to create a working folder in the user home directory. .
    
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
        WebDAV user or SFTP user
  
    -type [string: file|folder|zip]
         Type must be 'file', 'folder' or 'zip'. The 'file' option means that each file is 
         handled individually, the 'folder' option means that entire folders are transmitted
         only when all files in them are ready. The option 'zip' sends a folder zipped, only
         when all files in a folder are ready.
   
    -crt (Optional) [string]
         Path to server TLS certificate. Only needed if the server has a self signed certificate.

## Setup the efw on a Windows system
2) Download the **efw_run_example.bat**, the **efw.exe** and the **task_example.vbs** for your system [here](https://github.com/ComPlat/ELN_file_watcher/releases/tag/latest)
2) Copy the **efw_{system}.exe** and save it **as efw.exe**. Additionally, download the **efw_run_example.bat** to the target directory on your target machine
   - In the following we use the example "C:\Program Files\file_exporter".
3) Replace in the **task_example.vbs**:
   - "&lt;Full path to run_.bat&gt;" with "C:\Program Files\file_exporter\efw_run_example.bat"
4) Replace in the **efw_run_example.bat**:
   - &lt;Path to efw.exe&gt; with "C:\Program Files\file_exporter\"
   - Setup all parameter (hint: use _efw.exe -h_):
   - -dst, -src, -crt, -duration, -user, -pass, -crt, -zip, -name, -transfer, -type
5) copy the **task_example.vbs** into the startup directory
   - Hint: **Windows Key + R** to open run and type **shell:startup**. This will open Task Scheduler





  

