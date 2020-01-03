package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/metamemelord/portfolio-website/handlers"
)

func init() {
	rebuildAllFlag := flag.Bool("rebuild-all", false, "Rebuild project and dependencies")
	rebuildFlag := flag.Bool("rebuild", false, "Rebuild project")

	flag.Parse()

	rebuildAll := *rebuildAllFlag
	rebuild := *rebuildFlag || rebuildAll

	var err error

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Don't have permission to read FS")
	}
	_, err = os.Stat(path.Join(dir, "dist", "index.html"))

	if err == nil && !(rebuild || rebuildAll) {
		log.Println("Seems like project is already build, skipping this step.")
		return
	}

	rebuildAll = err != nil || rebuildAll

	if rebuildAll {
		log.Println("Performing cleanup for rebuilding..")
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

	log.Println("Node.js", string(out))
	if rebuildAll {
		log.Println("Attempting to build project dependencies..")
		if err = exec.Command("npm", "install").Run(); err != nil {
			log.Fatalln("Could not build dependencies!")
		}
		log.Println("Done!")
	}
	if rebuild || rebuildAll {
		log.Println("Building project...")
		if err = exec.Command("npm", "run", "build").Run(); err != nil {
			log.Fatalln("Could not build project!")
		}
		log.Println("Done!")
	}
}

func main() {
	g := gin.New()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	handlers.Register(g)
	g.Run(":" + PORT)
}
