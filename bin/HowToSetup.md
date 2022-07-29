# Setup the efw on a Windows system
1) Download the **run_example.bat**, the **efw.exe** and the **task_example.vbs** for your system [here](https://github.com/ComPlat/ELN_file_watcher/releases/tag/latest)
    - In the following introduction we will assume a 64 bit Windows system.
2) Copy the **efw_win64.exe** and the **run64_example.bat** to the target directory on your target machine
   - In the following we use the example "C:\Program Files\file_exporter".
3) Replace in the **task_example.vbs**: 
   - "&lt;Full path to run_.bat&gt;" with "C:\Program Files\file_exporter\run64_example.bat"
4) Replace in the **run64_example.bat**:
   - &lt;Path to efw_win64.exe&gt; with "C:\Program Files\file_exporter\"
   - Setup all parameter (hint: use _efw_win64.exe -h_):
   - -dst, -src, -crt, -duration, -user, -pass, -crt, -zip
5) copy the **task_example.vbs** into the startup directory 
   - Hint: **Windows Key + R** to open run and type **shell:startup**. This will open Task Scheduler
   
In case it is not working it might be that either the log file or the source directory can not be accessed by the executing user. In this case please make sure the log file and the src dirctory have the correct read & writing permissions.
A tutorial how to set permissions can be found at: [Microsoft answers](https://answers.microsoft.com/en-us/windows/forum/all/give-permissions-to-files-and-folders-in-windows/78ee562c-a21f-4a32-8691-73aac1415373)

Turn off UAC (User Account Control)


## Alternative setup the efw on a Windows 10 system
1) Download the **run_example.bat** and the **efw.exe** for your system [here](https://github.com/ComPlat/ELN_file_watcher/releases/tag/latest)
   - In the following introduction we will assume a 64 bit Windows system.
2) Copy the **efw_win64.exe** and the **run64_example.bat** to the target directory on your target machine
   - In the following we use the example "C:\Program Files\file_exporter".
3) Replace in the **run64_example.bat**:
   - &lt;Path to efw_win64.exe&gt; with "C:\Program Files\file_exporter\"
   - Setup all parameter (hint: use _efw_win64.exe -h_):
   - -dst, -src, -crt, -duration, -user, -pass, -crt, -zip
4) Create a Scheduled Task to trigger at log on.
   - How to create a Scheduled Task:
   1) Using the **Windows Key + R** to open run and type **taskschd.msc**. This will open Task Scheduler.
   2) Under the actions panel, you can choose to create a back task or create a task. Click **Create Task**.
   3) The **Create Task** screen will appear.
   4) In the Create Task dialog, select the following:
      1) In the **General** (tab) and select **Run with highest privileges** 
      2) In the **Triggers** (tab) press the **New** (button) and add **Begin the task, At log on**
