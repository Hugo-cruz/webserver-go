# Orange Webserver
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)


## Description

This is a basic CRUD API to explore the use of Docker, docker-compose, and an Orange Pi, to create a small webserver.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/Hugo-cruz/webserver-go
    ```

2. Change into the project directory:

    ```bash
    cd webserver-go
    ```

3. Run it:

    ```bash
    docker-compose up
    ```

## Usage

After running the above commands, the webserver should be up and running on port 8080. You can access the API endpoints to perform CRUD operations on devices. Here is a brief overview of the available endpoints:

- **GET /device/initialize**: Initialize the database with some example data.
- **GET /device/:id**: Retrieve a specific device by ID.
- **GET /device/search**: Search for devices by brand.
- **GET /device/list**: Retrieve a list of devices.
- **POST /device/create**: Create a new device.
- **DELETE /device/:id**: Delete a device by ID.
- **PUT /device/:id**: Update an existing device by ID.

You can use tools like `curl`, Postman, or your browser to interact with these endpoints.

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or issues, please open an issue or submit a pull request.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes and commit them (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
