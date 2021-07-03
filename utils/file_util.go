package utils

import (
	"bytes"
	crand "crypto/rand"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/websockets"
	"github.com/tidwall/gjson"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//const FFMPEG_PATH = "/usr/bin/ffmpeg"
const FFMPEG_PATH = "ffmpeg"

//const FFMPEG_PATH = "/usr/local/bin/ffmpeg"

/**
Check if file exist and its size
*/
func CheckIfFileExists(f string) bool {
	if f, err := os.Stat(f); err == nil && f.Size() > 0 {
		return true
	}
	return false
}

/**
Add audio to video and generate video
*/
func AddAudioToVideo(inFile string, outFile string) error {
	// check if input file exists
	if !CheckIfFileExists(inFile) {
		return errors.New("not such file")
	}

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
func ExtractAudioFromVideo(inFile string) (string, error) {
	outFile := inFile + "_audio.mp3"
	// check if input file exists
	if !CheckIfFileExists(inFile) {
		return "", errors.New("not such file")
	}

	// extract audio from video using ffmpeg library
	cmd := exec.Command(FFMPEG_PATH, "-y", "-i", inFile, "-f", "mp3", "-ab", "192000", "-vn", outFile)
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return outFile, nil
}

func WriteLog(message string) {
	date := fmt.Sprintf(time.Now().Format("2006-01-02"))
	var f *os.File
	var err error
	//var logPath = "/home/ecodadys/go/src/github.com/nikola43/ecoapigorm/logs/"
	var logPath = "logs/"

	if CheckIfFileExists(logPath + date + ".txt") {
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

func GenerateHologramVideo(inFile string) (string, error) {
	outFile := inFile + "_holo.mp4"
	command := "/Library/Frameworks/Python.framework/Versions/3.8/bin/python3"

	// check if input file exists
	if !CheckIfFileExists(inFile) {
		return "", errors.New("not such file")
	}

	err := ExecuteSystemCommandVerbose(command, inFile)
	if err != nil {
		return "", err
	}

	return outFile, nil
}

func CompressMP4V2(inFile, outFile string, file interface{}) error {

	a, err := ffmpeg.Probe(inFile)
	if err != nil {
		panic(err)
	}
	totalDuration := gjson.Get(a, "format.duration").Float()

	fmt.Println(totalDuration)

	err = ffmpeg.Input(inFile).
		Output(outFile, ffmpeg.KwArgs{"c:v": "libx264", "preset": "veryslow"}).
		GlobalArgs("-progress", "unix://"+TempSock(totalDuration, file)).
		OverWriteOutput().
		Run()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return nil
}

func RandomUint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}


func TempSock(totalDuration float64, file interface{}) string {
	// serve

	// rand.Seed(time.Now().Unix())



	sockFileName := path.Join(os.TempDir(), fmt.Sprintf("%d_sock", RandomUint64()))
	l, err := net.Listen("unix", sockFileName)
	if err != nil {
		panic(err)
	}

	go func() {
		re := regexp.MustCompile(`out_time_ms=(\d+)`)
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		buf := make([]byte, 16)
		data := ""
		progress := ""
		for {
			_, err := fd.Read(buf)
			if err != nil {
				return
			}
			data += string(buf)
			a := re.FindAllStringSubmatch(data, -1)
			cp := ""
			if len(a) > 0 && len(a[len(a)-1]) > 0 {
				c, _ := strconv.Atoi(a[len(a)-1][len(a[len(a)-1])-1])
				cp = fmt.Sprintf("%.2f", float64(c)/totalDuration/1000000)
			}
			if strings.Contains(data, "progress=end") {
				cp = "done"
			}
			if cp == "" {
				cp = ".0"
			}
			if cp != progress {
				progress = cp

				type VideoConversionProgress struct {
					ID   uint      `json:"id"`
					Progress string      `json:"progress"`
				}


				videoConversionProgress := VideoConversionProgress{
					ID:       file.(models.Video).ID,
					Progress: progress,
				}

				socketEvent := models.SocketEvent{
					Type:   "video",
					Action: "progress",
					Data:   videoConversionProgress,
				}

				b, _ := json.Marshal(socketEvent)
				websockets.SocketInstance.Emit(b)

				if socketError := recover(); socketError != nil {
					log.Println("panic occurred:", socketError)
				}

				fmt.Println("progress: ", progress)
			}
		}
	}()

	return sockFileName
}

func CompressMP4(inFile, outFile string) error {

	// check if input file exists
	if !CheckIfFileExists(inFile) {
		return errors.New("not such file")
	}

	// we can store the output of this in our out variable
	// and catch any errors in err
	//cmd := FFMPEG_PATH + " -i " + inFile +" -y " + outFile
	//fmt.Println(cmd)
	out, err := exec.Command(FFMPEG_PATH, "-i", inFile, "-y", outFile).Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)

	// extract audio from video using ffmpeg library
	// ffmpeg -i input.mp4 -vcodec h264 -acodec aac output.mp4
	//err = ExecuteSystemCommandVerbose(FFMPEG_PATH, "-y", "-i", inFile, "-vcodec", "h264", "-acodec", "aac", outFile)
	// -y -preset veryfast -c:v libx264 -crf 30 -c:a aac tatiana.mp4
	//err := ExecuteSystemCommandVerbose(FFMPEG_PATH, "-i", inFile, "-y", "-preset", "veryfast", "-c:v", "libx264", "-crf", "30", "-c:a", "aac", outFile)
	//err := ExecuteSystemCommandVerbose(FFMPEG_PATH, "-i", inFile, "-y", outFile, " >>", outFile+".txt")

	if err != nil {
		return err
	}

	return err
}

func ExecuteSystemCommandVerbose(commandName string, arg ...string) error {
	cmd := exec.Command(commandName, arg...)
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
		extension == ".jpeg" ||
		extension == ".png" {
		fileType = "image"
	} else if extension == ".mp4" ||
		extension == ".avi" ||
		extension == ".mpg" {
		fileType = "video"

		if uploadMode == 2 {
			fileType = "holographic"
		}
	} else if extension == ".mp3" ||
		extension == ".wav" {
		fileType = "heartbeat"
	} else {
		fileType = ""
	}

	return fileType
}

func ExtractThumbnailFromVideo(inFile string, outFile string) error {

	// check if input file exists
	if !CheckIfFileExists(inFile) {
		return errors.New("not such file")
	}

	// extract audio from video using ffmpeg library
	cmd := exec.Command(FFMPEG_PATH, "-y", "-i", inFile, "-ss", "00:00:05.000", "-vframes", "1", outFile)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return err
}

/*
func CompressImage(inputFilePath, outFilePath string) error {
	options := bimg.Options{
		Quality: 10,
	}

	// open file
	buffer, err := bimg.Read(inputFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// process file
	newImage, err := bimg.NewImage(buffer).Process(options)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// save file
	err = bimg.Write(outFilePath, newImage)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return err
}

*/
