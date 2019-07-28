package cmd

import (
	"github.com/jenkins-x-apps/jx-app-jacoco/internal/logging"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	logLevelOptionName = "log-level"
)

var (
	logger = logging.AppLogger().WithFields(log.Fields{"component": "root"})

	rootCmd = &cobra.Command{
		Use:              "jacoco",
		Short:            "Jenkins-x pipeline extension app to create JaCoCo reports",
		PersistentPreRun: configureLogging,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	logLevel string
)

func init() {
	cobra.OnInitialize(initViper)

	rootCmd.PersistentFlags().StringVar(&logLevel, logLevelOptionName, "", "Setting the log level")
	_ = viper.BindPFlag(logLevelOptionName, rootCmd.PersistentFlags().Lookup(logLevelOptionName))
	viper.SetDefault(logLevelOptionName, "info")

	rootCmd.AddCommand(configureCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(versionCmd)
}

func configureLogging(cmd *cobra.Command, args []string) {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	logLevel := viper.GetString(logLevelOptionName)
	err := logging.SetLevel(logLevel)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "unable to configure logging"))
	}
}

func initViper() {
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
}

// Execute executes the root CLI command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}
