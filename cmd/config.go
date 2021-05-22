/*
Copyright © 2021 Moldy-Community

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
	"github.com/Moldy-Community/CLI/functions"
	"github.com/spf13/cobra"
)

var createToogle bool

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure moldy for best and custom usage :D",
	Long: `Generate a config file with the basic specifications:

Usage:

moldy config --create
moldy cfg -c

This create a MoldyFile.toml in the current directory
You can change the file manualy or by commands ( RECOMMENDED )`,
	Run: func(cmd *cobra.Command, args []string) {
		if createToogle {
			functions.CreateConfigFile()
		}
	},
	Example: "moldy config --create",
	Aliases: []string{"cfg"},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolVarP(&createToogle, "create", "c", false, "Toggle the flag for create the config file")
}
