FROM golang:latest 

COPY kube /
CMD ["/kube"]
