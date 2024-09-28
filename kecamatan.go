package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type KecamatanResult struct {
	NamaKecamatan string
	JumlahSuara   int
	JumlahTPS     int
}

func main() {
	baseDir := "/home/kita/Documents/benermeriah" // Ganti dengan direktori yang sesuai
	csvFile := "results/hasil_kecamatan.csv"

	var kecamatanResults []KecamatanResult

	// Traverse subdirectories of kecamatan (e.g. 01-pintu-rime-gayo, 02-permata, etc.)
	err := filepath.Walk(baseDir, func(kecamatanPath string, kecamatanInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Pastikan kecamatan adalah folder, skip file
		if kecamatanInfo.IsDir() {
			// Hitung total suara dan TPS untuk kecamatan ini
			namaKecamatan := kecamatanInfo.Name()

			// Pastikan kecamatan bukan folder README atau folder lain yang tidak relevan
			if strings.Contains(namaKecamatan, "README") || strings.Contains(namaKecamatan, "sub") {
				return nil
			}

			totalSuara, totalTPS := calculateKecamatanData(kecamatanPath)

			if totalTPS > 0 {
				kecamatanResults = append(kecamatanResults, KecamatanResult{
					NamaKecamatan: namaKecamatan,
					JumlahSuara:   totalSuara,
					JumlahTPS:     totalTPS,
				})
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error traversing kecamatan directories:", err)
		return
	}

	// Tulis hasil ke CSV
	writeToCSVKecamatan(kecamatanResults, csvFile)
	fmt.Println("CSV file created:", csvFile)
}

// Fungsi untuk menghitung total suara dan TPS di dalam sebuah kecamatan
func calculateKecamatanData(kecamatanPath string) (int, int) {
	totalSuara := 0
	totalTPS := 0

	// Traverse subdirectories (kampung) in kecamatan
	err := filepath.Walk(kecamatanPath, func(kampungPath string, kampungInfo os.FileInfo, kampungErr error) error {
		if kampungErr != nil {
			return kampungErr
		}

		// Cek apakah kampung memiliki TPS sub-folder
		if kampungInfo.IsDir() && strings.HasPrefix(kampungInfo.Name(), "20") {
			// Iterasi setiap TPS di kampung ini
			err := filepath.Walk(kampungPath, func(tpsPath string, tpsInfo os.FileInfo, tpsErr error) error {
				if tpsErr != nil {
					return tpsErr
				}

				if tpsInfo.IsDir() && strings.HasSuffix(tpsInfo.Name(), "-tps") {
					totalTPS++

					// Dapatkan suara dari file paslon
					totalSuara += readVoteCount(filepath.Join(tpsPath, "sub"), "paslon-1.txt")
					totalSuara += readVoteCount(filepath.Join(tpsPath, "sub"), "paslon-2.txt")
					totalSuara += readVoteCount(filepath.Join(tpsPath, "sub"), "paslon-3.txt")
				}

				return nil
			})

			if err != nil {
				fmt.Println("Error reading TPS data:", err)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error reading kampung data:", err)
	}

	return totalSuara, totalTPS
}

// Helper function to read vote count from a file
func readVoteCount(path string, fileName string) int {
	filePath := filepath.Join(path, fileName)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", filePath, err)
		return 0
	}

	votes, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		fmt.Println("Error converting vote count:", err)
		return 0
	}
	return votes
}

// Helper function to write results to CSV
func writeToCSVKecamatan(results []KecamatanResult, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"nama kecamatan", "jumlah suara", "jumlah tps"})

	// Write data
	for _, result := range results {
		writer.Write([]string{
			result.NamaKecamatan,
			strconv.Itoa(result.JumlahSuara),
			strconv.Itoa(result.JumlahTPS),
		})
	}
}
