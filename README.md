# Kong Static Files Plugin


## Introduction

The Kong Static Files Plugin is a custom plugin for [Kong](https://konghq.com/), an open-source API Gateway and Microservices Management Layer. This plugin allows Kong to host and serve textual metadata files, making it easy to serve plain text content alongside your APIs. You can use it to serve simple configuration files, documentation, or any other textual data that needs to be delivered over HTTP.

This README provides information on how to install, configure, and use the Kong Static Files Plugin.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Dynamic Content Hosting**: Serve textual metadata files directly from Kong, allowing you to manage and update them easily.

- **Customizable Content Types**: You can specify the `Content-Type` for each metadata file, ensuring compatibility with various web services and browsers.

- **Configurable Paths**: Define which metadata files should be served at specific paths, making it easy to manage different metadata files in a single Kong instance.


## Prerequisites

Before using the Kong Static Files Plugin, you need to have the following components set up:

- [Kong](https://konghq.com/kong/) installed and configured on your system.
- Basic knowledge of Kong and its core concepts.

## Installation

To install the Kong Static Files Plugin, follow these steps:

1. Ensure that Kong is properly installed and running on your system.

2. Clone this repository or download the plugin source code to your local machine.

3. Build the plugin using the following command:

   ```shell
   go build -o static-files
   ```

4. Copy the generated binary (`static-files`) to a directory included in your system's `PATH` or within your Kong plugin directory.

5. Restart Kong or the Kong container to make the plugin available.

## Configuration

### Kong Configuration

To configure the "static-files" plugin, you need to set up the necessary Kong settings.

1. `kong_plugins`: In your Kong configuration (usually kong.conf), add the "static-files" plugin to the kong_plugins list to enable it:

    ```
    kong_plugins = bundled,static-files
    ```

2. `kong_pluginserver_names`: Configure the server names for the "static-files" plugin. In kong.conf, add an entry like this:

    ```
    kong_pluginserver_names = static-files
    ```

    This entry defines the server name for your plugin.

3. `kong_pluginserver_static_files_start_cmd`: Define the start command for the "static-files" plugin. In kong.conf, add an entry like this:

    ```
    kong_pluginserver_static_files_start_cmd = /usr/bin/static-files
    ```

    Replace /usr/bin/my-static-files-server with the actual path to your Go binary.

4. `kong_pluginserver_static-files_query_cmd`: Define the query command for the "static-files" plugin. In kong.conf, add an entry like this:

    ```
    kong_pluginserver_static-files_query_cmd = /usr/bin/static-files -dump
    ```

    Again replace /usr/bin/my-static-files-server with the actual path to your Go binary.

Check the [docker-compose.yml](docker-compose.yml) file in this repo for more details on how to setup Kong on Docker.

### Plugin Configuration via Kong Configuration File

The Kong Static Files Plugin can be configured through Kong's configuration file or using the Admin API. Here's how to configure the plugin:

1. Open your Kong configuration file (commonly `kong.conf`).

2. Add the following configuration for the Kong Static Files Plugin:

```yaml
_format_version: '3.0'
services:
  - name: static-files-service
    routes:
        - name: static-files-urls
        protocols:
        - http
        - https
        paths:
        - /assetlinks.json
        - /robots.txt
    plugins:
      - name: static-files
        config:
          /assetlinks.json:
            contentType: application/json
            content: '{"example": "metadata"}'
          /robots.txt:
            contentType: text/plain
            content: 'User-agent: *\nDisallow: /'
```

   Make sure to adjust the configuration according to your setup.

3. Save the configuration file.

4. Restart Kong to apply the changes.

### Configuration via Admin API

You can also configure the Kong Static Files Plugin using the Kong Admin API:

1. Create a new service if you haven't already:

   ```shell
   curl -i -X POST --url http://localhost:8001/services/ --data 'name=static-files-service' --data 'url=http://example.com'
   ```

   Replace `http://example.com` with the URL of your service.

2. Add the Kong Static Files Plugin to the service:

   ```shell
   curl -i -X POST --url http://localhost:8001/services/static-files-service/plugins/ --data 'name=static-files'
   ```

   This associates the plugin with the service.

3. Configure the plugin for the service:

   ```shell
   curl -i -X POST --url http://localhost:8001/services/static-files-service/plugins/static-files/config --data 'json={"paths": {"/path/to/your/file.txt": {"contentType": "text/plain", "content": "Your plain text content here"}}}'
   ```

   Replace `/path/to/your/file.txt` with the desired path and adjust the `contentType` and `content` fields accordingly.

4. The Kong Static Files Plugin is now configured and ready to use.

## Usage

The Kong Static Files Plugin serves static textual metadata files at the specified paths. Here's how to use it:

1. Make a request to your Kong service with a path that matches one of the configured paths.

   For example, if you configured the plugin with a path of `/path/to/your/file.txt`, you can access the content by making a request to:

   ```
   http://kong-host:port/path/to/your/file.txt
   ```

2. The plugin will return the content of the file with the specified `Content-Type` header.

   - If the file path matches one of the configured paths, the content will be served.
   - If the file path doesn't match or is not configured, a `404 Not Found` response will be returned.

3. Your plain text content will be served with the specified content type, such as `text/plain`.

## Contributing

We welcome contributions to the Kong Static Files Plugin. If you find a bug, have an enhancement in mind, or want to add new features, please open an issue on the GitHub repository. We appreciate your help in making this plugin even better.

## License

The Kong Static Files Plugin is released under the [MIT License](LICENSE). Please refer to the [LICENSE](LICENSE) file for more information.

---

**Disclaimer**: This Kong plugin is a community contribution and is not officially maintained or supported by Kong Inc. or the Kong community. Please use it at your own discretion.

Kong is a registered trademark of Kong Inc. All other trademarks are the property of their respective owners.