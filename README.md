# Go Weather CLI

Go Weather is a command-line application that fetches and displays the local weather based on a provided zipcode.

## Installation

To install the application, clone the repository and navigate to the project directory:

```bash
git clone https://github.com/yourusername/go-weather.git
cd go-weather
```

Then, build the application using the following command:

```bash
go build -o weather ./cmd/weather
```

## Usage

To run the application, use the following command format:

```bash
./weather -zipcode <zipcode>
```

### Example

To get the weather for zipcode 76180, run:

```bash
./weather -zipcode 76180
```

## Dependencies

This project uses Go modules for dependency management. Ensure you have Go installed and set up on your machine.

## Contributing

Feel free to submit issues or pull requests for any improvements or bug fixes. 

## License

This project is licensed under the MIT License. See the LICENSE file for details.