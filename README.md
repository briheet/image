# Image

This package provides Go methods to handle image file operations, specifically for compressing images.
Currently only works with .png | .jpg files.

### Prerequisites

This tool uses [pngquant](https://pngquant.org/) for image compression. Please make sure to install `pngquant` to ensure the tool works correctly.

### Usage

Run the executable with the following commands and arguments:

#### Commands

- **Help**  
  Display help information and usage examples.

  ```bash
  make
  ./bin/Image -h
  ```

- **Compress**  
  Compress a PNG file to reduce its size while retaining quality.

  ```bash
  make
  ./bin/Image -compress <SourceFile> <OutputFile>
  ```

- **GreyScale**  
  Convert an Given image file to GreyScale. It uses Average method to convert to greyscale by changing the pixel values.

  ```bash
  make
  ./bin/Image -greyscale <SourceFile> <OutputFile>
  ```
