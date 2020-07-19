package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metamemelord/portfolio-website/handlers"
	"github.com/metamemelord/portfolio-website/pkg/worker"
)

func init() {
	if os.Getenv("APP_AUTH") == "" {
		log.Fatalln("APP_AUTH not provided")
	}

	if os.Getenv("ENV") == "release" || os.Getenv("GIN_MODE") == "release" {
		return
	}

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

	g.Use(gin.LoggerWithWriter(os.Stdout))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	PORT = ":" + PORT
	handlers.Register(g)
	log.Println("Portfolio running on port", PORT)
	go func() {
		for {
			worker.RefreshData()
			time.Sleep(2 * time.Hour)
		}
	}()
	go worker.KeepAlive(time.Minute)
	_ = g.Run(PORT)
}
