package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/gin-gonic/gin"
)

func init() {
	rebuild := flag.Bool("rebuild", false, "Rebuild project and dependencies")
	flag.Parse()

	var err error
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Don't have permission to read FS")
	}
	_, err = os.Stat(path.Join(dir, "dist", "index.html"))
	if err == nil && !*rebuild {
		log.Println("Seems like project is already build, skipping this step.")
		return
	}
	if *rebuild {
		log.Println("Re-building project...")
		log.Println("Performing cleanup..")
		err = os.RemoveAll(path.Join(dir, "node_modules"))
		if err != nil {
			log.Fatalln("Could not delete node_modules")
		}
		err = os.RemoveAll(path.Join(dir, "dist"))
		if err != nil {
			log.Fatalln("Could not delete dist")
		}
		log.Print("Done!\n")
	}
	out, err := exec.Command("node", "-v").Output()
	if err != nil {
		log.Fatalln("Node not installed")
	}
	fmt.Println()
	log.Println("Node.js", string(out))
	log.Println("Attempting to build project dependencies..")
	if err = exec.Command("npm", "install").Run(); err != nil {
		log.Fatalln("Could not build dependencies!")
	}
	log.Println("Done!")
	log.Println("Building project...")
	if err = exec.Command("npm", "run", "build").Run(); err != nil {
		log.Fatalln("Could not build project!")
	}
	log.Println("Done!")
}

func main() {
	g := gin.New()
	g.Static("/js", "./dist/js")
	g.Static("/css", "./dist/css")
	g.Static("/img", "./dist/img")
	g.StaticFile("/favicon.ico", "./dist/favicon.ico")
	g.GET("/", htmlSupplier)
	g.GET("/blogs", htmlSupplier)
	g.Run(":3000")
}

func htmlSupplier(c *gin.Context) {
	file, _ := ioutil.ReadFile("./dist/index.html")
	c.Data(200, "text/html", file)
}
