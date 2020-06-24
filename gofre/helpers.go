package gofre

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Build a generic wrapper task
func BuildWrapperTask(taskId string) Task {
	return Task{
		Id: taskId,
		Execute: func(task *Task, c chan bool) {
			c <- true
		},
	}
}

// Build a generic task to create directories
func BuildCreateDirectoryTask(taskId string, params map[string]interface{}) Task {
	return Task{
		Id:     taskId,
		Params: params,
		Execute: func(task *Task, c chan bool) {
			path := filepath.Join(".", task.getParams()["directory"].(string))
			os.MkdirAll(path, os.ModePerm)

			c <- true
		},
		Check: func(task *Task) bool {
			info, err := os.Stat(task.getParams()["directory"].(string))
			return !os.IsNotExist(err) && info.IsDir()
		},
	}
}

// Build a generic task to download files
func BuildDownloadImageTask(taskId string, params map[string]interface{}) Task {
	return Task{
		Id:     taskId,
		Params: params,
		Execute: func(task *Task, c chan bool) {
			out, err := os.Create(task.getParams()["outfile"].(string))
			if err != nil {
				c <- false
			}
			defer out.Close()
			resp, err := http.Get(task.getParams()["url"].(string))
			if err != nil {
				c <- false
			}
			defer resp.Body.Close()
			n, err := io.Copy(out, resp.Body)
			if err != nil || n <= 0 {
				c <- false
			}

			c <- true
		},
		Check: func(task *Task) bool {
			_, err := os.Stat(task.getParams()["outfile"].(string))
			return !os.IsNotExist(err)
		},
	}
}
