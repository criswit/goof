package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/criswit/goof/rando"
)

var (
	rootCmd                    *cobra.Command
	length                     int
	characterSet               string
	includeSymbols             bool
	includeNumbers             bool
	includeLowercaseLetters    bool
	includeUppercaseLetters    bool
	excludeSimilarCharacters   bool
	excludeAmbiguousCharacters bool
	times                      int
)

func main() {
	rootCmd = &cobra.Command{
		Run:   goof,
		Use:   "goof",
		Short: "goofy randomizer",
	}

	rootCmd.PersistentFlags().IntVarP(&length, "length", "l", rando.DefaultConfig.Length, "Length of the password")
	rootCmd.PersistentFlags().StringVar(&characterSet, "characters", rando.DefaultConfig.CharacterSet, "Character set for the config")
	rootCmd.PersistentFlags().BoolVar(&includeSymbols, "symbols", rando.DefaultConfig.IncludeSymbols, "Include symbols")
	rootCmd.PersistentFlags().BoolVar(&includeNumbers, "numbers", rando.DefaultConfig.IncludeNumbers, "Include numbers")
	rootCmd.PersistentFlags().BoolVar(&includeLowercaseLetters, "lowercase", rando.DefaultConfig.IncludeLowercaseLetters, "Include lowercase letters")
	rootCmd.PersistentFlags().BoolVar(&includeUppercaseLetters, "uppercase", rando.DefaultConfig.IncludeSymbols, "Include uppercase letters")
	rootCmd.PersistentFlags().BoolVar(&excludeSimilarCharacters, "exclude-similar", rando.DefaultConfig.ExcludeSimilarCharacters, "Exclude similar characters")
	rootCmd.PersistentFlags().BoolVar(&excludeAmbiguousCharacters, "exclude-ambiguous", rando.DefaultConfig.ExcludeAmbiguousCharacters, "Exclude ambiguous characters")
	rootCmd.PersistentFlags().IntVarP(&times, "times", "n", 1, "How many passwords to generate")

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func goof(_ *cobra.Command, args []string) {
	config := rando.Config{
		Length:                     length,
		CharacterSet:               characterSet,
		IncludeSymbols:             includeSymbols,
		IncludeNumbers:             includeNumbers,
		IncludeLowercaseLetters:    includeLowercaseLetters,
		IncludeUppercaseLetters:    includeUppercaseLetters,
		ExcludeSimilarCharacters:   excludeSimilarCharacters,
		ExcludeAmbiguousCharacters: excludeAmbiguousCharacters,
	}
	g, err := rando.New(&config)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	pwds, err := g.GenerateMany(times)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, pwd := range pwds {
		fmt.Println(pwd)
	}
}
