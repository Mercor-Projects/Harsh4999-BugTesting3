package kubernetes

import (
	"github.com/civo/cli/common"
	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"

	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var kubernetesListVersionCmd = &cobra.Command{
	Use:     "versions",
	Aliases: []string{"version"},
	Example: `civo kubernetes versions ls`,
	Short:   "List all Kubernetes cluster versions",
	Long: `List all Kubernetes cluster versions.
If you wish to use a custom format, the available fields are:

	* version
	* type
	* default`,
	Run: func(cmd *cobra.Command, args []string) {
		utility.EnsureCurrentRegion()

		client, err := config.CivoAPIClient()
		if common.RegionSet != "" {
			client.Region = common.RegionSet
		}
		if err != nil {
			utility.Error("Creating the connection to Civo's API failed with %s", err)
			os.Exit(1)
		}

		kubeVersions, err := client.ListAvailableKubernetesVersions()
		if err != nil {
			utility.Error("%s", err)
			os.Exit(1)
		}

		ow := utility.NewOutputWriter()
		for _, version := range kubeVersions {
			if version.Type == "deprecated" {
				continue
			}

			ow.StartLine()

			ow.AppendDataWithLabel("version", version.Version, "Version")
			ow.AppendDataWithLabel("type", version.Type, "Type")
			ow.AppendDataWithLabel("default", strconv.FormatBool(version.Default), "Default")
		}

		switch common.OutputFormat {
		case "json":
			ow.WriteMultipleObjectsJSON(common.PrettySet)
		case "custom":
			ow.WriteCustomOutput(common.OutputFields)
		default:
			ow.WriteTable()
		}
	},
}
