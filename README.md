# Awesome Ascii

Awesome Ascii is a command-line interface (CLI) tool designed to transform images into ASCII text art.
It processes an input image by scaling, converting it to grayscale, and then mapping the pixel values to ASCII characters.
The resulting ASCII art can be output directly to the terminal or saved to a file.
The command also allows for adjusting the concurrency level to optimize performance.

# Example

This is a basic example on how to convert an image to Ascii

```
go run main.go -i images/darkness.png -a extended -w 150
```

---

- ### Before

![Old image placeholder](images/darkness.png)

- ### After

![after image](https://utfs.io/f/9a7666e2-2d5d-410b-9d25-f39c10e799d9-sz54oy.png)

This is a list of Flags you can use to manipulate the output:

```
Flags:
  -a, --ascii-type AsciiCharType   Determine which set of ascii characters will be used (default basic)
  -c, --concurrency int            Set GOMAXPROCS (default 16)
  -h, --help                       help for awesome-ascii
  -i, --input string               An image path which will be converted to ASCII
  -o, --output string              An output path for the converted image
  -v, --version                    version for awesome-ascii
  -w, --width uint16               An image path which will be converted to ASCII (default 154)
```

You can also choose which characters pool to use for conversion, this is the list of Ascii characters supported

- basic (@%#\*+=-:. )
- binary (01)
- contrast (@#S%?\*+;:,. )
- extended (@W#98B0%Zq6x2t!i\*|~-:. )
- high detail (@$B%8&WM#\*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-\*+~<>i!lI;:,\"^`'. )

---

You can achieve the same output in an interactive way, you just need to type the command below to open interactive CLI

```
go run main.go interactive
```

Running the following command will open the CLI in interactive mode

![interactive CLI](https://utfs.io/f/2af7e08f-4d8f-43bf-9ddc-89812a6c71d3-sy3kwy.png)

# Sobel

part of a command-line interface (CLI) application designed to process images using the Sobel edge detection algorithm.

This command reads an input image, scales it to the desired width, converts it to grayscale, and then applies the Sobel algorithm to highlight the edges.

The resulting edge-detected image is then transformed into ASCII art.
Users can choose to display the ASCII art directly in the terminal or save it to a file.

# Example

This is a basic example on how to convert an image to Ascii

```
go run main.go sobel -i images/darkness.png -a extended -w 150 -t 200
```

If we apply this command on the same input this is the output

![sobel image](https://utfs.io/f/3797c6e7-f0ef-4a4c-a03f-133be0f8d16a-sy2xuw.png)

You can find here the list of flags to manipulate the output of a Sobel image

```
Flags:
  -a, --ascii-type AsciiCharType   Determine which set of ascii characters will be used (default basic)
  -h, --help                       help for sobel
  -i, --input string               An image path which will be converted to ASCII
  -o, --output string              An output path for the converted image
  -t, --threshold uint8            Threshold between 0..255 to control intensity of assci in the edges of the image (default 130)
  -w, --width uint16               An image path which will be converted to ASCII (default 314)
```
