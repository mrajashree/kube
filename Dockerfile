FROM scratch

COPY kube .
CMD["/kube"]