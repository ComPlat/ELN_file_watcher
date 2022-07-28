Dim WinScriptHost
Set WinScriptHost = CreateObject("WScript.Shell")
WinScriptHost.Run Chr(34) & "<Full path to run_.bat>" & Chr(34), 0
Set WinScriptHost = Nothing
