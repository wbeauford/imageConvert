FROM centos:7
RUN yum -y install ImageMagick ImageMagick-devel ImageMagick-perl
COPY imageCOnvert /
COPY 163535-1608224541.pdf /
CMD /imageConvert
