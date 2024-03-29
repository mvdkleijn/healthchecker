# Healthchecker

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mvdkleijn/healthchecker?style=for-the-badge) [![Go Report Card](https://goreportcard.com/badge/github.com/mvdkleijn/healthchecker?style=for-the-badge)](https://goreportcard.com/report/github.com/mvdkleijn/healthchecker) [![Liberapay patrons](https://img.shields.io/liberapay/patrons/mvdkleijn?style=for-the-badge)](https://liberapay.com/mvdkleijn/) [![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/O4O7H6C73)

Simplistic Go utility that does a HTTP HEAD check to see if a server app is alive. Ideal for in distroless containers.

All it does (and likely will ever do) is:
- Send HTTP HEAD request to a specified URL;
- On return of HTTP status code 200, exit with value 0;
- Else exit with value 1;

## Usage

- Build it yourself or download a pre-built release.
- Point it to your server's URL.

### Examples

On the commandline:

```./healthchecker http://127.0.0.1:8080/api/v1/status```

or in a Dockerfile:

```
HEALTHCHECK --interval=5s --timeout=5s --retries=3 \
    CMD ["/healthchecker", "http://127.0.0.1:8080/api/v1/status"]
```

## Support

- Go versions, see: https://endoflife.date/go
- Architectures: amd64, arm64 on Windows and Linux

Source code and issues: https://github.com/mvdkleijn/healthchecker

## Licensing

Healthchecker is made available under the [MPL-2.0](https://choosealicense.com/licenses/mpl-2.0/)
license. The full details are available from the [LICENSE](/LICENSE) file.
