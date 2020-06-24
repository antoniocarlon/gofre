# Gofre

### Gofre is a simple pipeline builder written in Go.

It can be easily extended to support multiple functionalities by creating custom task builders like the ones in the `helpers.go` file.

**Disclaimer:** Gofre is the result of my efforts of leaning Go by doing, so please take it with a grain of salt because it's just a proof of concept.

Building a pipeline using Gofre is as easy as this (see the `example.go` file for reference):

```
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
```

(This pipeline creates a new directory and then downloads a couple of images to it, using a wrapper task to orchestrate the whole thing).