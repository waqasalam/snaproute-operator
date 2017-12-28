FROM debian
COPY ./bgp-crd /bgp-crd
ENTRYPOINT /bgp-crd
