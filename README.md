# Healthchecker

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

Supported Go versions, see: https://endoflife.date/go
Supported architectures: amd64, arm64

Source code and issues: https://github.com/mvdkleijn/healthchecker

# Licensing

Healthchecker is licensed under GPLv3. The full details are available from the [LICENSE](/LICENSE) file.