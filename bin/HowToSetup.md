# Setup the efw on a Windows system
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