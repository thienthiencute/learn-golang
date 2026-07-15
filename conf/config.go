package conf

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
)

var (
	Domain          = "http://localhost:3007"
	DefaultImg      = "/default.jpg"
	DefaultImgAdmin = "/user-dummy-img.jpg"
	UploadPathPublic = "public"
)

func init() {

	if err := ReadFileEnv(); err != nil {
		panic("read file env")
	}

	flag.StringVar(&Domain, "domain", "http://localhost:3007", "domain")
	flag.StringVar(&DefaultImg, "default-img", "/default.jpg", "path domain, should be domain")
	flag.StringVar(&DefaultImgAdmin, "default-img-admin", "/user-dummy-img.jpg", "path domain, should be domain")

	flag.Parse()

}

func ReadFileEnv() error {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}
	_, err := os.Stat(envFile)
	if err == nil {
		err := godotenv.Load(envFile)
		if err != nil {
			return err
		}
	}
	return nil
}
