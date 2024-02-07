package util

import (
	"archive/zip"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func DownloadFile(
	URL, filePath string,
	ignoreSignatures bool,
) error {
	// Ensure the directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// Download the file
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}

	if _, err = io.Copy(out, resp.Body); err != nil {
		out.Close() // Close the file handle explicitly in case of an error
		return err
	}
	out.Close() // Close the file handle

	// Remove signature files if it's a JAR
	if ignoreSignatures && strings.HasSuffix(filePath, ".jar") {
		if err := RemoveSignatureFiles(filePath); err != nil {
			return err
		}
	}

	return nil
}

func DownloadFileWithoutSignature(URL, filePath string) error {
	return DownloadFile(URL, filePath, true)
}

func RemoveSignatureFiles(jarPath string) error {
	tempJarPath := jarPath + ".tmp"

	// Rename original JAR to temporary
	if err := os.Rename(jarPath, tempJarPath); err != nil {
		return err
	}

	// Open the temporary JAR file
	r, err := zip.OpenReader(tempJarPath)
	if err != nil {
		return err
	}
	defer r.Close()

	// Create a new JAR file
	newJar, err := os.Create(jarPath)
	if err != nil {
		return err
	}
	defer newJar.Close()

	w := zip.NewWriter(newJar)
	defer w.Close()

	manifestContent := ""

	for _, f := range r.File {
		if strings.HasPrefix(f.Name, "META-INF/") &&
			(strings.HasSuffix(f.Name, ".SF") ||
				strings.HasSuffix(f.Name, ".RSA") ||
				strings.HasSuffix(f.Name, ".DSA")) {
			continue // Skip signature files
		}

		if f.Name == "META-INF/MANIFEST.MF" {
			// Read and modify the MANIFEST.MF
			manifestContent, err = readAndModifyManifest(f)
			if err != nil {
				return err
			}
		} else {
			if err := reAddFileToJar(w, f); err != nil {
				return err
			}
		}
	}

	// Add the modified MANIFEST.MF back into the JAR
	if manifestContent != "" {
		if err := addModifiedManifest(w, manifestContent); err != nil {
			return err
		}
	}

	if err := r.Close(); err != nil {
		return err
	}

	// Delete the temporary file
	if err := os.Remove(tempJarPath); err != nil {
		return err
	}

	return nil
}

func readAndModifyManifest(f *zip.File) (string, error) {
	rc, err := f.Open()
	if err != nil {
		return "", err
	}
	defer rc.Close()

	// Read the content of the MANIFEST.MF file
	manifestBytes, err := io.ReadAll(rc)
	if err != nil {
		return "", err
	}
	manifestContent := string(manifestBytes)

	// Remove SHA-256-Digest entries
	manifestContent = removeSHA256DigestEntries(manifestContent)

	return manifestContent, nil
}

func removeSHA256DigestEntries(manifestContent string) string {
	// Split the content into lines
	lines := strings.Split(manifestContent, "\n")
	linesCount := len(lines)
	var newLines []string

	for i := 0; i < linesCount; i++ {
		line := lines[i]
		if strings.Contains(line, "Name:") {
			j := i + 1
			for ; j < linesCount; j++ {
				if strings.Contains(lines[j], "SHA-256-Digest:") {
					break
				}
			}
			i += j - i
			if j+1 < linesCount && strings.TrimSpace(lines[j+1]) == "" {
				i++
			}
		} else {
			newLines = append(newLines, line)
		}
	}

	// Reassemble the manifest content
	return strings.Join(newLines, "\n")
}

func addModifiedManifest(w *zip.Writer, manifestContent string) error {
	header := &zip.FileHeader{
		Name:   "META-INF/MANIFEST.MF",
		Method: zip.Deflate,
	}
	writer, err := w.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(manifestContent))
	return err
}

func reAddFileToJar(w *zip.Writer, f *zip.File) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	fileHeader, err := zip.FileInfoHeader(f.FileInfo())
	if err != nil {
		return err
	}
	fileHeader.Name = f.Name
	fileHeader.Method = f.Method

	writer, err := w.CreateHeader(fileHeader)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, rc)
	return err
}

func GetFromJson(url string, target interface{}) error {
	time.Sleep(time.Second) // моджанги терпеть не могут частые запросы, после первого запроса все сдыхает

	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
