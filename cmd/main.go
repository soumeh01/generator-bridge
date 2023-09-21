/*
 * Copyright (c) 2023 Arm Limited. All rights reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stringer",
	Short: "stringer - a simple CLI to transform and inspect strings",
	Long: `stringer is a super fancy CLI (kidding)
   
One can use stringer to modify or inspect strings straight from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func main() {
	uuidWithHyphen, _ := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
	fmt.Println("Hello, World!")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
