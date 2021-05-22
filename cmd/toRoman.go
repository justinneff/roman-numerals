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
	"strconv"

	"github.com/justinneff/roman-numerals/convert"
	"github.com/spf13/cobra"
)

var input int

// toRomanCmd represents the toRoman command
var toRomanCmd = &cobra.Command{
	Use:   "toRoman [number]",
	Short: "Converts an Arabic number to a Roman numeral",
	Long: `Converts an Arabic number to a Roman numeral.
	
For example: 12 => XII`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a number argument")
		}
		i, err := strconv.Atoi(args[0])
		if err != nil || !convert.CanConvertToRoman(i) {
			return fmt.Errorf("requires an integer number between %d and %d", convert.MinArabic, convert.MaxArabic)
		}
		input = i
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		result := convert.ToRoman(input)
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(toRomanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toRomanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toRomanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
