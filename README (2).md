---
title: The CICD Project - Status
---

<div>
        <a href="/" >
          <button class= "hover:bg-blue-200 rounded-full ">
            <img src="/images/homeIcon_2.png" width="240" height="150" /></button
        ></a>
  </div>

<h1> &#x1F537; Status  </h1>


## Plan


### :diamond_shape_with_a_dot_inside: Overview

| Order | Status             | Action Item |
| :---: | :----------------- | :--------------------------------- |
|   1   | :footprints: | Are the Dependencies completed? |
|   2   | :footprints: | Has the Softwawre Product been deployed to Production for use by customers? |
|   3   | :footprints: | Have the Users Approved the solution? |
|   4   | :footprints: | Have the Users decided on if a Version 2 is required and if so, what are the additional feature requests and bug fixes? |

### :diamond_shape_with_a_dot_inside: Tasks & Dependencies

| Order | Status             | Action Item | Acountable | ETA | Latest |
| :---: | :----------------- | :------------------- | :----------- | :----------- | :----------- |
|   1   | :footprints: | Move all Stored Procedures to Github | Lauren/Matt | 7/10/2022  |   |
|   2   | :footprints: |  TBD. | TBD. | TBD. |

### :diamond_shape_with_a_dot_inside: Unforseen Potential Issues

| Order | Status             | History ETA  | Updates & Reasons |
| :---: | :----------------- | :--------------------------------- | :--------------------------------- |
|   1   | :question: :stop_sign: | TBD. | Potentially not a requirement | 
|   2   | :footprints: | Online User Manual may be required (Documentation Intern / Marcus Training) |  Marcus to train documentation Intern |
|   3   | :footprints: | Online Tour may be required (Documentation Intern / Marcus Training) |Marcus to train documentation Intern |

#### Initial Project Ideation and Slide

GTS Development Team Start Date: 7/1/2022

Powerpoint Slide template for End of Summer Demonstration

### :diamond_shape_with_a_dot_inside: Notes

## Class Room

#### Docker

1. Dockerfile

2. Docker build

```sh

$ docker build

```

3. Docker run

```sh

$ docker run

```

:::warning  Notes

Vue is a SPA - Single page application framework that makes creaing HTML, CSS and Javascript frontends simpler and more efficient.

For more information: 

