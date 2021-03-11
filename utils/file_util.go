package utils

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const FFMPEG_PATH = "/usr/local/bin/ffmpeg"

//const FFMPEG_PATH = "/home/ecodadys/bin/ffmpeg"

/**
Check if file exist and its size
*/
func CheckIfFileExists(f string) error {
	var err error
	if f, err := os.Stat(f); err == nil && f.Size() > 0 {
		return nil
	}
	return err
}

/**
Add audio to video and generate video
*/
func AddAudioToVideo(inFile string, outFile string) error {
	// check if input file exists
	inFileError := CheckIfFileExists(inFile)
	if inFileError != nil {
		fmt.Printf("Error: %s", inFileError.Error())
		return inFileError
	}

	// extract audio from video using ffmpeg library
	cmd := exec.Command(FFMPEG_PATH, "-y", "-i", inFile, "-i", "music_video.m4a", "-c", "copy", "-shortest", outFile)
	//cmd := execontext.Command("ffmpeg", "-i", inFile, "-i", "/home/ecodadys/go/src/github.com/nikola43/ecoapigorm/video_musicontext.mp3", "-c", "copy", "-map", "0:v", "-map", "1:a", outFile)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error injectando audio " + err.Error())
		return err
	}

	return nil
}

/**
Extract audio from video and generate audio file
*/
func ExtractAudioFromVideo(inFile string, outFile string) error {
	// check if input file exists
	inFileError := CheckIfFileExists(inFile)
	if inFileError != nil {
		fmt.Printf("Error: %s", inFileError.Error())
		return inFileError
	}

	// check if output file exists
	outFileError := CheckIfFileExists(outFile)
	if outFileError != nil {
		fmt.Printf("Error: %s", outFileError.Error())
		return outFileError
	} else {
		// // if exists then remove
		// removeFileError := os.Remove(outFile)
		// if removeFileError != nil {
		// 	fmt.Printf("Error: %s", removeFileError.Error())
		// 	return removeFileError
		// }
	}

	// extract audio from video using ffmpeg library
	cmd := exec.Command(FFMPEG_PATH, "-y", "-i", inFile, "-f", "mp3", "-ab", "192000", "-vn", outFile)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

/**
Extract audio from video and generate audio file
*/
func ExtractThumbnailFromVideo(inFile string, outFile string) error {
	// check if input file exists
	inFileError := CheckIfFileExists(inFile)
	if inFileError != nil {
		fmt.Printf("Error: %s", inFileError.Error())
		return inFileError
	}

	// check if output file exists
	outFileError := CheckIfFileExists(outFile)
	if outFileError != nil {
		fmt.Printf("Error: %s", outFileError.Error())
		return outFileError
	} else {
		// // if exists then remove
		// removeFileError := os.Remove(outFile)
		// if removeFileError != nil {
		// 	fmt.Printf("Error: %s", removeFileError.Error())
		// 	return removeFileError
		// }
	}

	// extract audio from video using ffmpeg library
	cmd := exec.Command(FFMPEG_PATH, "-y", "-i", inFile, "-ss", "00:00:05.000", "-vframes", "1", outFile)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func WriteLog(message string) {
	date := fmt.Sprintf(time.Now().Format("2006-01-02"))
	var f *os.File
	var err error
	//var logPath = "/home/ecodadys/go/src/github.com/nikola43/ecoapigorm/logs/"
	var logPath = "logs/"
	if fileExists(logPath + date + ".txt") {
		f, err = os.OpenFile(logPath+date+".txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		f, err = os.Create(logPath + date + ".txt")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	newLine := message + " | " + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))

	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
}

func GenerateHologramVideo(inFile string, outFile string) error {

	command := "/Library/Frameworks/Python.framework/Versions/3.8/bin/python3"

	// check if input file exists
	err := CheckIfFileExists(inFile)
	if err != nil {
		return err
	}

	err = executeCommandVerbose(command, inFile)

	if err != nil {
		return err
	}
	err = CheckIfFileExists(outFile)
	if err != nil {
		return err
	}
	return nil
}

func CompressMP4(inFile string, outFile string) error {
	// check if input file exists
	err := CheckIfFileExists(inFile)
	if err != nil {
		return err
	}

	/*
		// check if input file exists
		err = CheckIfFileExists(outFile)
		if err != nil {
			return err
		} else {
			//if exists then remove
			fmt.Println("output file called " + outFile + " already exist. Removing...")
			removeError := os.Remove(outFile)
			if removeError != nil {
				return removeError
			}
			fmt.Println("file " + outFile + " has been removed successfully")
		}
	*/

	// extract audio from video using ffmpeg library
	// ffmpeg -i input.mp4 -vcodec h264 -acodec aac output.mp4
	//err = executeCommandVerbose(FFMPEG_PATH, "-y", "-i", inFile, "-vcodec", "h264", "-acodec", "aac", outFile)
	// -y -preset veryfast -c:v libx264 -crf 30 -c:a aac tatiana.mp4
	err = executeCommandVerbose(FFMPEG_PATH, "-i", inFile, "-y", "-preset", "veryfast", "-c:v", "libx264", "-crf", "30", "-c:a", "aac", outFile)

	if err != nil {
		return err
	}
	err = CheckIfFileExists(outFile)
	if err != nil {
		return err
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func executeCommandVerbose(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	fmt.Println("Result: ")
	fmt.Println(out.String())
	return nil
}

// use godot package to load/read the .env file and
// return the value of the key
func GetEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GetFileType(file string, uploadMode uint) string {
	// fileType image -> 1
	// fileType video -> 2
	// fileType holo -> 3
	// fileType heartbeat -> 4

	fileType := ""
	extension := filepath.Ext(file)

	if extension == ".jpg" ||
		extension == ".png" {
		fileType = "image"
	} else if extension == ".mp4" ||
		extension == ".avi" ||
		extension == ".mpg" {
		fileType = "video"

		if uploadMode == 2 {
			fileType = "holo"
		}
	} else if extension == ".mp3" ||
		extension == ".wav" {
		fileType = "heartbeat"
	} else {
		fileType = ""
	}

	return fileType
}
