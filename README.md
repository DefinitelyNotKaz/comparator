
# Comparator

Comparator is a tool designed to identify pixels that do not match a [pxls.space](https://pxls.space/) template.

## Installation

### Precompiled Binaries

1. Download the appropriate binary for your operating system from the [Releases](https://github.com/DefinitelyNotKaz/comparator/releases) page.

2. On macOS or Linux, you may need to grant execute permissions:
   ```bash
   chmod +x ./comparator-linux-amd64
   ```

3. Open a terminal and execute the binary:
   ```bash
   ./comparator
   ```

### Building from Source

1. Ensure you have Go (Golang) installed on your system.

2. Clone or download this repository.

3. Install the necessary packages:
   ```bash
   go mod tidy
   ```

4. Build the application:
   ```bash
   go build
   ```

## Usage

To use Comparator, follow these steps:
   ```bash
   comparator -i image.png -p palette.json
   ```
   Replace `image.png` with the path to your image and `palette.json` with the path to pxls palette.

If you want to display what pixels are wrong on the console add `--verbose`.

You can see the options with `comparator --help`
