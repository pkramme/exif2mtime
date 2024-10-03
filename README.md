# exif2mtime

If you have a JPEG collection that has been copied badly and has lost its original file creation and modification timestamps, this tool resets file modification dates to the EXIF timestamp. This allows operating systems to sort your images by when the image was created, not by when it was last copied.

Also this tool detects JPEGs based on MIME sniffing, not file extension and can automatically rewrite the file extension to `.jpg` if its something else like `.jpeg` or `.JPG`. 

# Usage
Clone and build with `go build`. Then,
```
./exif2mtime --doit file1.jpg file2.jpg
```
and to additionally rewrite the file extension
```
./exif2mtime --fixext --doit file1.jpg file2.jpg
```

Output is something like
```
2016/2016IMG_4431.JPG exif creation date 2016-07-22 22:57:16 +0000 UTC
2016/2016IMG_4432.JPG exif creation date 2016-07-23 00:30:46 +0000 UTC
2016/2016IMG_4433.JPG exif creation date 2016-07-23 00:30:50 +0000 UTC
2016/2016IMG_4434.JPG exif creation date 2016-07-23 00:31:02 +0000 UTC
2016/2016IMG_4435.JPG exif creation date 2016-07-23 00:32:10 +0000 UTC
2016/2016IMG_4436.JPG exif creation date 2016-07-23 00:32:21 +0000 UTC
```

# Note
This was written really quickly and worked for me, in my specific setup and might not work for you or even damage your files. Don't trust it, read the code and have backups.
