package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/nguyenthenguyen/docx"
	"github.com/gorilla/mux"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form in the request
	err := r.ParseMultipartForm(10 << 20) // Max memory 10MB
	if err != nil {
		http.Error(w, "Error parsing multipart form: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve the file from form data
	file, _, err := r.FormFile("document") // This needs to be the name used in the form
	if err != nil {
		http.Error(w, "Error retrieving the file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Create a temporary file within our temp-documents directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-documents", "upload-*.docx")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")

	// Now we read the document
	docxReader, err := docx.ReadDocxFile(tempFile.Name())
	if err != nil {
		http.Error(w, "Error opening document: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer docxReader.Close()

	doc := docxReader.Editable()
	fmt.Fprintf(w, "Content: %s\n", doc.GetContent())
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/extract", handleRequest).Methods("POST")
	http.ListenAndServe(":8000", r)
}
