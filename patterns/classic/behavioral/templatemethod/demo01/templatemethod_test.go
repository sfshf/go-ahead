package templatemethod

func ExampleHTTPDownloader() {
	var downloader Downloader = NewHTTPDownloader()

	downloader.Download("http://example.com/abc.zip")
	// Output:
	// Prepare downloading ...
	// [HTTP] downloading from <http://example.com/abc.zip> ...
	// [HTTP] saving ...
	// Finish downloading!
}

func ExampleFTPDownloader() {
	var downloader Downloader = NewFTPDownloader()

	downloader.Download("ftp://example.com/abc.zip")
    // Output:
	// Prepare downloading ...
	// [FTP] downloading from <ftp://example.com/abc.zip> ...
	// [FTP] saving ...
	// Finish downloading!
}
