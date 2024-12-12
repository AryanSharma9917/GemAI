package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

var (
	imageFilePath   string
	imageFileFormat string
)

var imageCmd = &cobra.Command{
	Use:     "image [your question] --path [image path] --format [image format]",
	Example: "gencli image 'What this image is about?' --path cat.png --format png",
	Short:   "Know details about an image (Please put your question in quotes)",
	Long:    "Ask a question about an image and get a response. You need to provide the path of the image and the format of the image. The supported formats are jpg, jpeg, png, and gif.",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := imageFunc(args)
		fmt.Println(res)
	},
}




