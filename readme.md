# tftp server in kubernetes

> Most network drivers for Kubernetes use NAT to expose services to the outside world which is not compatible with the way TFTP works natively. TFTP creates ephemeral port for each incoming connection. Kubernetes proxy cannot NAT these ports and as a result the initial connection over 69/UDP is established but then times out when the file download is handed over to an ephemeral port.
> Dnsmasq comes with a "single-port-mode" which strictly uses only port 69/UDP for all communication. While this is not 100% according to the TFTP RFC, it has been proven to be compatible with all common TFTP clients by the dnsmasq developers.

This is a basic server that use pin/tftp library to serve file using single-port-mode mode.

That allow to handle issues with NAT and Pod Network using a kubernetes service


Refs:
 - https://github.com/pin/tftp
 - https://github.com/siderolabs/metal-controller-manager/blob/master/internal/tftp/tftp_server.go