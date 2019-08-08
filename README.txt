########
#README#
########

Mozilla-observatory to prometheus exporter

Installation: (without docker)
1: Make
2: ./exporter_observatory

Installation: (with docker)
1: Make
1: Make docker
3: docker run -d -p 9229:9229 exporter_observatory

Then you can use it easily by using your navigator at  localhost:9229/metrics/host=... or `ip`:9229/metrics/host=.... if you want to ask from another computer
If you want to do multiple host you can separe them by a "," : "host=HOST1,HOST2,HOST3,HOST4..."


Khan S.
