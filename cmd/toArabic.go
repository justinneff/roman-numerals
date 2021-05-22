/*
Copyright © 2021 Justin Neff <neffjustin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/justinneff/roman-numerals/convert"
	"github.com/spf13/cobra"
)

// toArabicCmd represents the toArabic command
var toArabicCmd = &cobra.Command{
	Use:   "toArabic",
	Short: "Converts a Roman numeral to an Arabic number",
	Long: `Converts a Roman numeral to a base 10 Arabic number
	
For example: XIV => 14
	`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a single string argument")
		}

		if !convert.CanConvertToArabic(args[0]) {
			return errors.New("string is not valid Roman numeral")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		result := convert.ToArabic(args[0])
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(toArabicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toArabicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toArabicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
