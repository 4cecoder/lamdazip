package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strings"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var (
    configFile string
    rootCmd    = &cobra.Command{
        Use:   "lambda-packager",
        Short: "A tool to package Lambda functions",
        Run:   run,
    }
)

func init() {
    cobra.OnInitialize(initConfig)
    rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.lambda-packager.yaml)")
}

func initConfig() {
    if configFile != "" {
        viper.SetConfigFile(configFile)
    } else {
        home, err := os.UserHomeDir()
        cobra.CheckErr(err)

        viper.AddConfigPath(home)
        viper.SetConfigType("yaml")
        viper.SetConfigName(".lambda-packager")
    }

    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}

func run(cmd *cobra.Command, args []string) {
    // Lambda function names
    functionNames := viper.GetStringSlice("function_names")

    // Virtual environment site-packages directory
    sitePackagesDir := viper.GetString("site_packages_dir")

    // Destination directory for packaged Lambda functions
    destDir := viper.GetString("dest_dir")

    for _, functionName := range functionNames {
        // Create a zip file with the function name
        zipFileName := functionName + ".zip"
        zipFile, err := os.Create(zipFileName)
        if err != nil {
            fmt.Printf("Error creating zip file: %v\n", err)
            return
        }
        defer zipFile.Close()

        // Create a new zip writer
        zipWriter := zip.NewWriter(zipFile)

        // Add the function file to the zip archive
        functionFile := filepath.Join("funx", functionName+".py")
        addFileToZip(functionFile, functionName+".py", zipWriter)

        // Add the site-packages directory to the zip archive
        addDirectoryToZip(sitePackagesDir, zipWriter)

        // Close the zip writer
        err = zipWriter.Close()
        if err != nil {
            fmt.Printf("Error closing zip writer: %v\n", err)
            return
        }

        // Move the zip file to the destination directory
        destPath := filepath.Join(destDir, functionName, zipFileName)
        err = os.Rename(zipFileName, destPath)
        if err != nil {
            fmt.Printf("Error moving zip file: %v\n", err)
            return
        }
    }

    fmt.Println("Lambda functions packaged successfully.")
}

func addFileToZip(filePath, fileName string, zipWriter *zip.Writer) {
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()

    fileWriter, err := zipWriter.Create(fileName)
    if err != nil {
        fmt.Printf("Error creating file in zip: %v\n", err)
        return
    }

    _, err = io.Copy(fileWriter, file)
    if err != nil {
        fmt.Printf("Error copying file to zip: %v\n", err)
        return
    }
}

func addDirectoryToZip(dirPath string, zipWriter *zip.Writer) {
    filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if info.IsDir() {
            return nil
        }

        relPath := strings.TrimPrefix(path, dirPath+string(os.PathSeparator))
        fileWriter, err := zipWriter.Create(relPath)
        if err != nil {
            fmt.Printf("Error creating file in zip: %v\n", err)
            return err
        }

        file, err := os.Open(path)
        if err != nil {
            fmt.Printf("Error opening file: %v\n", err)
            return err
        }
        defer file.Close()

        _, err = io.Copy(fileWriter, file)
        if err != nil {
            fmt.Printf("Error copying file to zip: %v\n", err)
            return err
        }

        return nil
    })
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
