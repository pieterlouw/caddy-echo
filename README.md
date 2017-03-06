This is project's goal is to serve (excuse the pun..) as an example/template to create a ServerType for [Caddy Server](https://github.com/mholt/caddy)

The server type is called `tcpecho` which will basically echo any traffic back to the caller. Potentially this could be used for other server types like TCP/TLS Proxying, HTTP Reverse Proxying etc.

## Resources ##

[Writing a Plugin: Server Type](https://github.com/mholt/caddy/wiki/Writing-a-Plugin:-Server-Type)

[Caddy Forum discussion](https://forum.caddyserver.com/t/server-types-other-than-http/65)

[CoreDNS (example of a server type other than HTTP)](https://github.com/coredns/coredns)