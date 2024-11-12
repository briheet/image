# Image

This package provides Go methods to handle image file operations, specifically for compressing PNG images.
Currently only works with .png files.

### Prerequisites

This tool uses [pngquant](https://pngquant.org/) for image compression. Please make sure to install `pngquant` to ensure the tool works correctly.

### Usage

Run the executable with the following commands and arguments:

#### Commands

- **Help**  
  Display help information and usage examples.

  ```bash
  ./binary -h
  ```

- **Compress**  
  Compress a PNG file to reduce its size while retaining quality.

  ```bash
  ./binary -compress <SourceFile.png> <OutputFile.png>
  ```
