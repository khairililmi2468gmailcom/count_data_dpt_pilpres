// package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"strings"
// )

// type KampungResult struct {
// 	NamaKampung string
// 	JumlahSuara int
// 	JumlahTPS   int
// }

// func main() {
// 	baseDir := "/home/kita/Documents/benermeriah" // Ganti dengan direktori dasar
// 	csvFile := "results/hasil_kampung.csv"

// 	results := []KampungResult{}

// 	// Traverse all subdirectories
// 	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		// Check if the directory contains TPS folders (e.g., 001-tps, 002-tps, etc.)
// 		if strings.Contains(path, "sub") {
// 			kampung := filepath.Base(filepath.Dir(path)) // Get the village (kampung) name
// 			tpsDirs, _ := ioutil.ReadDir(path)           // Read all TPS directories

// 			totalVotes := 0
// 			tpsCount := 0

// 			for _, tpsDir := range tpsDirs {
// 				if tpsDir.IsDir() && strings.Contains(tpsDir.Name(), "-tps") {
// 					tpsPath := filepath.Join(path, tpsDir.Name(), "sub")
// 					tpsCount++

// 					// Get vote counts for paslon-1, paslon-2, paslon-3
// 					totalVotes += readVoteCount(tpsPath, "paslon-1.txt")
// 					totalVotes += readVoteCount(tpsPath, "paslon-2.txt")
// 					totalVotes += readVoteCount(tpsPath, "paslon-3.txt")
// 				}
// 			}

// 			if tpsCount > 0 {
// 				results = append(results, KampungResult{
// 					NamaKampung: kampung,
// 					JumlahSuara: totalVotes,
// 					JumlahTPS:   tpsCount,
// 				})
// 			}
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		fmt.Println("Error traversing directories:", err)
// 		return
// 	}

// 	// Write the results to a CSV file
// 	writeToCSV(results, csvFile)
// 	fmt.Println("CSV file created:", csvFile)
// }

// // Helper function to read vote count from a file
// func readVoteCount(path string, fileName string) int {
// 	filePath := filepath.Join(path, fileName)
// 	data, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		fmt.Println("Error reading file:", filePath, err)
// 		return 0
// 	}

// 	votes, err := strconv.Atoi(strings.TrimSpace(string(data)))
// 	if err != nil {
// 		fmt.Println("Error converting vote count:", err)
// 		return 0
// 	}
// 	return votes
// }

// // Helper function to write results to CSV
// func writeToCSV(results []KampungResult, fileName string) {
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		fmt.Println("Error creating CSV file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	// Write CSV header
// 	writer.Write([]string{"nama kampung", "jumlah suara", "jumlah tps"})

// 	// Write data
// 	for _, result := range results {
// 		writer.Write([]string{
// 			result.NamaKampung,
// 			strconv.Itoa(result.JumlahSuara),
// 			strconv.Itoa(result.JumlahTPS),
// 		})
// 	}
// }
