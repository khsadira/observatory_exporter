########
#README#
########

Mozilla-observatory to prometheus exporter

Installation:

1: Make
2: Make docker
3: docker run -d -p 9229:9229 exporter_observatory

Then you can use it easly by using your navigator at localhost:9229/metrics/host=....
If you want to do multiple host you can separe them by a "," : localhost:9229/metrics/host=HOST1,HOST2,HOST3



Khan S.
