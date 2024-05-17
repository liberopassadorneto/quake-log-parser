package cmd

import (
	"fmt"
	"github.com/liberopassadorneto/quake/handler"
	"github.com/liberopassadorneto/quake/logger"
	"github.com/spf13/cobra"
	"io/ioutil"
	"path/filepath"
)

var uploadFilePath string
var outputFileName string

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a log file and process it",
	Long: `Upload a log file and process it to generate a report in JSON format.

Example usage:
  quake upload --file /path/to/log/file.log --output output.json

Flags:
  -f, --file   Path to the log file to upload
  -o, --output Name of the output JSON file (default is output.json and stored in the report directory)`,
	Run: func(cmd *cobra.Command, args []string) {
		if uploadFilePath == "" {
			fmt.Println("Please provide a file path using the --file option.")
			return
		}

		// Set default output file name if not provided
		if outputFileName == "" {
			outputFileName = "output.json"
		}

		// Ensure the output file is stored in the "report" directory
		outputFilePath := filepath.Join("report", outputFileName)

		h := handler.NewHandler()
		jsonData, err := h.UploadFile(uploadFilePath)
		if err != nil {
			logger.Log.Fatal(err)
		}

		err = ioutil.WriteFile(outputFilePath, jsonData, 0644)
		if err != nil {
			logger.Log.Fatalf("Failed to write JSON to file: %v", err)
		}

		fmt.Printf("JSON data has been written to %s\n", outputFilePath)
	},
}

func init() {
	uploadCmd.Flags().StringVarP(&uploadFilePath, "file", "f", "", "Path to the log file to upload")
	uploadCmd.Flags().StringVarP(&outputFileName, "output", "o", "output.json", "Name of the output JSON file (default is output.json and stored in the report directory)")
	uploadCmd.MarkFlagRequired("file")
}
