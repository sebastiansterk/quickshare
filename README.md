# Quickshare
Quickly upload and share text-based snippets. Quickshare is going to upload text-based files (ASCII characters) to https://hastebin.wiuwiu.de and return the link to the uploaded snippet.

# Usage
```
./quickshare <path/to/file>
```

# Compile and run from source:
```
git clone https://github.com/sebastiansterk/quickshare.git
cd quickshare
go build quickshare.go
```

# Download and run Linux x64 binary:
```
curl https://github.com/sebastiansterk/quickshare/releases/download/0.1.0/quickshare_0.1.0_linux_x86_64 -o quickshare; chmod o+x quickshare; mv quickshare /usr/local/bin/
```
