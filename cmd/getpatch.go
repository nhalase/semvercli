/*
Copyright © 2020 Nicholas J Halase <njhalase@gmail.com>

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
	"github.com/blang/semver/v4"
	"github.com/spf13/cobra"
)

var getPatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Extract the PATCH part of <version>.",
	Long:  "Extract the PATCH part of <version> where <version> is MAJOR.MINOR.PATCH.",
	Args:  ValidateGetCmd,
	Run:   runGetPatchCmd,
}

func runGetPatchCmd(_ *cobra.Command, args []string) {
	version, _ := semver.Make(args[0])
	_, _ = printf("%d", version.Patch)
}

func init() {
	getCmd.AddCommand(getPatchCmd)
}
