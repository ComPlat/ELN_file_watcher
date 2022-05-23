package main

import (
	"os"
	"path/filepath"
	"time"
)

// ProcessManager manges the file watching process.
// As soon as all files in a subdirectory of <CMD arg -src>
// (or a file directly in <CMD arg -src>) are
// not changed for almost exactly <CMD -duration> seconds,
// the subdirectory will be pushed into the channel 'done_files'.
type ProcessManager struct {
	args       *Args
	done_files chan string
}

// doWork runs in a endless loop. It watches the files in the <CMD arg -src> directory.
// It terminates as soon as a value is pushed into quit. Run in extra goroutine.
func (m ProcessManager) doWork(quit chan int) {
	InfoLogger.Println("Started watch process.")
	for {
		select {
		case <-quit:
			return
		default:
			now := time.Now()
			done_folders := make(map[string]bool)
			// Checking all files in <CMD arg -src>.
			err := filepath.Walk(m.args.src,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if !info.IsDir() {
						modifiedtime := info.ModTime()
						diff := now.Sub(modifiedtime)
						if diff < 2*m.args.duration {
							if relpath, err := filepath.Rel(m.args.src, path); err == nil {
								folder := get_root_dir(relpath)

								if _, ok := done_folders[folder]; !ok {
									done_folders[folder] = true
								}
								if diff <= m.args.duration {
									done_folders[folder] = false
								}
							} else {
								ErrorLogger.Println(err)
							}
						}
					}
					return nil
				})

			if err != nil {
				ErrorLogger.Println(err)
			}

			// Pushing all complete subdirectory into done_files channel.
			for k, v := range done_folders {
				if v {
					InfoLogger.Println("Folder ready to send: ", k)
					m.done_files <- filepath.Join(m.args.src, k)
				}
			}

			time.Sleep(m.args.duration - time.Now().Sub(now))
		}
	}
}

// newProcessManager factory for ProcessManager struct
func newProcessManager(args *Args, done_files chan string) ProcessManager {
	return ProcessManager{args: args, done_files: done_files}
}
