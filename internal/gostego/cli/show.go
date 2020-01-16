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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zWaR/gostego/pkg/gostego/providers"
)

var imageFile string
var length int

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays the message that's hidden in the image.",
	Long: `Displays the message hidden in the image. If the optional length parameter is specified,
	it will return only first length number of characters of the extracted string. If the length parameter
	is not give, it will return all printable ASCII characters between the char codes 31 and 123. Please
	note that if the length is not given, result will include a lot of garbage.`,
	Run: func(cmd *cobra.Command, args []string) {
		stegoService := providers.CreateStegoService()
		text := stegoService.Show(imageFile, length)

		fmt.Println("Message is:")
		fmt.Println(text)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().StringVarP(&imageFile, "image", "i", "", "Image to extract the message from (required).")
	showCmd.Flags().IntVarP(&length, "length", "l", 0, "Length of the embedded text (optional).")

	showCmd.MarkFlagRequired("image")
	rootCmd.AddCommand(showCmd)
}
