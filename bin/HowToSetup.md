# Setup the efw on a Windows system
1) Download the **efw_run_example.bat**, the **efw.exe** and the **task_example.vbs** for your system [here](https://github.com/ComPlat/ELN_file_watcher/releases/tag/latest)
2) Copy the **efw_{system}.exe** and sve it **as efw.exe**. Additionally, download the **efw_run_example.bat** to the target directory on your target machine
   - In the following we use the example "C:\Program Files\file_exporter".
3) Replace in the **task_example.vbs**: 
   - "&lt;Full path to run_.bat&gt;" with "C:\Program Files\file_exporter\efw_run_example.bat"
4) Replace in the **efw_run_example.bat**:
   - &lt;Path to efw.exe&gt; with "C:\Program Files\file_exporter\"
   - Setup all parameter (hint: use _efw.exe -h_):
   - -dst, -src, -crt, -duration, -user, -pass, -crt, -zip, -name, -transfer, -type
5) copy the **task_example.vbs** into the startup directory 
   - Hint: **Windows Key + R** to open run and type **shell:startup**. This will open Task Scheduler

Please Note: If the tool is used with the SFTP protocol under Windows XP, you must also save a portable version of [WinSCP](https://winscp.net/download/WinSCP-5.21.5-Portable.zip) in the same folder as the efw.exe.
