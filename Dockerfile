FROM scratch

COPY kube /kube
ENTRYPOINT ["/kube"]