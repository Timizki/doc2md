# Docx Text Extractor

This is a simple REST API server written in Go that extracts text from a .docx file. It uses the `github.com/nguyenthenguyen/docx` library to read the .docx file and the `github.com/gorilla/mux` library to handle HTTP requests.

## Installation

1. Install Go: https://golang.org/doc/install
2. Clone this repository: `git clone https://github.com/timizki/doc2md.git`
3. Navigate to the cloned repository: `cd doc2md`
4. Install the required Go packages: `go get github.com/nguyenthenguyen/docx` and `go get -u github.com/gorilla/mux`
5. Build the program: `go build`

## Usage

1. Run the server: `./doc2md`
2. The server will start and listen on port 8000.
3. To extract text from a .docx file, make a POST request to the `/extract` endpoint and include the .docx file in the form data with the name "document". For example, you can use the following `curl` command:
curl -X POST -F "document=@/path/to/your/document.docx" http://localhost:8000/extract

Replace `/path/to/your/document.docx` with the actual path to your .docx file.

The server will respond with the text content of the .docx file.
