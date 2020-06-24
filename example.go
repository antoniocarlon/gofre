package main

import gofre "./gofre"

func main() {
	buildCreateDirectoryTask := gofre.BuildCreateDirectoryTask("create_directory", map[string]interface{}{"directory": "images"})

	downloadImageDocTask := gofre.BuildDownloadImageTask("download_image_doc",
		map[string]interface{}{"url": "https://golang.org/doc/gopher/doc.png",
			"outfile": "./images/doc.png"})
	downloadImageDocTask.DependsOn = map[string]gofre.Task{"create_directory_task": buildCreateDirectoryTask}

	downloadImageTalksTask := gofre.BuildDownloadImageTask("download_image_talks",
		map[string]interface{}{"url": "https://golang.org/doc/gopher/talks.png",
			"outfile": "./images/talks.png"})
	downloadImageTalksTask.DependsOn = map[string]gofre.Task{"create_directory_task": buildCreateDirectoryTask}

	wrapperTask := gofre.BuildWrapperTask("wrapper_task")
	wrapperTask.DependsOn = map[string]gofre.Task{"download_image_doc": downloadImageDocTask,
		"download_image_talks": downloadImageTalksTask}

	wrapperTask.Run()
}
