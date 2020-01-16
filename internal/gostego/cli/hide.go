/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cli

import (
	"github.com/spf13/cobra"
	"github.com/zWaR/gostego/pkg/gostego/providers"
)

var (
	textFile string
	image    string
)

// hideCmd represents the hide command
var hideCmd = &cobra.Command{
	Use:   "hide",
	Short: "Hides the message from the given file in an image.",
	Long: `Can hide any kind of message stored in the text. Only ASCII characterss are supported.
	Returned image will be saved in the same location as the given image and will have the suffix
	_steg.`,
	Run: func(cmd *cobra.Command, args []string) {
		stegoService := providers.CreateStegoService()
		stegoService.Hide(image, textFile)
	},
}

func init() {

	hideCmd.Flags().StringVarP(&textFile, "msg-file", "m", "", "File with the message to hide (required)")
	hideCmd.Flags().StringVarP(&image, "image", "i", "", "Image to hide the message to (required)")

	hideCmd.MarkFlagRequired("msg-file")
	hideCmd.MarkFlagRequired("image")
	rootCmd.AddCommand(hideCmd)
}
