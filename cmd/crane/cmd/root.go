// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/logs"
	"github.com/spf13/cobra"
)

func init() {
	Root.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable debug logs")
	Root.PersistentFlags().BoolVar(&insecure, "insecure", false, "Allow image references to be fetched without TLS")
}

var (
	verbose  = false
	insecure = false

	// Crane options for this invocation.
	options = []crane.Option{}

	// Root is the top-level cobra.Command for crane.
	Root = &cobra.Command{
		Use:               "crane",
		Short:             "Crane is a tool for managing container images",
		Run:               func(cmd *cobra.Command, _ []string) { cmd.Usage() },
		DisableAutoGenTag: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if verbose {
				logs.Debug.SetOutput(os.Stderr)
			}
			if insecure {
				options = append(options, crane.Insecure)
			}
		},
	}
)
