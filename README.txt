########
#README#
########

Mozilla-observatory to prometheus exporter

Installation: (without docker)
1: Make
2: ./http-observatory_exporter

Installation: (with docker)
1: Make
1: Make docker
3: docker run -d -p 9230:9230 http-observatory_exporter

Then you can use it easily by using your navigator at  localhost:9230/http-observatory/metrics/host=... or `ip`:9230/metrics/host=.... if you want to ask from another computer
If you want to do multiple host you can separe them by a "," : "host=HOST1,HOST2,HOST3,HOST4..."


Khan S.
