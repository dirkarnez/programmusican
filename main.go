package main

import (
	"archive/zip"
	"encoding/xml"
    "fmt"
    "io"
    "os"
)

type Song struct {
	Devices Devices
}

type Devices struct {
	AudioMixer  AudioMixer
}

type AudioMixer struct {
	Attributes []Attribute `xml:"Attributes"`
}

type Attribute struct{
	Id        string     	`xml:"x:id,attr"`
	Name      string      	`xml:"name,attr"`
	Flags     int      		`xml:"flags,attr"`
}

func main() {
	song := Song{
		Devices: Devices {
			AudioMixer: AudioMixer {
				Attributes: []Attribute{
					Attribute{
						Id: "channels",
						Name: "Channels",
						Flags: 1,
					},
				},
			},
		},
	}
    output := "MySong.song"

    if err := ZipFiles(output, song); err != nil {
        panic(err)
	}
	
    fmt.Println("Zipped File:", output)
}

// ZipFiles compresses one or many files into a single zip archive file.
// Param 1: filename is the output zip file's name.
// Param 2: files is a list of files to add to the zip.
func ZipFiles(filename string, song Song) error {
    newZipFile, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer newZipFile.Close()

    zipWriter := zip.NewWriter(newZipFile)
    defer zipWriter.Close()

	if err = AddToZip(zipWriter, "Devices/audiomixer.xml", song.Devices.AudioMixer); err != nil {
		return err
	}
    return nil
}

func AddToZip(zipWriter *zip.Writer, filename string, data interface{}) error {
	var writer io.Writer
	writer, err := zipWriter.Create(filename)
	if err != nil {
		return err
	}

	output, err := xml.Marshal(data)

	_, err = writer.Write(output)
	if err != nil {
		return err
	}

	return nil
}