[WIKI - Single-page application](https://en.wikipedia.org/wiki/Single-page_application)


[What is a Single Page Application](https://www.codecademy.com/article/fecp-what-is-a-spa)

"vuepress" is a static site generator written in vue.js and was used to create the original Team Portal web application.

[Vuepress](https://vuepress.vuejs.org/)

:::

#### Example:  Below is the dockerfile for the teamportal vue based web application.


```sh

FROM node:16.12.0 as build-stage
RUN mkdir app
WORKDIR /app
COPY package.json yarn.lock ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx as production-stage
RUN mkdir /app
COPY --from=build-stage /app/src/.vuepress/dist /app
COPY nginx.conf /etc/nginx/nginx.conf

```

#### Build a CLI application in the GOLANG computer language.

:::warning Notes

[What is a CLI?](https://www.w3schools.com/whatis/whatis_cli.asp)


:::

### :diamond_shape_with_a_dot_inside: Golang CLI Application

Hello World - Simplest Golang CLI Application - also known as a Terminal Application (Mac and Unix) or a Command Prompt Application in Windows

[A tour of GO - Interactive Learning of GOLANG language on the web](https://go.dev/tour/welcome/1)

![Image](../../.vuepress/public/images/interns_dockerize_3.png "interns_dockerize_3.png")


![Image](../../.vuepress/public/images/interns_dockerize_4.png "interns_dockerize_4.png")


#### Build a web application in the GOLANG computer language and then run that application into a Docker Container.


### :diamond_shape_with_a_dot_inside: Golang Web Application

Create a basic golang web application program.

The code illustrated below will create a basic static and stateless web application to test that our environment setup is working and that we can package our code into Docker.

```go

package main

import (
	"log"
	"net/http"
)

// main - set up a function handler to send some HTML and CSS to a browser
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
        <html>
        <head>
		<title>GOLANG TEST</title>
		</head>
		<style type="text/css">
        .cell-green {
            background: rgba(0,255,0,0.1);
        }
		</style>
        <body>
		<p class="cell-green">Hello World</p>
		</body>
        </html>`))
	})

	//start the webserver and listen on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("LisenAndServer:", err)
	}
}
```

:::tip More Information

[Golang By Example](https://gobyexample.com/)

:::

### :diamond_shape_with_a_dot_inside: Building (Compiling) your program

#### Building your program for the MAC

:::warning NOTES

GOLANG ARCHITECURE - (GOARCH) i.e. the chips that you can compile for

GOLANG OPERATING SYSTEMS - (GOOS) i.e. the operating systems that you can complile for

more information ...

[Golang Documentation](https://go.dev/doc/install/source#environment)

[How to cross-compile Go programs for Windows, macOS, and Linux](https://freshman.tech/snippets/go/cross-compile-go-programs/)

:::


Build your Golang program into a binary file by running the following Golang Build command.

::: tip

Note that it is required to have the GOOS=darwin for running on the MAC operating System

:::



```
$ CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o tutorialexample tutorialexample.go
```

### :diamond_shape_with_a_dot_inside: Run your program

```sh
$ go run tutorialexample.go
```

You can also execute your compiled program at the terminal prompt

```
./tutorialexample
```

![Image](../../.vuepress/public/images/interns_dockerize_1.png "interns_dockerize_1.png")



#### Output

#### At localhost:8080 your browser will now direct to the following page as illustrated below:

![Image](../../.vuepress/public/images/interns_dockerize_2.png "interns_dockerize_2.png")


<!-- 
![Image](../.vuepress/public/images/golangTutorialexamplerun2.png "Running the golang code")
-->

### :diamond_shape_with_a_dot_inside: Building the program for Docker

::: tip

Note that it is required to have the GOOS=linux for docker images that are using a linux operating system like "ALPINE"

:::

#### Build your Golang program into a binary file by running the following Golang Build command.

```
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tutorialexample tutorialexample.go
```

After a successful build, the binary file - "tutorialexample" is created.

<!-- 
![Image](../.vuepress/public/images/tutorialexamplebuildforMAC.png "Note the new file created")
-->

::: danger

If you run ./tutorialexample from your MAC, it will fail with the following error:

This is due to the fact that the docker image will create a linux environment and we set our Dockerfile GOOS setting = linux.

:::

<!-- 
![Image](../.vuepress/public/images/LinuxNotDarwinErrorOnMac.png "Running the golang code from executable with incorrect GOOS")
-->


### :diamond_shape_with_a_dot_inside:  Create a Dockerfile.

Create a Dockerfile in the same directory as your code.

This step is composed of editing a Dockerfile and building that Dockerfile into a docker container that can be run locally via the **$ docker run** command.

Inorder to put our code into a docker image formatted file, we will first create a Dockerfile in the same directory as our source code location. The Dockerfile is a simple text file with no extension.

```
FROM scratch
COPY tutorialexample /
CMD ["/tutorialexample"]
EXPOSE 8080
```
### :diamond_shape_with_a_dot_inside: Build a Dockerfile

In order to build your code into a docker image format, run the following command:

```
$ docker build -t tutorialexample:v1.0.1 .
```

**Note** The -t parameter option lets us specify a name and a corresponding tag name. In this case, the image name will be tutorialexample and the tag will be v1.0.1

**Note** The "dot" or "period" in the above command is required..  This signifies that we are building to the current directory that we are in.

To list your local Docker Images, type the following docker command in your terminal window:

```
$ docker images
```

#### Output

```
REPOSITORY                                TAG                 IMAGE ID            CREATED              SIZE
tutorialexample                           v1.0.1              e6f7a15424fd        About a minute ago   7.31MB
registry.ng.bluemix.net/draco/draco-ui    2                   f5374333b8cc        4 months ago         30.1MB
registry.ng.bluemix.net/draco/draco-api   2                   f47bfe5a6bcb        4 months ago         195MB
<none>                                    <none>              b70a946113ca        4 months ago         97.2MB
<none>                                    <none>              666d54a95a03        4 months ago         200MB
<none>                                    <none>              a7d2539cafc4        4 months ago         30.1MB
<none>                                    <none>              fa0f047d89bd        5 months ago         28.6MB
nimmis/alpine-glibc                       latest              e0e2abe2b977        18 months ago        14.3MB
```

#### Tag names often used to display versions.

Note that in the above example, that we added, after the colon, the "tag" name.  In this example, we chose to give our image name the "tag" - "v1.0.1". Typically, version numbers follow a convention such as [semantic versioning](http://semver.org/).

#### If you try to execute the program, after building it for docker, you will get an error

```
./tutorialexample
````

#### Error Message

```
-bash: ./tutorialexample: cannot execute binary file
```

### :diamond_shape_with_a_dot_inside: Run your docker container.

```
$ docker run --rm -p 8080:8080 tutorialexample:v1.0.1
```

The result should be a "hanging cursor".  This simply tells us that the golang webserver is running and waiting for requests on port 8080 inside our docker file. The command also displays that our local desktop port 8080 is (-p) "published" to match with the docker container that is now running in a background process.

To view the running program, open a web browser and link to localhost port 8080 and it will display as below.

